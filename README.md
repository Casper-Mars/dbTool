# dbTool

## 程序主界面

![avatar](imgs/mainWindow.png)


## 数据项说明

* ipPort

> 数据库的ip和端口，形式：ip:port。例如：192.168.123.74:3306

* username

> 数据库的登录用户名

* password

> 数据库的登录密码

* 数据库名称

> 要导出表结构的数据库名称，多个数据库用"，"分隔。不能出现不必要的空格，否则会认为空格也是数据库名称的一部分

* 存储位置

> 导出的word文档存放的位置

## 数据库导出支持列表

|数据库|是否支持|
|---|---|
|mysql|支持|
|sqlserver|不支持|
|oracle|不支持|
|SQLit|不支持|

## 平台支持

|平台|是否支持|
|---|---|
|Windows|支持|
|linux|支持|
|macOS|支持|

## 环境要求

### macOS

* glib-2.64

> 可以通过安装gtk+3. Mac系统通过brew可快速安装：brew install gtk+3

### linux

桌面版的发行版一般都自带gtk环境。

### windows

解压后运行dbtool.exe