clean:
	rm -rf ./bin
	rm coverage.out

build:
	go build ./...

clean.build: clean build

fmt:
	gofmt -w ./

doc:
	godoc -http=:6060

test:
	go clean -testcache
	go test ./... -cover -coverprofile=coverage.out

test.v:
	go clean -testcache
	go test ./... -cover -coverprofile=coverage.out -v

test.cov:
	go tool cover -html=coverage.out

test.bench:
	go clean -testcache
	go test -bench=. -benchmem

test.bench.profile:
	go clean -testcache
	go test -bench=. -benchmem -memprofile memory.out

publish:
	git tag $(tag)
	git push origin $(tag)
	go list -m github.com/thegogod/rum@$(tag)

.PHONY: test
