module.exports = {
  presets: ['@vue/cli-plugin-babel/preset'],

  // 配置Ant-ui---按需引入
  plugins: [
    [
      'import',
      { libraryName: 'ant-design-vue', libraryDirectory: 'es', style: 'css' }
    ]
  ]
}
