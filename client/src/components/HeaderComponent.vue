<template>
  <div class="header">
    <v-app-bar flat>

    </v-app-bar>
    <v-navigation-drawer
      v-model="drawer"
      absolute
      bottom
      temporary
    >
      <v-list>
        <v-list-item>

          <v-list-item-title>
            <router-link> Home </router-link>
          </v-list-item-title>

          <v-list-item-title>
            <router-link> {{this.userName}} </router-link>
          </v-list-item-title>

          <v-list-item-title>
            <router-link> Logout </router-link>
          </v-list-item-title>

        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
import axios from "../plugins/axios";
export default {
  name: 'Header',
  data(){
      return{
        userName: null,
      };
  },

  methods:{
    getUserName(){
      let url = 'http://localhost:8080/login';
      axios .get(url)
            .then(response => {
              this.userName = response.data['message'];
            })
            .catch(e => {
              console.log(e);
            })
    },

    mounted(){
      this.getUserName();
    }
  }
}
</script>