# webpack
在 echo 中使用 webpack

使用了 `symfony` 的 encore, 更多使用请参考

https://symfony.com/doc/current/frontend/encore/installation.html

https://symfony.com/doc/current/frontend.html

## 环境需求
## 安装

执行命令`go get -u github.com/gosoeasy/webpack`

依赖于 `@symfony/webpack-encore`, 安装 `yarn add @symfony/webpack-encore --dev`

新建 webpack.config.js
```js
var Encore = require('@symfony/webpack-encore');

Encore
// directory where compiled assets will be stored
    .setOutputPath('template/build/')
    // public path used by the web server to access the output path
    .setPublicPath('/build')
    // only needed for CDN's or sub-directory deploy
    //.setManifestKeyPrefix('build/')

    /*
     * ENTRY CONFIG
     *
     * Add 1 entry for each "page" of your app
     * (including one that's included on every page - e.g. "app")
     *
     * Each entry will result in one JavaScript file (e.g. app.js)
     * and one CSS file (e.g. app.css) if you JavaScript imports CSS.
     */
    .addEntry('app', './assets/js/app.js')
    //.addEntry('page1', './assets/js/page1.js')
    //.addEntry('page2', './assets/js/page2.js')

    // will require an extra script tag for runtime.js
    // but, you probably want this, unless you're building a single-page app
    .enableSingleRuntimeChunk()

    .cleanupOutputBeforeBuild()
    .enableSourceMaps(!Encore.isProduction())
    // enables hashed filenames (e.g. app.abc123.css)
    .enableVersioning(Encore.isProduction())

// uncomment if you use TypeScript
//.enableTypeScriptLoader()

// uncomment if you use Sass/SCSS files
//.enableSassLoader()

// uncomment if you're having problems with a jQuery plugin
//.autoProvidejQuery()
;

module.exports = Encore.getWebpackConfig();
```

```


新建 `assets/css/app.css` `assets/js/app.js`
```html
// app.css


// assets/js/app.js

require('../css/app.css');

console.log('Hello Webpack Encore');
```


根目录下:
```
|____webpack.config.js
|____assets
| |____css
| | |____app.css
| |____js
| | |____app.js
```

## 使用

### yarn encore

```bash
# 开发环境
yarn encore dev
# 实时编译, 刷新浏览器即可查看最新变更
yarn encore dev --watch
# 生产环境
yarn encore production
```

以上命令将会在 `pubic` 下生成 `build` 目录, 目录中是编译好的文件

```
|____build
| |____entrypoints.json
| |____runtime.js
| |____app.css
| |____manifest.json
| |____app.js
```

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>think-webpack</title>
    {:encore_link_tags('app')}
</head>
<body>
    <h1>think-webpack</h1>
    {:encore_script_tags('app')}
</body>
</html>
```