/* eslint-disable @typescript-eslint/no-var-requires */
const webpack = require('webpack');
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = {
  entry: './src/index.tsx',
  mode: 'development',
  output: {
    filename: 'index.js',
    path: path.resolve(__dirname, 'dist'),
  },
  devtool: 'source-map',
  resolve: {
    extensions: ['.json', '.ts', '.tsx', '.js'],
  },
  module: {
    rules: [
      { test: /\.json$/, use: 'json' },
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ],
  },
  plugins: [
    new webpack.optimize.ModuleConcatenationPlugin(),
    new HtmlWebpackPlugin({
      template: 'static/index.html',
      filename: 'index.html',
    }),
    new CopyWebpackPlugin([{ from: 'static/icons', to: 'icons' }]),
  ],
  devServer: {
    contentBase: path.join(__dirname, 'dist'),
    compress: true,
    port: 3001,
    historyApiFallback: true,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        pathRewrite: { '^/api': '' },
      },
    },
  },
};
