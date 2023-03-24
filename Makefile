build:
	go mod init app && go get github.com/sujjadapp/videoproxy/route && GO111MODULE=auto CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o videoproxy -a -ldflags "-s -w" main.go && docker run -d -m 100m \ --name videoproxy \ --restart always \ -p 0.0.0.0:8080:6060 \ e YOUTUBE_API_KEY="yourkey" \ sujjadapp/videoproxy

docker:
	make build && \
	docker build -t=sujjadapp/videoproxy .
