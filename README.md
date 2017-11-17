# GoGoGo

### 什么是Go
    Go是一种新的语言，一种并发的、带垃圾回收的、快速编译的语言。
   
###  Go开发环境搭建
#### Windows Go环境
1)	下载Go安装包（https://studygolang.com/dl）， 32 位请选择名称中包含 windows-386 的 msi 安装包，64 位请选择名称中包含 windows-amd64 的。
2)	下载好后运行，不要修改默认安装目录 C:\Go\，若安装到其他位置会导致不能执行自己所编写的 Go 代码。安装完成后默认会在环境变量 Path 后添加 Go 安装目录下的 bin 目录 C:\Go\bin\，并添加环境变量 GOROOT，值为 Go 安装根目录 C:\Go\ 。
3)	验证是否安装成功：在运行中输入 cmd 打开命令行工具，在提示符下输入 go，检查是否能看到 Usage 信息。输入 cd %GOROOT%，看是否能进入 Go 安装目录。若都成功，说明安装成功。
4) [开发工具](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.4.md)：
    ##### LiteIDE
    LiteIDE是一款专门为Go语言开发的跨平台轻量级集成开发环境（IDE），由visualfc编写。

    下载：http://sourceforge.net/projects/liteide/files
    ##### Visual Studio Code
    vscode是微软基于Electron和web技术构建的开源编辑器, 是一款很强大的编辑器。

    安装Visual Studio Code 最新版：https://code.visualstudio.com/

    安装Go插件点击右边的 Extensions 图标 搜索Go插件 在插件列表中，选择 Go，进行安装，安装之后，系统会提示重启 Visual Studio Code。
    ##### Gogland或者IntelliJ IDEA+Go插件
    Gogland是JetBrains公司推出的Go语言集成开发环境，是Idea Go插件是强化版。Gogland同样基于IntelliJ平台开发，支持JetBrains的插件体系。    
    下载地址:https://www.jetbrains.com/go/



###	 参考资料
#### 书籍
1)	[Go语言圣经（中文版）](http://books.studygolang.com/gopl-zh/)（《The Go Programming Language》 中文版本）
4)	[Go语言标准库](http://books.studygolang.com/The-Golang-Standard-Library-by-Example/)
####	网站
1)	[Golang中文社区](https://studygolang.com/)
2)	[菜鸟教程](http://www.runoob.com/go/go-tutorial.html)
3)	[Go by Example 中文](http://books.studygolang.com/gobyexample/)
4)	[Go指南](http://tour.studygolang.com/welcome/1)
####	GitHub
1)	[Go 编程基础](https://github.com/Unknwon/go-fundamental-programming)
2)	[Go Web基础](https://github.com/Unknwon/go-web-foundation/tree/v1)
3)	[build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md)
4)	GO学习笔记（自己记录的）
5)  [opensds](https://github.com/opensds)
####	Go
1.go连接数据库
  
  使用[Mysql数据库](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.2.md)
#### Go Web
1.Beego快速开始

下载安装

    go get github.com/astaxie/beego

创建文件 hello.go

    package main

    import "github.com/astaxie/beego"

    func main() {
        beego.Run()
    }

运行hello.go,访问http://localhost:8080

2.快速使用Bee工具

Bee工具安装

    go get github.com/beego/bee

new命令创建新项目

    bee new myproject

运行项目

    cd $GOPATH/src/myproject 
    bee run

访问http://localhost:8080/