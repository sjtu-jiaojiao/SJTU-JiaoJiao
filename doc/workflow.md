# workflow
You should follow these git workflow if you want to contribute to this project.

You can use `git-flow` to automatically manage the branch. See [cheat sheet](http://danielkummer.github.io/git-flow-cheatsheet/) for more information. 

You can also run `gitflow.sh` to auto configure the git flow.

## Branch
- master: last stable and deployed version. PR must be merged from `dev` branch and approved by 4 team members(currently all teammates).
- develop: current develop version, may have bugs or incomplete features. PR must be approved by 1 team members.
- fea_[name]: feature branch.
- bug_[name]: bugfix branch.
- fix_[name]: hotfix branch.
- rel_[name]: release branch.