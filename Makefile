build:
	go mod init main && go get github.com/sujjadapp/videoproxy/route && GO111MODULE=auto CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o videoproxy -a -ldflags "-s -w" main.go

docker:
	make build && \
	docker build -t=sujjadapp/videoproxy .
