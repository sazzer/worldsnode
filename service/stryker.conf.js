module.exports = function(config) {
  config.set({
    testRunner: 'jest',
    coverageAnalysis: 'off',
    timeoutFactor: 10,

    reporters: [
        'progress', 
        'clear-text', 
        'dots', 
        'html'
    ],
    htmlReporter: {
        baseDir: 'build/reports/mutation/html'
    },


    mutator: 'javascript',
    mutate: [
        'src/**/*.js',
        '!src/**/__tests__/*.js'
    ],
      
    transpilers: [
        'babel'
    ],
    babelrcFile: './.babelrc',

    jest: {
        config: require('./unit-test.json')
    }
  });
};

