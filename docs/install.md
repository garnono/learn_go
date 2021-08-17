# 安装

### 支持的系统
* Linux
* FreeBSD
* Mac OS X（也称为 Darwin）
* Windows

### 安装方式
* 安装包
    * Windows 下载扩展名为 msi 的文件进行安装
    ```
  默认情况下 .msi 文件会安装在 c:\Go 目录下。
  你可以将 c:\Go\bin 目录添加到 Path 环境变量中。
  添加后你需要重启命令窗口才能生效
  ```
  * Mac OS X 下载扩展名为 pkg 的文件进行安装
   ```
  安装目录在 /usr/local/go/ 下
  或者通过 brew 进行安装：brew install go
  ```
  
* 源码——UNIX/Linux/Mac OS X, 和 FreeBSD 安装
    * 下载 tar.gz 的文件
    ```
  下载地址：
    https://golang.org/dl/ 
    https://golang.google.cn/dl/
  ```
    * 解压
    ```
  tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz
  ```
    * 添加环境变量
    ```
  export PATH=$PATH:/usr/local/go/bin
  ```
  
```
通过 go version 检测是否安装成功
```