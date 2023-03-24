# One-Time Pad Service

One-Time Pad Service 是一个基于 Go 语言的 Web 应用程序，它提供了用户注册、登录、添加等基本功能。
# 依赖

该应用程序依赖于以下第三方库：

    html/template
    math/rand
    net/http
    strconv
    strings
    time

# 运行

要运行该应用程序，你需要在终端中进入项目根目录并执行以下命令：
```
go run main.go
```

接下来，你可以在浏览器中访问 http://localhost:8080 来使用该应用程序。
# 功能

该应用程序提供了以下功能：

    用户注册
    用户登录
    添加数据

# 代码结构

该应用程序的代码结构如下：

    main.go: 应用程序的主要入口点，它包含了路由和处理器函数。
    MailUtils: 发送邮件的工具函数。
    Otp: 生成一次性密码的工具函数。
    User: 用户相关的数据结构和数据库操作。

# 联系方式

liweijun0302@gmail.com
