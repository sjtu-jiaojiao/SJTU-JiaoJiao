# SJTU-JiaoJiao
## Dockerize
In the root of Admin/docker file folder.               
`docker build -f Dockerfile -t ng-app .`            
Then           
`docker run -d -p 8099:80     --name jojoadmin    ng-app`            
And in the 8099 port of the virtual machine address you will see the webfrontend          
All your need is to proxy pass the /api/ uri to /v1/ outside the docker inside the server by nginx Reverse Proxy
## Tools
- gitflow.sh: 自动配置gitflow.


## 后台Demo
### 用户管理
![avatar](/codes/Admin/demo/用户管理详情.JPG)
### 交易管理
![avatar](/codes/Admin/demo/交易管理详情.JPG)
### 活动管理
![avatar](/codes/Admin/demo/活动详情.JPG)
### 首页
![avatar](/codes/Admin/demo/首页.JPG)
### 统计
![avatar](/codes/Admin/demo/统计.JPG)
