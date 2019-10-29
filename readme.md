# CLI 命令行实用程序开发实战 - Agenda

## 任务目标

1. 熟悉 go 命令行工具管理项目
2. 综合使用 go 的函数、数据结构与接口，编写一个简单命令行应用 agenda
3. 使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令,不会影响其他命令的代码
4. 项目部署在 Github 上，合适多人协作，特别是代码归并
5. 支持日志（原则上不使用debug调试程序）

## 知识点

### JSON序列化与反序列化

```go
date,err:=json.Marshal(users)//将users结构体数组变成json格式
json.Unmarshal([]byte(jsonStr), &users)//将json格式的json字符串解析到结构体数组里
```

## Cobra库的应用

```bash
cobra init --pkg-name agenda
cobra add register
cobra add delete
cobra add query
cobra add login
cobra add logout
```

## 实现功能

### register

register an account

```bash
agenda register -u [username] -p [password] -e [email] -t [telephone]
```

### delete

delete current account, should login first

```bash
agenda delete
```

### query

query all acounts information, should login first

```bash
agenda query
```

### login

log in

```bash
agenda login -u [username] -p [password]
```

### logout

log out

```bash
agenda logout
```



