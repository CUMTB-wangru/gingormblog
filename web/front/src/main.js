import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import day from 'dayjs'

import './plugins/http'

// 时间格式过滤器
Vue.filter('dateformat', function(indate, outdate) {
  return day(indate).format(outdate)
})

Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
