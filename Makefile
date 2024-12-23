# go makefile

program != basename $$(pwd)

latest_release != gh release list --json tagName --jq '.[0].tagName' | tr -d v
version != cat VERSION

gitclean = if git status --porcelain | grep '^.*$$'; then echo git status is dirty; false; else echo git status is clean; true; fi

build: fmt
	fix go build

fmt: go.sum
	fix go fmt . ./...

go.mod:
	go mod init

go.sum: go.mod
	go mod tidy

install: build
	go install

test:
	fix -- go test -failfast -v .
	fix -- go test -failfast -v ./...

release:
	@$(gitclean) || { [ -n "$(dirty)" ] && echo "allowing dirty release"; }
	@$(if $(update),gh release delete -y v$(version),)
	gh release create v$(version) --notes "v$(version)"

testclean:
	rm -f testdata/*.out
	rm -f testdata/*.err

clean: testclean
	rm -f $(program)
	go clean

sterile: clean
	go clean -r
	go clean -cache
	go clean -modcache
	rm -f go.mod go.sum
