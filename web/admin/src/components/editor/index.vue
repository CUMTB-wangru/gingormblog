<template>
  <div>
     <!-- props属性实现父传子： 父组件index 给子组件Editor 传递数据init、content，需要在子组件Editor中用props声名接收的变量 -->
    <Editor :init="init" v-model="content"></Editor>
  </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
import './tinymce.min.js'
import './icons/default/icons.min.js'
import '../../assets/tinymce/themes/silver/theme.min.js'

import './langs/zh_CN'

// 注册插件
import './plugins/preview/plugin.min.js'
import './plugins/paste/plugin.min.js'
import './plugins/wordcount/plugin.min.js'
import './plugins/code/plugin.min.js'

import './plugins/image/plugin.min.js'
import './plugins/imagetools/plugin.min.js'
import './plugins/media/plugin.min.js'
import './plugins/codesample/plugin.min.js'
import './plugins/lists/plugin.min.js'
import './plugins/table/plugin.min.js'

export default {
  components: { Editor },

  // 不是路由props 这里接收的是AddArt组件传过来的context内容
  props: {
    value: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      init: {
        language: 'zh_CN',
        height: '600px',
        plugins: 'preview paste wordcount code imagetools image media codesample lists table',
        branding: false,
        paste_data_images: true,
        // 富文本框需要的功能 
        toolbar: [
          'undo redo | styleselect |fontsizeselect| bold italic underline strikethrough |alignleft aligncenter alignright alignjustify |blockquote removeformat |numlist bullist table',
          'preview paste code codesample |image media',
        ],
        //上传图片
        images_upload_handler: async (blobInfo, succFun, failFun) => {
          let formdata = new FormData()
          formdata.append('file', blobInfo.blob(), blobInfo.name())
          const { data: res } = await this.$http.post('upload', formdata)
          succFun(res.url)
          failFun(this.$message.error('上传图片失败'))
        },
        imagetools_cors_hosts: ['*'],
        imagetools_proxy: '',
      },
      content: this.value,
    }
  },
  // 事件监听： value发生改变就调用回调函数，重新赋值context---简写形式，直接写成函数形式    
  watch: {
    value(newV) {
      this.content = newV
    },
    content(newV) {
      // this.$emit 触发事件
      // 监听input事件，有改变，就把newV传入context
      this.$emit('input', newV)
    },
  },
}
</script>

<style scoped>
/* 导入样式 */
@import url('../../assets/tinymce/skins/ui/oxide/skin.min.css');
@import url('../../assets/tinymce/skins/content/default/content.min.css');
</style>
