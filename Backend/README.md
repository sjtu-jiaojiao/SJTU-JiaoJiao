# 后端
## 配置
- 依赖管理采用go mod，建议使用最新版golang。
- 国内网络建议使用goproxy或采用vendor模式以加快依赖下载速度。

## 目录结构

    api    RESTAPI入口
        auth    认证服务API
    srv    微服务组件
        auth    认证服务
    utils   通用模块
