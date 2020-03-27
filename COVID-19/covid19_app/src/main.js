import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'bootstrap';
import "./assets/css/app.css"
import "./assets/vendors/mdi/css/fontawesome.min.css"
import "./assets/vendors/mdi/js/all.js"

import "./assets/vendors/aos/css/aos.css"
import "./assets/vendors/jquery-flipster/css/jquery.flipster.css"
import "./assets/css/style.css"
import jQuery from 'jquery';
window.$ = window.jQuery = jQuery;
import 'popper.js';

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
