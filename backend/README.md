# 后端
## 环境依赖
- golang
- consul
- nodejs/npm
- 建议系统环境为Linux系统，windows下可能出现各种问题

## 其他依赖

    go get github.com/micro/micro
    go get github.com/smartystreets/goconvey
    go get github.com/oxequa/realize
    go get github.com/micro/protoc-gen-micro
    go get github.com/golang/protobuf/protoc-gen-go
    npm install apidoc -g
    
## 本地启动
一键启动：

    go run run.go run_*.go

## 本地默认地址
- consul服务：`localhost:8500`
- micro web：`localhost:8082`
- APIGateway：`localhost:8080`
- goconvey：`localhost:8400`
- realize热启动：`localhost:5002`
- API地址：`localhost:8080/[version]/[service]/[router]`
- doc文档：`localhost:8080/[version]/doc`

## 配置
- 依赖管理采用go mod，建议使用最新版golang。
- 国内网络建议使用goproxy以加快依赖下载速度。
- `config.json` 为总配置文件，项目配置请在 `consul.json` 中添加以支持配置中心。

### 配置文件
`config.json` 配置：

- `deploy`：部署设置，`develop` 为开发模式，`product` 为生产模式
- `hosts`：服务地址配置
  - `consul_*`：consul服务地址
  - `url_*`：url地址
  - `test_*`：测试配置，请勿修改
- `config_ttl`：远程配置超时时间（秒）
- `test`：测试配置，请勿修改

`consul.json` 配置：

- `test`：测试配置，请勿修改
- `srv_config`：微服务配置
  - `namespace`：Service的namespace

- `api_config`：API配置
  - `version`：API版本
  - `namespace`：APIGateway的namespace

- `sys_config`：系统配置
  - `auth_url`：OAuth auth url
  - `token_url`：OAuth token url
  - `logout_url`：OAuth logout url

### 环境变量
**环境变量设置敏感信息，相关内容禁止提交到仓库**

- `JJ_CONFIGPATH`：config路径，覆盖默认值，发布二进制文件需设置
- `JJ_CLIENTID`：OAuth client_id
- `JJ_CLIENTSECRET`：OAuth client_secret
- `JJ_MARIADBUSER`：MariaDB用户名
- `JJ_MARIADBPWD`：MariaDB密码
- `JJ_MONGODBUSER`：MongoDB用户名
- `JJ_MONGODBPWD`：MongoDB密码
- `JJ_JWTSECRET`：JWT secretkey

## 目录结构

    api             RESTAPI入口
    doc             API文档
    srv             微服务组件
    database        数据库模块
    utils           通用模块
    vendor          go库缓存

## JWT Token
token存在三个字段：

- `id`：唯一id
- `role`：用户身份
  - 1：用户
  - 10：系统管理员
- `exp`：过期时间，默认30分钟
