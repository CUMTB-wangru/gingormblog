import Vue from 'vue'
import VueRouter from 'vue-router'
const Login = () => import(/* webpackChunkName: "Login" */ '../views/Login.vue')
const Admin = () => import(/* webpackChunkName: "Admin" */ '../views/Admin.vue')

// 页面路由组件
const Index = () => import(/* webpackChunkName: "Index" */ '../components/admin/Index.vue')
const AddArt = () => import(/* webpackChunkName: "AddArt" */ '../components/article/AddArt.vue')
const ArtList = () => import(/* webpackChunkName: "ArtList" */ '../components/article/ArtList.vue')
const CateList = () => import(/* webpackChunkName: "CateList" */ '../components/category/CateList.vue')
const UserList = () => import(/* webpackChunkName: "UserList" */ '../components/user/UserList.vue')
const Profile = () => import(/* webpackChunkName: "UserList" */ '../components/user/Profile.vue')
const CommentList = () => import(/* webpackChunkName: "UserList" */ '../components/comment/commentList.vue')

// 路由重复点击捕获错误
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location, onResolve, onReject) {
  if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

// 前端路由：这里的路由负责前端页面组件之间的跳转，不涉及后端api
const routes = [
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '请登录'
    },
    component: Login
  },
  {
    path: '/',
    name: 'admin',
    meta: {
      title: 'GinBlog 后台管理页面'
    },
    component: Admin,
    children: [
      {
        path: 'index',
        component: Index,
        meta: {
          title: 'GinBlog 后台管理页面'
        }
      },
      {
        path: 'addart',
        component: AddArt,
        meta: {
          title: '新增文章'
        }
      },
      {
        path: 'addart/:id',
        component: AddArt,
        meta: {
          title: '编辑文章'
        },
        // vue-router 布尔模式：传递所有的params参数，通过props 传值 props写在哪个组件就给哪个组件传值
        // 在AddArt组件中用 props: ['id'] 表示要接收的数据
        props: true
      },
      {
        path: 'artlist',
        component: ArtList,
        meta: {
          title: '文章列表'
        }
      },
      {
        path: 'catelist',
        component: CateList,
        meta: {
          title: '分类列表'
        }
      },
      {
        path: 'userlist',
        component: UserList,
        meta: {
          title: '用户列表'
        }
      },
      {
        path: 'profile',
        component: Profile,
        meta: {
          title: '个人设置'
        }
      },
      {
        path: 'commentlist',
        component: CommentList,
        meta: {
          title: '评论管理'
        }
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

// 前置路由守卫
router.beforeEach((to, from, next) => {

  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()

  // 拿到存储在浏览器里的token
  const userToken = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!userToken) {
    // 强制跳转到/Login
    next('/login')
  } else {
    next()
  }
})

export default router
