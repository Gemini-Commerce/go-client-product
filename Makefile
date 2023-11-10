version=1.0.1

generate:
	$(MAKE) -C ./modules/go-client-generator/ generate-go-client service=product version=${version}
push:
	bash git_push.sh
publish:
	GOPROXY=proxy.golang.org go list -m github.com/Gemini-Commerce/go-client-product@v${version}