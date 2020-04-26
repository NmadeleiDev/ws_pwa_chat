module.exports = {
    publicPath: process.env.NODE_ENV === 'production' ? process.env.PUBLIC_PATH : '/',
    outputDir: process.env.BUILD_DIR || '../nginx/static',
    lintOnSave: true,
    devServer: {
        proxy: {
            '/api/v1': {
                pathRewrite: { '^/api/v1': '' },
                target: process.env.BACKEND_HOST || 'http://localhost:2222',
                changeOrigin: true,
                overlay: {
                    warnings: true,
                    errors: true
                }
            },
        },
    },
};
