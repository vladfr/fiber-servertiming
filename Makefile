publish:
	go mod tidy && go test ./...
	git push origin main --tags
	GOPROXY=proxy.golang.org go list -m github.com/vladfr/fiber-servertiming/v2@v2.0.2
