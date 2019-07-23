# workflow
本文档是工作流简介，我们使用git flow作为工作流，相关使用说明参考 [cheat sheet](http://danielkummer.github.io/git-flow-cheatsheet/) 。

*nix系统可以使用根目录下的`gitflow.sh`自动初始化gitflow环境。

## 分支
- master: 稳定分支. PR必须从`develop`分支合并且有4人通过review（目前为小组全体成员）
- develop: 当前开发分支，可能存在bug或不稳定新特性. PR必须有至少1人review
- feature/[name]: feature 分支
- bugfix/[name]: bugfix 分支
- hotfix/[name]: hotfix 分支

## 创建新feature
使用下面的命令创建新特性并提交远程
    
    git flow feature start [name]
    git flow feature publish [name]

示例:

    git flow feature start test
    git flow feature publish test

## 个人开发
首先checkout个人工作分支:

    git checkout -b [Github用户名] [base分支]
    
部分完成工作，提交远程:

    // git add/commit.
    git push -u origin [Github用户名]

当工作全部完成，请向[base分支]提交PR，**合并成功后你的分支会被自动删除。**

你可以使用

    git pull -p

或

    git branch -d [Github用户名]

清理你的本地分支。