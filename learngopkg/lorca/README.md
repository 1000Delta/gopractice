# Lorca 练习

## 简介

使用 HTML5 + Go 来创建使用 Chrome 浏览器作为UI层的桌面应用。

根据官方介绍，比较 Electron 更加轻量，不会绑定一个特定版本的 Chrome ，更倾向于复用已经安装的 Chrome。

## 入门

提供了基本API：

`UI#Load` 加载指定页面，参数为页面URL，可以服务器地址或页面数据

`UI#Bind` 绑定 Golang 函数到 Javascript

`UI#Eval` 解析 js 代码并执行

`UI#SetBounds`

`UI#Bounds`

溜了溜了