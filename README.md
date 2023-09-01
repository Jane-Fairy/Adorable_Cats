### 1、项目介绍
这个项目是想打造一个猫咖平台，后期会做流浪猫平台，目的是为了让城市里流浪的猫猫有一个美好的家，让每个猫猫都能被人类爱戴.

adulatian has been take part in this project.


### 10、帮助文档

生成api代码go-ctl示例

```cmd
goctl api go --api ./app/usercenter/cmd/api/desc/usercenter.api  --dir ./app/usercenter/cmd/api -style gozero
```

生成proto代码go-ctl示例

```cmd
goctl rpc protoc ./pb/usercenter.proto --go_out=. --go-grpc_out=. --zrpc_out=. -style goZero     
```


根据sql文件生成template代码
```cmd
goctl model mysql ddl --src user.sql --dir ../      -------------->  template: go init template #获取原始模板
```


### 11、额外控制
```cmd
关闭日志：logx.DisableStat() //去除控制台输出
```
