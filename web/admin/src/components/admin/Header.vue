<template>
  <a-button type="danger" @click="loginOut">退出</a-button>
</template>

<script>
export default {
  methods: {
    loginOut() {
      // 清除浏览器里的token
      window.sessionStorage.clear('token')
      // 编程式导航  强制跳转
      this.$router.push('/login')
    }
  }
}
</script>
