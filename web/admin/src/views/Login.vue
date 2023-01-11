<template>
  <div class="container">
    <div class="loginBox">
      <!-- :model="formdata" 动态绑定formdata对象的数据 
        :rules="rules" 绑定验证规则 
        ref="loginFormRef" 表单引用对象 -->
      <a-form-model ref="loginFormRef" :rules="rules" :model="formdata" class="loginForm">

        <!-- prop="username" 指定验证username属性 要和data中的属性对上 -->
        <a-form-model-item prop="username">
          <!-- v-model="formdata.username"  动态绑定formdata对象的username属性 -->
          <a-input v-model="formdata.username" placeholder="请输入用户名">
            <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)" />
          </a-input>
        </a-form-model-item>

        <!-- prop="password" 指定验证password属性 要和data中的属性对上 -->
        <a-form-model-item prop="password">
          <!-- 给键盘的回车键绑定监听事件login -->
          <a-input
            v-model="formdata.password"
            placeholder="请输入密码"
            type="password"
            v-on:keyup.enter="login"
          >
            <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)" />
          </a-input>
        </a-form-model-item>

        <a-form-model-item class="loginBtn">
          <a-button type="primary" style="margin:10px" @click="login">登录</a-button>
          <a-button type="info" style="margin:10px" @click="resetForm">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formdata: {
        username: '',
        password: '',
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          {
            min: 4,
            max: 12,
            message: '用户名必须在4到12个字符之间',
            trigger: 'blur',
          },
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          {
            min: 6,
            max: 20,
            message: '密码必须在6到20个字符之间',
            trigger: 'blur',
          },
        ],
      },
    }
  },
  methods: {
    resetForm() {
      // 通过拿到表单引用对象loginFormRef调用resetFields() 重置表单
      this.$refs.loginFormRef.resetFields()
    },
    login() {
      // 通过拿到表单引用对象loginFormRef调用validate() 验证登录信息
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return this.$message.error('输入非法数据，请重新输入')

        // this.$http.post('login', this.formdata)  向后端请求url：'http://localhost:3000/api/v1/login'
        //  post方式  携带参数：formdata
        // 将后端返回的参数 结构解析 赋值给res
        const { data: res } = await this.$http.post('login', this.formdata)
        if (res.status != 200) return this.$message.error(res.message)

        // 将浏览器生成的token存储到res.token
        window.sessionStorage.setItem('token', res.token)
        // 编程式路由导航  跳转到/index
        this.$router.push('/index')
      })
    },
  },
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: #282c34;
}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: #fff;
  position: absolute;
  top: 50%;
  left: 70%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}

.loginBtn {
  display: flex;
  justify-content: flex-end;
}
</style>
