# workflow
You should follow these git workflow if you want to contribute to this project.

You can use `git-flow` to automatically manage the branch. See [cheat sheet](http://danielkummer.github.io/git-flow-cheatsheet/) for more information. 

You can also run `gitflow.sh` to auto configure the git flow.

## branch
- master: last stable and deployed version. PR must be merged from `develop` branch and approved by 4 team members(currently all teammates).
- develop: current develop version, may have bugs or incomplete features. PR must be approved by 1 team members.
- feature_[name]: feature branch.
- bug_[name]: bugfix branch.
- hotfix_[name]: hotfix branch.
- release_[name]: release branch.

## start a feature
This will create new feature branch and push it to remote.
    
    git flow feature start [name]
    git flow feature publish [name]

example:

    git flow feature start test
    git flow feature publish test

## start your contribution
First checkout your working branch:

    git checkout -b [your_github_name] [based_on_branch]
    
When finished, but not complete the whole work:

    // git add/commit or some operation to your branch.
    git push -u origin [your_github_name]

When finished the whole work, go github and submit a PR to [based_on_branch] with [your_github_name]. After being merged, your branch will be deleted.
You can use

    git pull -p

or

    git branch -d [your_github_name]

to tidy up your local workspace.