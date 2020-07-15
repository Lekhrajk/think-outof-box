import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/list',
    name: 'list',
    component: () => import(/* webpackChunkName: "create" */ '../views/List.vue')
  },
  {
    path: '/todo/:id',
    name: 'fullview',
    component: () => import(/* webpackChunkName: "fullview" */ '../views/FullView.vue')
  },
  {
    path: '*',
    name: "error",
    component: () => import(/* webpackChunkName: "fullview" */ '../views/Error.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
