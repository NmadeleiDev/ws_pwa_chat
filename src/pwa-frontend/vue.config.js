module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  "publicPath": "/",
  "outputDir": "../nginx/static",
  "lintOnSave": true,
  "devServer": {
    "proxy": {
      "/api/v1": {
        "pathRewrite": {
          "^/api/v1": ""
        },
        "target": "http://localhost:2222",
        "changeOrigin": true,
        "overlay": {
          "warnings": true,
          "errors": true
        }
      }
    }
  },
}