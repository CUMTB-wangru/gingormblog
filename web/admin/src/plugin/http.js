import Vue from 'vue'
import axios from 'axios'

let Url = 'http://localhost:3000/api/v1/'

axios.defaults.baseURLbaseURL = Url

// axios的请求拦截器，和后端的jwt鉴权验证规范匹配  固定写法
axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

// 将axios挂载到vue的全局组件上，这样所有的vue组件都可以用
Vue.prototype.$http = axios

export { Url }
