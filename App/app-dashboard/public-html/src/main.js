import Vue from 'vue'
import App from './components/App.vue'
import VueResource from 'vue-resource'

import About from './components/About.vue'

import Header from './components/shared/Header.vue'
import Footer from './components/shared/Footer.vue'

import BackupOverview from './components/server/BackupOverview.vue'

Vue.use(VueResource);

Vue.component("app-header", Header);
Vue.component("app-footer", Footer);

Vue.component("app-backup-overview", BackupOverview);

const routes = {
  '/': App,
  '/about': About
}

new Vue({
  el: '#app',
  data: {
    currentRoute: window.location.pathname
  },
  computed: {
    ViewComponent () {
      return routes[this.currentRoute] || NotFound
    }
  },
  render (h) { return h(this.ViewComponent) }
})
