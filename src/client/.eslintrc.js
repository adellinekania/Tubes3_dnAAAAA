const path = require('path');

module.exports = {
  env: {
    browser: true,
    es2021: true,
    'vue/setup-compiler-macros': true,
  },
  extends: [
    'plugin:vue/vue3-recommended',
    'airbnb-base',
  ],
  parserOptions: {
    ecmaVersion: 2022,
    sourceType: 'module',
  },
  plugins: [
    'vue',
    'import',
  ],
  rules: {
    'vue/script-setup-uses-vars': 'error',
    'no-plusplus': ['error', { allowForLoopAfterthoughts: true }],
    'import/prefer-default-export': 'off',
  },
  settings: {
    'import/resolver': {
      alias: {
        map: [
          ['@', path.resolve(__dirname, 'src')],
        ],
      },
    },
  },
};
