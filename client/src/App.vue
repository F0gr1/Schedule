<template>
  <v-app>

    <header>
      <HeaderComponent/>
    </header>

    <v-main>
      <router-view/>
    </v-main>

    <footer>
      <FooterComponent/>
    </footer>

  </v-app>
</template>

<script>
import HeaderComponent from './components/HeaderComponent.vue'
import FooterComponent from './components/FooterComponent.vue'

export default {
  name: 'App',

  data: () => ({
    //
  }),

  components:{
    HeaderComponent,
    FooterComponent
  }
};
</script>
