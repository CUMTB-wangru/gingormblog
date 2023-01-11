import Vue from 'vue'
import App from './App.vue'
import router from './router'

// 资源在哪里需要用到就在哪里导入import
import './plugin/http'
import './plugin/antui'
import './assets/css/style.css'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
