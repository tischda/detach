# go-cli-template

Template for small [Go](https://www.golang.org) CLI projects.

## Get started

Name your project:
~~~
set PROJECT=my-project
~~~

Create repository and project folder with [Github CLI](https://github.com/cli/cli):
~~~
gh config set git_protocol ssh
cd %PROJECT% || gh repo create %PROJECT% --public --clone --template github.com/tischda/go-cli-template
cd %PROJECT%
init-project.cmd
~~~

Start coding.

## Release project

~~~
make test

git tag -a v1.0.0 -m "First release"
git push origin v1.0.0
make release
~~~
