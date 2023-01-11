import Vue from 'vue'

// 使用antui，按需引入，用什么就import什么  Vue.use(xxx)：注册成全局组件
import {
  ConfigProvider,
  Button,
  FormModel,
  Input,
  Icon,
  message,
  Layout,
  Menu,
  Row,
  Col,
  Table,
  Card,
  Pagination,
  Modal,
  Select,
  Switch,
  Upload
} from 'ant-design-vue'

// 配置message
message.config({
  top: `60px`,
  duration: 2,
  maxCount: 3
})

// 挂在到vue的原型上，所有vue组件都可以使用
Vue.prototype.$message = message
Vue.prototype.$confirm = Modal.confirm

// 注册成全局组件：所有的vue组件都可以用
Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Menu)
Vue.use(Row)
Vue.use(Col)
Vue.use(Table)
Vue.use(Card)
Vue.use(Pagination)
Vue.use(ConfigProvider)
Vue.use(Modal)
Vue.use(Select)
Vue.use(Switch)
Vue.use(Upload)
