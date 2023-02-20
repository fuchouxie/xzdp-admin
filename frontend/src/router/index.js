import Vue from 'vue'
import Router from 'vue-router'

// consts _import = require('./_import_' + process.env.NODE_ENV)
// in development-env not use lazy-loading, because lazy-loading too many pages will cause webpack hot update too slow. so only in production use lazy-loading;
// detail: https://panjiachen.github.io/vue-element-admin-site/#/lazy-loading

Vue.use(Router)

/* Layout */
import Layout from '../views/layout/Layout'

/**
* hidden: true                   if `hidden:true` will not show in the sidebar(default is false)
* alwaysShow: true               if set true, will always show the root menu, whatever its child routes length
*                                if not set alwaysShow, only more than one route under the children
*                                it will becomes nested mode, otherwise not show the root menu
* redirect: noredirect           if `redirect:noredirect` will no redirct in the breadcrumb
* name:'router-name'             the name is used by <keep-alive> (must set!!!)
* meta : {
    title: 'title'               the name show in submenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar,
  }
**/
export const constantRouterMap = [
  // {
  //   path: '/',
  //   component: Layout,
  //   redirect: '/login',
  //   name: 'login',
  //   hidden: true
  // },
  {
    path: '',
    component: Layout,
    redirect: '/dashboard/dashboard'
  },
  { path: '/login', component: () => import('@/views/login'), name: '登录NxAdmin', hidden: true },
  { path: '/404', component: () => import('@/views/errorPage/404'), hidden: true },
  { path: '/401', component: () => import('@/views/errorPage/401'), hidden: true },

  // 报表
  {
    path: '/dashboard',
    component: Layout,
    meta: { title: 'dashboard', icon: 'dashboard' },
    children: [
      {
        path: 'dashboard',
        name: 'dashboard',
        component: () => import('@/views/dashboard/dashboard'),
        meta: { title: 'dashboard', icon: 'dashboard' }
      }
    ]
  },
  {
    path: '/table',
    component: Layout,
    redirect: '/table/ShopList',
    name: 'shop',
    hidden: false,
    alwaysShow: true,
    meta: {
      title: '商店管理',
      icon: 'table'
    },
    children: [
      {
        path: 'ShopList',
        name: 'ShopList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '商店列表' }
      },
      {
        path: 'VoucherList',
        name: 'VoucherList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '代金券管理' }
      }
    ]
  },

  {
    path: '/table',
    component: Layout,
    redirect: '/table/BlogList',
    name: 'blog',
    hidden: false,
    alwaysShow: true,
    meta: {
      title: '笔记管理',
      icon: 'table'
    },
    children: [
      {
        path: 'BlogList',
        name: 'BlogList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '笔记列表' }
      }
    ]
  },

  {
    path: '/table',
    component: Layout,
    redirect: '/table/OrderList',
    name: 'order',
    hidden: false,
    alwaysShow: true,
    meta: {
      title: '订单管理',
      icon: 'table'
    },
    children: [
      {
        path: 'OrderList',
        name: 'OrderList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '订单列表' }
      }
    ]
  },

  {
    path: '/table',
    component: Layout,
    redirect: '/table/UserList',
    name: 'user',
    hidden: false,
    alwaysShow: true,
    meta: {
      title: '权限管理',
      icon: 'table'
    },
    children: [
      {
        path: 'UserList',
        name: 'UserList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '用户列表' }
      }
    ]
  }
]

export default new Router({
  // mode: 'history', //后端支持可开
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})
export const asyncRouterMap = [
]
