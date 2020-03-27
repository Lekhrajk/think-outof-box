/*=========================================================================================
  File Name: router.js
  Description: All routes
  ----------------------------------------------------------------------------------------
  Item Name: Xongl user panel
  Author: Xongl cloud private limited
==========================================================================================*/


import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const router = new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    scrollBehavior () {
        return { x: 0, y: 0 }
    },
    routes: [

        {
    // =============================================================================
    // MAIN LAYOUT ROUTES
    // =============================================================================
            path: '/dashboard',
            component: () => import('./layouts/main/Main.vue'),
            children: [
              {
                path: '/home',
                name: 'home',
                component: () => import('./views/Home.vue')
              },
              {
                path: '/select-template',
                name: 'select-template-page',
                component: () => import('@/views/pages/SelectTemplate.vue'),
                meta: {
                  breadcrumb: [
                      { title: 'Home', url: '/home' },
                      { title: 'SelectTemplate', active: true },
                  ],
                  pageTitle: 'Deploy VM',
                },
              },
              {
                path: '/select-template/category/:id',
                name: 'deploy-vm-page',
                component: () => import('@/views/pages/CreateVm.vue'),
                meta: {
                    breadcrumb: [
                        { title: 'Home', url: '/home' },
                        { title: 'Select Template', url: '/select-template' },
                        { title: 'Deploy VM', active: true },
                    ],
                    parent: 'select-template-page',
                    rule: 'editor'
                },
              },
              {
                path: '/select-template/category/vmstatus/:id',
                name: 'vm-status-page',
                component: () => import('@/views/pages/VmStatus.vue'),
                meta: {
                    breadcrumb: [
                        { title: 'Home', url: '/home' },
                        { title: 'VM Status', active:true },
                        { title: 'Current Status'},
                      
                    ],
                    parent: 'select-template-page',
                    rule: 'editor'
                },
              },
              {
                path: '/vm-usage',
                name: 'vm-usage-page',
                component: () => import('@/views/pages/VmUsage.vue')
              },
            ],
        },
    // =============================================================================
    // FULL PAGE LAYOUTS
    // =============================================================================
        {
            path: '',
            component: () => import('@/layouts/full-page/FullPage.vue'),
            children: [
        // =============================================================================
        // PAGES
        // =============================================================================
              {
                path: '/login',
                name: 'page-login',
                component: () => import('@/views/pages/Login.vue')
              },
              {
                path: '/resetPassword',
                name: 'reset-password',
                component: () => import('@/views/login/ResetPassword.vue')
              },
              {
                path: '/',
                redirect: '/login'
              },
              {
                path: '/pages/error-404',
                name: 'page-error-404',
                component: () => import('@/views/pages/Error404.vue')
              },
            ]
        },
        // Redirect to 404 page, if no match found
        {
            path: '*',
            redirect: '/pages/error-404'
        }
    ],
})

router.afterEach(() => {
  // Remove initial loading
  const appLoading = document.getElementById('loading-bg')
    if (appLoading) {
        appLoading.style.display = "none";
    }
})

export default router
