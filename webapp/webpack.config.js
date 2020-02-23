const webpack = require('webpack');
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    entry: './src/index.ts',
    mode: 'development',
    output: {
        filename: 'index.js',
        path: path.resolve(__dirname, 'dist')
    },
    devtool: 'source-map',
    resolve: {
        extensions: ['.json', '.ts', '.tsx', '.js'],
    },
    module: {
        rules: [
            {test: /\.json$/, use: 'json'},
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
            filename: 'index.html',
        }),
    ]
};
