import About from './components/About.vue'
import BackupOverview from './components/server/BackupOverview.vue'
import DetailsOverview from './components/server/DetailsOverview.vue'

export const routes = [
    { path: '/', component: BackupOverview},
    { path: '/about', component: About},
    { path:'/backups', component: BackupOverview},
    { path:'/details/:id', component: DetailsOverview}
];