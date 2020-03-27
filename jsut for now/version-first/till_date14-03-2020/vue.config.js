/*=========================================================================================
  File Name: vue.config.js
  Description: configuration file of vue
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
==========================================================================================*/


module.exports = {
publicPath: '/',
devServer: {
/*    port: 8080,*/
    host: 'localhost',
proxy: {
      '/*': {
        target: 'http://localhost:9090',
        ws: true,
        changeOrigin: true
      }
  }
},
  transpileDependencies: [
    'vue-echarts',
    'resize-detector'
  ],
  configureWebpack: {
    optimization: {
      splitChunks: {
        chunks: 'all'
      }
    }
  }
}

