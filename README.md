gli
====

GitLab Commandline Interface

## Synopsis

* show project list

```
$ gli projects
```

* show project issues

```
$ gli issues {PROJECT_ID}
```

* checkout issue (git-flow) with peco example

```
if exists peco; then

    function peco-issue-checkout () {
        local branch_name=$(gli projects | grep $(git remote -v | grep push | awk {'print $2'} | sed -e "s#ssh://##g" | sed -e "s#/#:#") | awk {'print $1'} | sed -e "s/#//" | xargs gli issues | peco | awk {'print "issue/" $1'} | sed -e "s/#//")
        if [ -n "$branch_name" ]; then
            BUFFER="git flow feature start ${branch_name}"
            zle accept-line
        fi
        zle clear-screen
    }

    zle -N peco-issue-checkout
    bindkey '^T' peco-issue-checkout

fi
```

## Usage

## Author

[Yusuke Hatanaka](http://twitter.com/Hatajoe)
