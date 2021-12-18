# mvc based on gin & gorm

## reference

- [gin](https://gin-gonic.com/docs/)
- [gorm](https://gorm.io/docs/index.html)
- [validate](https://gookit.github.io/validate/#/)

## quick start

```shell
make build
./mvc
```

- 浏览器打开`localhost:5000/admin`可看到前端页面
- 测试sql中用户名密码分别为admin 123456
- 测试登陆，POST如下json
```json
{
    "username": "admin",
    "password": "123456"
}
```

## modules

### config

- 以yml文件放在`config/${env}`目录下，可配置不同环境下的配置文件
- `pkg/config` 下定义了配置结构体，在`internal/global`包中初始化
- 具体运行参数可运行`./mvc -h`查看

### mysql

- 在`pkg/db`下定义初始化，在`internal/global`包中初始化

### redis

- 在`pkg/redis`下定义初始化，在`internal/global`包中初始化
- 通过怕日志redis配置中的env设置redis中key的prefix，实际存入的key为`${prefix}.${key}`

```go
package redis_sample

import "mvc/internal/global"

func redis_sample() {
    _ = global.RedisClient.PutToRedis("key", "value", 3600)
    data, _ = global.RedisClient.GetFromRedis("key")
}
```

### router

- 定义在`internal/server/api_router`和`internal/server/backend_router`下
- 使用`router.StaticFS("/admin", http.Dir("./web"))`挂载前端项目目录

### log

- 在config中定义输出位置
- 使用`tools/logger`下的方法输出log

### app
`internal/app`下定义
- 以模块为单位，`controller`, `model`, `validate`应写在统一包下
- 返回格式在`tools/formatter`下定义

### validate

- 定义`validate`变量
- `rules`, `scenes`分别为字段规则与场景，使用时，应先把struct使用`tools/validate`下的`Struck2Map`进行转换

```go
type paramType struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
var params paramType

// validate
if err := Validate.ValidateMap(validate.Struct2Map(params), "login"); err != nil {
	c.JSON(http.StatusOK, formatter.ApiReturn(constants.CodeError, "参数验证失败", err.Error()))
	return
}
```