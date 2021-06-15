package main

import (
	sessioninfo "Schedule/server/SessionInfo"
	"Schedule/server/controller"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var LoginInfo sessioninfo.SessionInfo

func main() {
	engine := gin.Default()
	store := cookie.NewStore([]byte("select"))
	engine.Use(sessions.Sessions("mysession", store))
	engine.Use(static.Serve("/", static.LocalFile("../client/src", false)))
	engine.GET("/login", func(c *gin.Context) {
		c.String(200, `{"message":"hello, hello, hello"}`)
	})

	engine.POST("/login", controller.NewLogin().LoginK)

	engine.GET("/singup", func(c *gin.Context) {
		c.HTML(200, "singup.html", gin.H{})
	})
	engine.POST("/singup", controller.NewLogin().SingUp)
	menu := engine.Group("/menu")
	menu.Use(sessionCheck())
	{
		menu.GET("/top", controller.GetMenu)
	}

	engine.POST("/logout", controller.PostLogout)
	engine.Run(":8080")
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		LoginInfo.Name = session.Get("name")

		// セッションがない場合、ログインフォームをだす
		if LoginInfo.Name == nil {
			log.Println(session)
			log.Println("ログインしていません")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort() // これがないと続けて処理されてしまう
		} else {
			c.Set("name", LoginInfo.Name) // ユーザidをセット
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}
