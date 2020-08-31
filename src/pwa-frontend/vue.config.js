module.exports = {
  "pwa": {
    "manifestOptions": {
      "name": "Enchat",
      "short_name": "Enchat",
    }
  },
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
        "target": "http://localhost:2223",
        "changeOrigin": true,
        "overlay": {
          "warnings": true,
          "errors": true
        }
      },
      "/api/media": {
        "pathRewrite": {
          "^/api/media": ""
        },
        "target": "http://localhost:2224",
        "changeOrigin": true,
        "overlay": {
          "warnings": true,
          "errors": true
        }
      }
    }
  },
}