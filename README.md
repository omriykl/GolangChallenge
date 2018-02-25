# GolangChallenge

after installing Go export these env vars:
export GOROOT=/usr/lib/golang 

export GOPATH=$HOME/YOUR_REPO_FOLDER

export PATH=$PATH:$GOPATH/bin

install 3rd party pacakages:

go get -u github.com/valyala/quicktemplate

go get -u github.com/valyala/quicktemplate/qtc

go get github.com/gin-gonic/gin

cd to your repo directory and run "go build" 
now, to run the app "./GolangChallenge"

the app should be up an running, on port 8080

example running app can be found here: http://35.168.81.14:8080/
