# miniKeys——基于GO语言的钥匙管理工具

----

## 背景：

互联网普及，越来越多的网站都需要用户进行登录注册，而这就出现了一个用户的痛点问题：用户经常忘记密码，普通网站通常设计忘记密码这个选项，但是用户操作起来还得输入验证码等等信息，即使有Chrome浏览器的记住密码功能，但是如果更换了设备，用户无法输入密码，仍然无法保证用户使用密码的便捷性。

## 架构

miniKeys使用前后端分离的方式完成用户功能，前端使用微信小程序，后端使用GO语言完成项目。

前端在微信小程序中使用[LinUi](https://doc.mini.talelin.com/)，详细介绍可[查看](https://github.com/yamonc/mini_program_wechat)。

后端采用Go语言，web框架使用的是[gin](https://github.com/gin-gonic/gin)，ORM使用的是[Gorm](https://gorm.io/zh_CN/docs/index.html)。

## 主要完成功能

用户登录：输入唯一秘钥登录系统。可以查看用户目录下的所有钥匙信息。

获得当前用户下的所有钥匙信息。

查看指定钥匙信息。

保存和更新指定钥匙信息。

删除指定钥匙信息。

## 测试

得到当前用户下的所有钥匙信息：

```json
[
    {
        "ID": 9,
        "KeyName": "11",
        "Url": "11",
        "Password": "11",
        "Account": "11",
        "Back": "11",
        "Category": 0,
        "UserId": 7
    },
    {
        "ID": 10,
        "KeyName": "11",
        "Url": "11",
        "Password": "11",
        "Account": "11",
        "Back": "11",
        "Category": 0,
        "UserId": 7
    },
    {
        "ID": 11,
        "KeyName": "11",
        "Url": "11",
        "Password": "11",
        "Account": "11",
        "Back": "11",
        "Category": 0,
        "UserId": 7
    },
    {
        "ID": 12,
        "KeyName": "测试钥匙1",
        "Url": "yamong.top",
        "Password": "yyhqqq11",
        "Account": "yamon.top",
        "Back": "备注",
        "Category": 0,
        "UserId": 7
    }
]
```

## 项目总结

- [miniKeys——钥匙库项目系列一（项目介绍以及Go项目部署）](https://yamon.top/blog/93)

## 待优化

连接数据库需要优化，目前是每次发出请求都需要连接数据库。

