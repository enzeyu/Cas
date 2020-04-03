//Vue-loader在15.*之后的版本都是 vue-loader的使用都是需要伴生 VueLoaderPlugin的(看下方注释)
const path = require('path');

const VueLoaderPlugin = require('vue-loader/lib/plugin');
const HTMLPlugin = require('html-webpack-plugin');
const webpack = require('webpack');
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const {
  CleanWebpackPlugin
} = require('clean-webpack-plugin')

const isDev = process.env.NODE_ENV === 'development';

const config = {
  target: 'web', //入口
  entry: path.join(__dirname, 'src/index.js'),
  output: {
    filename: 'bundle.[hash:8].js',
    path: path.join(__dirname, 'dist'),
  },
  //webpack原生只支持js文件类型，只支持ES5语法，我们使用以.vue文件名结尾的文件时，需要为其指定loader
  module: {
    rules: [{
        test: /\.vue$/,
        loader: 'vue-loader'
      },
      {
        test: /\.jsx$/,
        loader: 'babel-loader'
      },
      //CSS文件没有用到，这部分配置仅是展示，可忽略掉
      // {
      //   test: /\.css$/,
      //   use: [
      //     'style-loader',
      //     'css-loader'
      //   ]
      // },

      //将小于1024d的图片转为base64，减少http请求
      {
        test: /\.(gif|png|jpg|jpeg|svg)$/,
        use: [{
          loader: 'url-loader',
          options: {
            limit: 1024,
            name: '[name]-aaa.[ext]',
            outputPath:'assets/img/'
          }
        }]
      }
    ]
  },
  plugins: [
    //判断相关环境
    // new webpack.DefinePlugin({
    //   'process.env': {
    //     NODE_ENV: isDev ? '"development"' : '"production"'
    //   }
    // }),
    //在每次构建前清理 /dist 文件夹，因此只会生成用到的文件。
    new CleanWebpackPlugin(),
    new HTMLPlugin(),
    new VueLoaderPlugin(),
    new MiniCssExtractPlugin({
      //类似于webpackOptions.output中相同选项的选项
      //所有选项都是可选的
      filename: '[name].css',
      chunkFilename: '[id].css',
      ignoreOrder: false //启用以删除有关冲突顺序的警告
    })
  ],
  optimization: {
    splitChunks: {
      chunks(chunk) {
        return chunk.name !== 'my-excluded-chunk'
      }
    }
  }
}

if (isDev) {
  config.module.rules.push({
    //css预处理器，使用模块化的方式写css代码
    //stylus-loader专门用来处理stylus文件，处理完成后变成css文件，交给css-loader.webpack的loader就是这样一级一级向上传递，每一层loader只处理自己关心的部分。
    //styl配置需要根据开发环境来配置，开发环境就是这么用。正式环境中↓ ↓ ↓
    test: /\.styl/,
    use: [
      'vue-style-loader',
      'css-loader',
      {
        loader: 'postcss-loader',
        options: {
          sourceMap: true
        }
      },
      'stylus-loader'
    ]
  });
  // config.devtool = '#cheap-module-eval-source-map'
  config.devServer = {
    port: 8000, //端口号
    host: '0.0.0.0',
    //提示错误与警告功能，
    overlay: {
      errors: true
    },
    hot: true
    //这个功能是当我们运行devServer的时候，会自动帮我们打开浏览器
    // open:true
  }
} else {
  config.output.filename = '[name].[chunkhash:8].js'
  config.module.rules.push({
    test: /\.styl/,
    use: [{
        loader: MiniCssExtractPlugin.loader,
        options: {
          publicPath: './',
          hmr: process.env.NODE_ENV === 'development'
        }
      },
      'css-loader',
      {
        loader: 'postcss-loader',
        options: {
          sourceMap: true
        }
      },
      'stylus-loader'
    ]
  });
  config.plugins.push(
    // new webpack.HotModuleReplacementPlugin(),
    // new webpack.NoEmitOnErrorsPlugin(),
    new MiniCssExtractPlugin({
      filename: 'styles.[chunkhash].[name].css',
      chunkFilename: '[id].css',
      ignoreOrder: false
    })
  )
}

module.exports = config;