# Jiaojiaoadmin

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 8.1.0.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI README](https://github.com/angular/angular-cli/blob/master/README.md).

## Dockerize
In the root of Admin/docker file folder.               
`docker build -f Dockerfile -t ng-app .`            
Then           
`docker run -d -p 8099:80     --name jojoadmin    ng-app`            
And in the 8099 port of the virtual machine address you will see the webfrontend          
All your need is to proxy pass the /api/ uri to /v1/ outside the docker inside the server by nginx Reverse Proxy. The sample nginx.conf is under the same file folder

## Tools
- gitflow.sh: 自动配置gitflow.

## Module
### 导航
Powerd by NG-ZORRO-ANTD by Ali,分为首页/管理:用户+交易/统计/活动/站点/账号:登录+注销
### 账号
Powered by NG-ALAIN by Ali, JWT格式,输入Token后自动加入请求头,点击注销清空Token。
### 首页
本视图为展示宏观数据的入口,原型拟定为总用户数/活动日程/站点信息等,非最终版本
### 管理
本视图为管理员查询用户或交易信息，并进行统一管理，详情入口，只读，按列表形式排列，分页，可以根据条件作出对应的举动（如封禁用户，终止交易），可以根据名称或标签进行筛选
### 管理详情
本视图为管理员修改用户或交易信息并查询更丰富的细节，左侧为信息，部分可写，如状态，封禁，标签，简介等。重要信息不可修改。右侧为可视化，视具体精力决定是否进行实现。
### 统计
Powered by Echarts by Baidu,本视图是企业商业智能可视化分析，左上角是展示交易关系的交易网络，可以直观看出交易频繁的团体，以及人员的核心度。中上为价格趋势图，可以展示各类交易品的价格波动情况。右上为标签词云，反应商品热度。左下为总交易热度的日历图，反应交易热度。右下为购售对比趋势图，可以研究供求关系。交互未定，具体内容未定，仅作为demo，视具体精力决定是否进行实现。
### 活动
本视图为管理员进行查询信息，并进行统一管理，详情入口，只读
### 活动详情
本视图为管理员修改活动信息并查询更丰富的细节，左侧为信息，部分可写。所有信息均可修改。右侧为可视化，视具体精力决定是否进行实现。
### 站点
本视图为管理员修改站点信息，具体内容不详

## Test
无E2E测试,使用Jasmine框架,结果如下:    
- Statements   : 90.6% ( 376/415 )
- Branches     : 75.79% ( 72/95 )
- Functions    : 92.95% ( 145/156 )
- Lines        : 90.19% ( 331/367 )
## 后台Demo
### Jaccount登录
为了减少后端工作量，目前首先使用Jaccount登录后获取Token，由管理员手动输入Token登录。JWT会进行路由守卫
![avatar](/admin-frontend/pic/J+Docker.JPG)
### 用户管理
![avatar](/admin-frontend/pic/用户管理详情.JPG)
### 交易管理
![avatar](/admin-frontend/pic/交易管理详情.JPG)
### 活动管理
![avatar](/admin-frontend/pic/活动详情.JPG)
### 首页
![avatar](/admin-frontend/pic/首页.JPG)
### 统计
![avatar](/admin-frontend/pic/统计.JPG)
