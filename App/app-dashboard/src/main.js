import Vue from 'vue'
import App from './components/App.vue'
import VueResource from 'vue-resource'

import About from './components/About.vue'
import Backup from './components/Backup.vue'
import License from './components/License.vue'

import Header from './components/shared/Header.vue'
import Sidebar from './components/shared/Sidebar.vue'
import Footer from './components/shared/Footer.vue'

import CustomerOverview from './components/server/UpgradableOverview.vue'
import BackupOverview from './components/server/BackupOverview.vue'
import LicenseOverview from './components/server/LicenseOverview.vue'
import AboutOverview from './components/server/AboutOverview.vue'

Vue.use(VueResource);

Vue.component("app-header", Header);
Vue.component("app-sidebar", Sidebar);
Vue.component("app-footer", Footer);

Vue.component("app-upgradable-overview", CustomerOverview);

Vue.component("app-backup-overview", BackupOverview);
Vue.component("app-license-overview", LicenseOverview);
Vue.component("app-about-overview", AboutOverview);

const routes = {
  '/': App,
  '/about': About,
  '/backup': Backup,
  '/license': License
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
