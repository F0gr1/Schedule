import Vue from 'vue'
import './plugins/bootstrap-vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'

require('bootstrap')

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
