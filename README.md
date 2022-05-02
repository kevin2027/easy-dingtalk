# dingtalk

dingtalk playground

## 安装

```shell
  go get -u github.com/Kevin2027/easy-dingtalk
```

## 引入代码

```go
import (
    "github.com/Kevin2027/easy-dingtalk/dingtalk"
    "github.com/Kevin2027/easy-dingtalk/utils"
)
```

## 创建服务

```go
srv, _, err = dingtalk.NewDingtalk(utils.DingtalkOptions{
    AppKey:    config.AppKey,
    AppSecret: config.AppSecret,
    AgentId:   config.AgentId,
})

```
