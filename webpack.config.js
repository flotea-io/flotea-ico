var path = require('path')
var webpack = require('webpack')
//var Visualizer = require('webpack-visualizer-plugin');

module.exports = {
  mode: 'development',
  entry: {
    'main': './src/main.ts',
    'icoBuy': './src/icoBuy.ts',
    'voting': ['./src/addCarrier.ts', './src/removeCarrier.ts'], 
    'icoVoting': './src/icoVoting.ts',
    'applicants': './src/applicants.ts',
    'carriers': './src/carriers.ts',
    'style': './src/css/index.scss'
  },
  output: {
    path: path.resolve(__dirname, './static/js'),
    publicPath: '/static/js/',
    filename: "[name].js",
  },
  optimization: {
    minimize: true,
    splitChunks: {
      chunks: 'all',
      automaticNameDelimiter: '_',
    },
  },
  module: {
    rules: [
    {
      test: /\.s[ac]ss$/i,
      use: [
      {
        loader: 'file-loader',
        options: {
          name: '../css/[name].css',
        }
      },
      {
        loader: 'extract-loader'
      },
      {
        loader: 'css-loader?-url'
      },
      
      {
        loader: 'sass-loader'
      }
      ],
    },
    {
      test: /\.vue$/,
      loader: 'vue-loader',
      options: {
        loaders: {
            // Since sass-loader (weirdly) has SCSS as its default parse mode, we map
            // the "scss" and "sass" values for the lang attribute to the right configs here.
            // other preprocessors should work out of the box, no loader config like this necessary.
            'scss': 'vue-style-loader!css-loader!sass-loader',
            'sass': 'vue-style-loader!css-loader!sass-loader?indentedSyntax',
          }
          // other vue-loader options go here
        }
      },
      {
        test: /\.tsx?$/,
        loader: 'ts-loader',
        exclude: /node_modules/,
        options: {
          appendTsSuffixTo: [/\.vue$/],
        }
      },
      {
        test: /\.(png|jpg|gif|svg)$/,
        loader: 'file-loader',
        options: {
          name: '[name].[ext]?[hash]'
        }
      }
      ]
    },
    resolve: {
      extensions: ['.ts', '.js', '.vue', '.json'],
      alias: {
        'vue$': 'vue/dist/vue.esm.js'
      }
    },
    devServer: {
      historyApiFallback: true,
      noInfo: true
    },
    performance: {
      hints: false
    },
    devtool: (process.env.NODE_ENV === 'production') ? false : 'eval',
    plugins:[
      new webpack.ContextReplacementPlugin(/moment[\/\\]locale$/, /(en|pl|cs)$/),
    //new Visualizer()
    ]
  }

  if (process.env.NODE_ENV === 'production') {
    module.exports.devtool = '#source-map'
    module.exports.plugins = (module.exports.plugins || []).concat([
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: 'production'
      }
    }),
    new webpack.optimize.UglifyJsPlugin({
      mangle:   true,
      //sourceMap: true,
      extractComments: 'all',
      compress: {
        warnings: false
      }
    }),
    new webpack.LoaderOptionsPlugin({
      minimize: true
    })
    ])
}
