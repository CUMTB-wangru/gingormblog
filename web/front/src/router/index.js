import Vue from 'vue'
import VueRouter from 'vue-router'

const ArticleList = () =>
  import(/* webpackChunkName: "group-index" */ '../components/ArticleList.vue')
const Detail = () =>
  import(/* webpackChunkName: "group-detail" */ '../components/Details.vue')
const Category = () =>
  import(/* webpackChunkName: "group-category" */ '../components/CateList.vue')
const Search = () =>
  import(/* webpackChunkName: "group-search" */ '../components/Search.vue')

// 将vue-router 注册为全局组件
Vue.use(VueRouter)

//获取原型对象上的push函数
const originalPush = VueRouter.prototype.push
//修改原型对象中的push方法
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

// 动态路由 参数都是由编程式导航的url 携带过来的params参数
const routes = [
  { 
    path: '/',
    component: ArticleList,
    meta: { title: '欢迎来到GinBlog' }
  },
  {
    path: '/article/detail/:id',
    component: Detail,
    meta: { title: window.sessionStorage.getItem('title') },
    // 动态路由  传递所有的 params参数到Detail组件
    props: true
  },
  {
    path: '/category/:cid',
    component: Category,
    meta: { title: '分类信息' },
    // 动态路由  传递所有的 params参数到Category组件
    props: true
  },
  {
    path: '/search/:title',
    component: Search,
    meta: { title: '搜索结果' },
    // 动态路由  传递所有的 params参数到Search 组件
    props: true
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title ? to.meta.title : '加载中'
  }
  next()
})

export default router
