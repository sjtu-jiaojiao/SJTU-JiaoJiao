# 后端API更新日志
## v0.5.0
### 新增
1. 推送模块上线

### BUG修复
1. 修复message被错误标记为已读的问题

## v0.4.1
### BUG修复
1. `GET /tag` 权限错误修复

## v0.4.0
### 新增
1. `GET /tag` 获取AI tag

## v0.3.3 2019-09-09
### 移除
1. `transaction` 模块被完全移除

### 新增
1. `PUT "/buyInfo/:buyInfoID"` 修改状态
2. `PUT "/sellInfo/:sellInfoID"` 修改状态

### BUG修复
1. 修复message被错误标记为已读的问题

## v0.3.2 2019-09-08
### 新增
1. `GET /message/userID` 增加获取旧消息的功能

### BUG修复
1. `GET /message` 修复`offset`和`limit`失效问题

## v0.3.1 2019-09-06
### BUG修复
1. 修复文件上传相关检测

## v0.3.0 2019-08-02
### 修复
1. `DELETE /content` 删除单独content file已可使用

### 新增
1. message相关功能已可以使用

## v0.2.2 2019-08-01
### 禁用
1. `DELETE /content` 删除单独content file暂不可用

### 新增
1. tag相关功能已可以使用

### API修改
1. `PUT /content` 参数均为必选，将删除content file功能移至 `DELETE`
2. `DELETE /content` 允许删除content file（暂不可用）

### BUG修复
1. 字段值检查已添加，非法数值会被视为0

## v0.2.1 2019-07-30
### 修复
1. content相关API已可使用

### API修改
1. `POST /transaction` `userID`参数修改为`fromUserID`
2. `POST /content` 现在会多返回新创建的fileID

### 参数修改
1. 查询接口限制单条最大值为100

## v0.2.0 2019-07-29
### 新增
1. `GET /content/:contentID` 获取content
2. `GET /transaction` 获取交易
3. `POST /transaction` 添加交易
4. `PUT /transaction` 修改交易

## v0.1.4 2019-07-23
### API修改
1. 所有字段包含`Id`的修改为`ID`

## v0.1.3 2019-07-23
### 恢复
1. `AddContent`、`DeleteContent`、`AddAvatar`接口恢复使用

### 新增
1. `POST /buyInfo` 添加求购信息
2. `GET /buyInfo/:buyInfoId` 查询求购信息详情
3. `GET /buyInfo` 按条件查询求购信息
4. `PUT /content` 修改content
5. `GET /sellInfo` 支持更多查询条件

### 参数修改
1. 图片上传大小限制为5M
2. 视频上传大小限制为50M

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
