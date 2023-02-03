const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
  entry: "./index.js",
  output: {
    filename: "[name].bundle.js",
    path: path.resolve(__dirname, "dist"),
  },
  module: {
    rules: [
      {
        test:  /\.(s(a|c)ss)$/,
        use: ["style-loader", "css-loader", 'sass-loader'],
      },
      {
        test: /\.html$/,
        use: ["html-loader"],
      },
      {
        test: /\.m?js$/,
        exclude: /node_modules/,    
      },
      {
        test: /\.(png|jpg|gif)$/,
        use: [
          {
            loader: "file-loader",
            options: {
              outputPath: "assets/images",
            },
          },
        ],
      },
    ],
  },
  plugins: [
    // Plugin for copying the index.html file to the dist directory.
    new HtmlWebpackPlugin({
      template: "./index.html",
    }),
  ],
  devServer: {
    static: "./dist",
    hot: true,
  },
  devtool: "inline-source-map",
};
