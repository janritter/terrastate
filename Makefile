
prepare:
	go mod download

build: prepare
	go build -o ./bin/terrastate -v -ldflags "-X github.com/janritter/terrastate/cmd.gitSha=`git rev-parse HEAD` -X github.com/janritter/terrastate/cmd.buildTime=`date +'%Y-%m-%d_%T'` -X github.com/janritter/terrastate/cmd.version=LOCAL_BUILD"

tests:
	go test ./... -v  --cover

run:
	go run main.go --var-file ./test/test.tfvars

tests-e2e: build
	rm -rf /tmp/terrastate-e2e
	mkdir -p /tmp/terrastate-e2e
	rm -f /tmp/terrastate-e2e/terrastate.tf
	cd /tmp/terrastate-e2e && $(CURDIR)/bin/terrastate --var-file $(CURDIR)/test/test.tfvars
	terraform fmt /tmp/terrastate-e2e/terrastate.tf
	diff /tmp/terrastate-e2e/terrastate.tf $(CURDIR)/test/expected_terrastate.tf
