<template>
  <a-layout class="container">
    <!-- 组件Nav展示 -->
    <Nav></Nav>
    <a-layout>
      <a-layout-header class="headerBtn">
        <!-- 展示Header组件 -->
        <Header></Header>
      </a-layout-header>
      <a-layout-content>
        <!-- :key="$route.fullPath" 绑定一个独一无二的数据 -->
        <router-view :key="$route.fullPath"></router-view>
      </a-layout-content>
      <a-layout-footer class="footer">
        <!-- 展示Footer组件 -->
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
// 在Admin组件中导入其它的组件
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'
import Header from '../components/admin/Header'
export default {
  // 在组件中导入别的组件需要注册导入的组件
  components: { Nav, Footer, Header },
}
</script>

<style scoped>
.container {
  height: 100%;
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
