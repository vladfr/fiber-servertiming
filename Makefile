publish:
	go mod tidy && go test ./...
	git tag v0.1.0
	git push origin main --tags
	GOPROXY=proxy.golang.org go list -m github.com/vladfr/fiber-servertiming/v2@v0.2.0
