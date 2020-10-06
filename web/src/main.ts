import Vue from 'vue'
import { App } from './components'
import { configureOneSignal, constructFirebaseApp } from './config'
import { HomePage } from '@/features/home'
import { MembersPage } from '@/features/members'
import { CreateEventPage } from '@/features/event-create'
import {
  LoginPage,
  LogoutPage,
  installFirebaseAuthHooksToVueRouter,
} from '@/features/auth'
import VueRouter from 'vue-router'
import VueHotkey from 'v-hotkey'
import 'firebase/app'
import 'firebase/auth'
import 'firebase/functions'

Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(VueHotkey)

const firebase = constructFirebaseApp()

const routes = [
  { path: '/', component: HomePage, meta: { requiresAuth: true } },
  { path: '/members', component: MembersPage, meta: { requiresAuth: true } },
  {
    path: '/create-event',
    component: CreateEventPage,
    meta: { requiresAuth: true },
  },
  { path: '/login', component: LoginPage },
  { path: '/logout', component: LogoutPage },
]
const router = new VueRouter({ routes })
installFirebaseAuthHooksToVueRouter(router)

new Vue({
  provide() {
    return {
      firebase,
    }
  },
  beforeMount() {
    configureOneSignal()
  },
  router,
  render: h => h(App),
}).$mount('#app')
