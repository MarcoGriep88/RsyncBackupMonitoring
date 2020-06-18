import Vue from 'vue'
import VueResource from 'vue-resource'
import VueRouter from 'vue-router';

import { routes } from './routes';

import App from './App.vue'
import Header from './components/shared/Header.vue'
import Footer from './components/shared/Footer.vue'

import BackupOverview from './components/server/BackupOverview.vue'
import DetailsOverview from './components/server/DetailsOverview.vue'

Vue.use(VueResource);
Vue.use(VueRouter);

const router = new VueRouter({
  routes
});

Vue.component("app-header", Header);
Vue.component("app-footer", Footer);

Vue.component("app-backup-overview", BackupOverview);
Vue.component("app-backup-details", DetailsOverview);

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
