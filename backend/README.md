# 后端
## Docker测试环境

    sudo docker build -t sjtujj/test . -f Dockerfile_test

或者使用预编译环境（自行修改***为正确值）

    sudo docker run -d -p 10000:22 -p 10001:80 -v $(pwd):/root/backend -e JJ_CLIENTID=*** -e JJ_CLIENTSECRET=*** -e JJ_JWTSECRET=*** sjtujj/test

注意在linux环境挂载后本目录权限可能会变成`root:root`，请预先备份

默认密码：
- ssh: `root/sjtujj123123`
- mysql: `root/123456`
- mongodb: `root/123456`

端口：
- ssh: 10000
- web: 10001

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

## Test
测试json，一个创建文件的例子：
```
{
    "case": [
        {
            "input": {
                "file": ""
            },
            "output": {
                "status": -1,
                "fileID": ""
            }
        },
        {
            "input": {
                "file": "valid"
            },
            "output": {
                "status": 1,
                "fileID": "#NOTEMPTY#"
            },
            "verify": {
                "fileID": "#fileID#",
                "file": "valid"
            }
        }
    ]
}
```
一个查询文件的例子：
```
{
    "data": [
        {
            "fileID": "000000000000000000000001",
            "file": "valid"
        }
    ],
    "case": [
        {
            "input": {
                "fileID": ""
            },
            "output": {
                "status": -1,
                "file": "",
                "size": 0
            }
        },
        {
            "input": {
                "fileID": "invalid"
            },
            "output": {
                "status": -1,
                "file": "",
                "size": 0
            }
        },
        {
            "input": {
                "fileID": "100000000000000000000000"
            },
            "output": {
                "status": 2,
                "file": "",
                "size": 0
            }
        },
        {
            "input": {
                "fileID": "000000000000000000000001"
            },
            "output": {
                "status": 1,
                "file": "valid",
                "size": 5
            }
        }
    ]
}
```
一个删除文件的例子：
```
{
    "data": [
        {
            "fileID": "000000000000000000000001",
            "file": "valid1"
        },
        {
            "fileID": "000000000000000000000002",
            "file": "valid2"
        }
    ],
    "case": [
        {
            "input": {
                "fileID": ""
            },
            "output": {
                "status": -1
            }
        },
        {
            "input": {
                "fileID": "invalid"
            },
            "output": {
                "status": -1
            }
        },
        {
            "input": {
                "fileID": "100000000000000000000000"
            },
            "output": {
                "status": 2
            }
        },
        {
            "input": {
                "fileID": "000000000000000000000001"
            },
            "output": {
                "status": 1
            },
            "verify": [
                {
                    "fileID": "000000000000000000000001",
                    "file": "valid1",
                    "_exist": false
                },
                {
                    "fileID": "000000000000000000000002",
                    "file": "valid2",
                    "_exist": true
                }
            ]
        },
        {
            "input": {
                "fileID": "000000000000000000000002"
            },
            "output": {
                "status": 1
            },
            "verify": [
                {
                    "fileID": "000000000000000000000001",
                    "file": "valid1",
                    "_exist": false
                },
                {
                    "fileID": "000000000000000000000002",
                    "file": "valid2",
                    "_exist": false
                }
            ]
        }
    ]
}
```

- `data` 内为预存数据
- `case` 内为测试用例，每个测试用例参数如下：
  - `input`：必须，测试输入
  - `output`：必须，期望输出，不允许省略，特殊参数如下
    - `_error`: 可选，默认值为`false`，为`true`则期望抛出异常（忽略其他输出参数）
  - `verify`:可选，验证数据库中是否存在/不存在指定值，特殊参数如下
    - `#var#`：使用`output`中的`var`的值代替（部分参数支持）
    - `_exist`：可选，默认值为`true`，为`false`则为判断不存在
    - 
通用特殊参数（部分参数支持）：
- `#NOTEMPTY#`：字符串非空
- `#NOTZERO#`：数字非0
- `#NOT#xxx`：字符串非`xxx`
- `#NOT#x`：数字非`x`