
test:
	mockgen github.com/pefish/go-core-type/api-session IApiSession > api-strategy/mock/mock-api-session/mock.go
	go build ./...
	go test -cover ./...
