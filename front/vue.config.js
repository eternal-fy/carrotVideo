module.exports = {
    publicPath: '/',
    devServer:{
        proxy:{
            '/api':{
                target:"http://localhost:9999/",//API服务器地址（更换为你的路径）
                ws:true,//代理websocked
                changeOrigin:true,//虚拟的站点需要更管origin
                pathRewrite:{
                    '^/api':'api'//重写路径
                }
            }
        },
        historyApiFallback: true,
        allowedHosts:"all"
    }
}