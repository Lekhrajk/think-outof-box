import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/Class I',
    name: 'classI',
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "Class I" */ '../views/ClassI.vue')
  },
  {
    path: '/Class II',
    name: 'classII',
    component: () => import(/* webpackChunkName: "Class II" */ '../views/ClassII.vue')
  },
  {
    path: '/Class III',
    name: 'classIII',
    component: () => import(/* webpackChunkName: "Class III" */ '../views/ClassIII.vue')
  },
  {
    path: '/Class IV',
    name: 'classIV',
    component: () => import(/* webpackChunkName: "Class IV" */ '../views/ClassIV.vue')
  },
  {
    path: '/Class V',
    name: 'classV',
    component: () => import(/* webpackChunkName: "Class V" */ '../views/ClassV.vue')
  },
  {
    path: '/Class VI',
    name: 'classVI',
    component: () => import(/* webpackChunkName: "Class VI" */ '../views/ClassVI.vue')
  },
  {
    path: '/Class VII',
    name: 'classVII',
    component: () => import(/* webpackChunkName: "Class VII" */ '../views/ClassVII.vue')
  },

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
