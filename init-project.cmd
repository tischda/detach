@echo off
setlocal

REM -- Extract go module name for from git project and replace in generator source code
for /f %%i in ('git remote get-url origin ^| sed -e "s#git@github.com:#github.com/#" -e "s#\.git$##"') do (
	sed -i "s#{{ \.GitHubRepository }}#%%i#g" template.go 
)

REM -- Generate additional files
go generate template.go

REM -- Clean up
REM del /q template.go
REM del /q init-project.cmd

endlocal
