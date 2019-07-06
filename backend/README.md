# 后端
## 其他依赖

    go get github.com/micro/micro
    go get github.com/smartystreets/goconvey
    go get github.com/oxequa/realize
    go get github.com/micro/protoc-gen-micro
    go get github.com/golang/protobuf/protoc-gen-go
    npm install apidoc -g
    
## 本地启动
一键启动：

    ./runsrv.sh

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
- `JJ_CLIENTID`：OAuth client_id
- `JJ_CLIENTSECRET`：OAuth client_secret

## 目录结构

    api             RESTAPI入口
        auth        认证服务API
    srv             微服务组件
        auth        认证服务
    utils           通用模块
    apidoc.json     apidoc配置文件
    config.json     总配置文件
    consul.json     项目配置文件
    go.mod          go mod文件
    runsrv.sh       本地服务启动脚本
    servedoc.go     doc文档服务
    generate.sh     proto文件生成脚本

