import Vue from 'vue'
import Router from 'vue-router'

// const _import = require('./_import_' + process.env.NODE_ENV)
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
  {
    path: '',
    component: Layout,
    redirect: '/dashboard/dashboard'//默认页
  },
  { path: '/login', component: () => import('@/views/login'), name: '登录NxAdmin', hidden: true },
  { path: '/404', component: () => import('@/views/errorPage/404'), hidden: true },
  { path: '/401', component: () => import('@/views/errorPage/401'), hidden: true },
  // 报表
  {
    path: '/dashboard',
    component: Layout,
    meta: { title: '首页', icon: 'dashboard' },
    children: [
      {
        path: 'dashboard',
        name: 'dashboard',
        component: () => import('@/views/dashboard/dashboard'),
        meta: { title: '首页', icon: 'dashboard' }
      }
    ]
  },
  // 表格
  {
    path: '/shop',
    component: Layout,
    redirect: '/shopList',
    name: 'shop',
    alwaysShow:true,
    meta: {
      title: '商户管理',
      icon: 'table'
    },
    children: [
      {
        path: 'shopList',
        name: 'shopList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '商户列表' }
      },
      {
        path: 'BaseForm',
        name: 'BaseForm',
        hidden: true, // 不在侧边栏线上
        component: () => import('@/views/form/BaseForm'),
        meta: { title: '修改商户' }
      }
    ]
  },

  {
    path: '/user',
    component: Layout,
    redirect: '/userList',
    name: 'user',
    alwaysShow:true,
    meta: {
      title: '用户管理',
      icon: 'table'
    },
    children: [
      {
        path: 'userList',
        name: 'userList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '用户列表' }
      }
    ]
  },
  {
    path: '/order',
    component: Layout,
    redirect: '/orderList',
    name: 'order',
    alwaysShow:true,
    meta: {
      title: '订单管理',
      icon: 'table'
    },
    children: [
      {
        path: 'orderList',
        name: 'orderList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '订单列表' }
      }
    ]
  },
  {
    path: '/blog',
    component: Layout,
    redirect: '/blogList',
    name: 'blog',
    alwaysShow:true,
    meta: {
      title: '笔记管理',
      icon: 'table'
    },
    children: [
      {
        path: 'blogList',
        name: 'blogList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '笔记列表' }
      }
    ]
  },
  {
    path: '/admin',
    component: Layout,
    redirect: '/adminList',
    name: 'admin',
    alwaysShow:true,
    meta: {
      title: '系统管理',
      icon: 'table'
    },
    children: [
      {
        path: 'adminList',
        name: 'adminList',
        component: () => import('@/views/table/complex-table'),
        meta: { title: '管理员列表' }
      }
    ]
  },
  {
    path: '/self',
    component: Layout,
    redirect: '/updatePwd',
    name: 'self',
    alwaysShow:true,
    meta: {
      title: '个人中心',
      icon: 'table'
    },
    children: [
      {
        path: 'updatePwd',
        name: 'updatePwd',
        component: () => import('@/views/form/BaseForm'),
        meta: { title: '修改密码' }
      }
    ]
  },

]

export default new Router({
  // mode: 'history', //后端支持可开
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})
export const asyncRouterMap = []
