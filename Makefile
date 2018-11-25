build: clean-code install test

clean-code:
	go get golang.org/x/tools/cmd/goimports && goimports -w .
	go get golang.org/x/lint/golint && golint -set_exit_status ./...

install:
	go get -v github.com/mtojek/aws-closest-region

test: install
	go get -t ./...
	go test -v ./...
	aws-closest-region
	aws-closest-region --help || test -n "$$?"
	aws-closest-region --verbose
	aws-closest-region polly
	aws-closest-region --verbose polly
	aws-closest-region harrypotter || test -n "$$?"
