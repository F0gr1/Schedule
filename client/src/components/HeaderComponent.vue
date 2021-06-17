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
            <router-link> {{userName}} </router-link>
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
export default {
  name: 'Header',
  data(){
      return{
        userName: '',
      };
  },

  methods:{
    getUserName(){
      
    },

    mounted(){
      this.getUserName();
    }
  }
}
</script>