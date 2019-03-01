
prepare:
	dep ensure -v

build: prepare
	go build -o ./bin/terrastate -v -ldflags "-X github.com/janritter/terrastate/cmd.gitSha=`git rev-parse HEAD` -X github.com/janritter/terrastate/cmd.buildTime=`date +'%Y-%m-%d_%T'` -X github.com/janritter/terrastate/cmd.version=LOCAL_BUILD"

tests:
	go test ./... -v

run:
	go run main.go