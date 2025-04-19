import Vue from 'vue'
import VueRouter from 'vue-router'
import ChatPage from '../views/ChatPage'
import Contacts from '../views/Contacts'
import Dialog from '../views/Dialog'
import SettingsPage from '../views/SettingsPage'
import MainLayout from '../views/MainLayout'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'main',
    component: MainLayout,
    children:[
      {
        path: '',
        name: 'chat',
        component: ChatPage
      },
      {
        path: '/contacts',
        name: 'contacts',
        component: Contacts
      },
      {
        path: '/dialog',
        name: 'dialog',
        component: Dialog
      },
    ]
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsPage
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
