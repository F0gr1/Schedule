package controller

import (
	"log"
	"net/http"

	"Schedule/server/crypto"
	"Schedule/server/db"
	"Schedule/server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

func LoginM(c *gin.Context, name string) {
	session := sessions.Default(c)
	session.Set("name", name)
	session.Save()
}
func getUser(username string) model.Login {
	db := db.Connection()
	var user model.Login
	db.First(&user, "name = ?", username)
	db.Close()
	return user
}
func (l *Login) LoginK(c *gin.Context) {
	db := db.Connection()
	defer db.Close()
	log.Println("ログイン処理")
	name := c.PostForm("name")

	LoginM(c, name) // // 同じパッケージ内のログイン処理

	dbPassword := getUser(c.PostForm("name")).Pass
	log.Println(dbPassword)
	// フォームから取得したユーザーパスワード
	formPassword := c.PostForm("pass")

	// ユーザーパスワードの比較
	if err := crypto.CompareHashAndPassword(dbPassword, formPassword); err != nil {
		log.Println("ログインできませんでした")

		c.Abort()
	} else {
		log.Println("ログインできました")
		c.Redirect(http.StatusMovedPermanently, "/menu/top")
	}
}
func (l *Login) SingUp(c *gin.Context) {
	var form Login
	if err := c.Bind(&form); err != nil {
		c.Abort()
	} else {
		username := c.PostForm("name")
		password := c.PostForm("pass")
		// 登録ユーザーが重複していた場合にはじく処理PasswordEncrypt

		passwordEncrypt, _ := crypto.PasswordEncrypt(password)
		db := db.Connection()
		defer db.Close()
		if err := db.Create(&model.Login{Name: username, Pass: passwordEncrypt}).GetErrors(); err != nil {

		}
		c.Redirect(302, "/login")
	}
}
func PostLogout(c *gin.Context) {
	log.Println("ログアウト処理")
	Logout(c) // 同じパッケージ内のログアウト処理

	// ログインフォームに戻す
	c.HTML(http.StatusOK, "login.html", gin.H{
		"name":         "",
		"ErrorMessage": "",
	})
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	log.Println("セッション取得")
	session.Clear()
	log.Println("クリア処理")
	session.Save()
}
func GetMenu(c *gin.Context) {
	name, _ := c.Get("name") // ログインユーザの取得

	c.HTML(http.StatusOK, "menu", gin.H{"name": name})
}
