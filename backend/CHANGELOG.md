# 后端API更新日志
## v0.1.2 2019-07-20
### 禁用
1. `AddContent`、`DeleteContent`、`AddAvatar`接口已声明但功能暂不可用

### 新增
1. `GET /file/:fileId` 获取文件
2. `POST /avatar` 添加头像（暂不可用）

### API修改
1. API参数错误将统一返回400
2. `PUT /sellInfo` 修改为POST
3. `PUT /content` 修改为POST
4. `PUT /user` 修改为POST
5. `POST /user` 修改为PUT

### 权限修改
1. `GET /user/:userId` 允许管理员访问
2. `PUT /user` 用户自身不允许修改部分字段

## v0.1.1 2019-07-17
### 新增
1. `GET /auth` 登陆接口添加状态码3为用户被冻结
2. `GET/PUT /user` 用户查询添加`role`字段标识用户身份

### BUG修复
1. `GET /sellInfo/:sellInfoId` 返回字段`sellInfoState`修正为`status`（文档错误实际返回正常）
2. `PUT /content` 修复`content`字段检测问题

## v0.1.0 2019-07-17
第一个稳定版本