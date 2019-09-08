# 后端
## 环境依赖
- python3
- golang
- consul
- nodejs/npm
- **建议系统环境为Linux系统，windows下make命令可能无法正常工作**

模型配置文档：https://github.com/sjtu-jiaojiao/SJTU-JiaoJiao/tree/master/backend/srv/tag

依赖安装：

    make deps
    
## 本地开发
### Makefile一键命令
- 查看帮助：`make` 或 `make help`
- 依赖安装：`make deps`
- 构建二进制文件：`make build`
- 构建docker：`make docker`
- 单元测试：`make test`
- 清理目录：`make clean`
- 生成文档：`make doc`
- 生成proto：`make proto`
- 添加新服务：`make add`
- 部署服务：`make deploy` （**警告：本地开发禁用**）

### 服务管理程序
一键启动：

linux:

    go run run.go run_linux.go

windows:

    go run run.go run_windows.go

## 本地服务默认地址
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
    template        模板目录，用于自动生成

## JWT Token
token存在三个字段：

- `id`：唯一id
- `role`：用户身份
  - 1：用户
  - 10：系统管理员
- `exp`：过期时间，默认30分钟
