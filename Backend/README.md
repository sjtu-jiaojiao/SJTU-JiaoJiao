# 后端
## 配置
- 依赖管理采用go mod，建议使用最新版golang。
- 国内网络建议使用goproxy或采用vendor模式以加快依赖下载速度。
- `config.json` 为总配置文件，项目配置请在 `consul.json` 中添加以支持配置中心。

`config.json` 配置：
- `deploy`：部署设置，`develop` 为开发模式，`product` 为生产模式
- `hosts`：公共服务地址配置
  - `consul-*`：consul服务地址

`consul.json` 配置：
- `test`：测试配置，请勿修改
- `namespace`：APIGateway的namespace

## 本地测试

    consul agent -ui -bind=127.0.0.1 -dev
    consul kv import @consul.json
  
## 目录结构

    api    RESTAPI入口
        auth    认证服务API
    srv    微服务组件
        auth    认证服务
    utils   通用模块
    test    单元测试
