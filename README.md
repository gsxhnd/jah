# jaha

环境变量模版文件 [template](template.env)

启动服务

```shell
dotenvx run -f .dev.env -- go run ./server
```

```bash
swag init -d ./server/router -g init.go --outputTypes yaml --pdl 3
```
