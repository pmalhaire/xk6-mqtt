
const path = require('path');

module.exports = {
  mode: 'production',
  entry: {
    vehicle_state_ota_test: './examples/vehicle_state_ota_test.ts',
  },
  output: {
    path: path.resolve(__dirname, 'dist'), // eslint-disable-line
    libraryTarget: 'commonjs',
    filename: '[name].bundle.js',
  },
  module: {
    rules: [{ test: /\.ts$/, use: 'babel-loader' }],
  },
  target: 'web',
  externals: /k6(\/.*)?/,
};
