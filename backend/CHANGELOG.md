# 后端API更新日志
## v0.1.1 2019-07-17
### 新增
1. `GET /auth` 登陆接口添加状态码3为用户被冻结
2. `GET/PUT /user` 用户查询添加`role`字段标识用户身份

### BUG修复
1. `GET /sellInfo/:sellInfoId` 返回字段`sellInfoState`修正为`status`（文档错误实际返回正常）
2. `PUT /content` 修复`content`字段检测问题

## v0.1.0 2019-07-17
第一个稳定版本