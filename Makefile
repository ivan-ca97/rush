.PHONY: setup keys run test build update lint generate format vulncheck

format:
	go mod tidy
	go fmt ./...

keys:
	mkdir -p keys
	openssl genrsa 1024 | base64 | tr -d '\n' > keys/private.pem
	base64 --decode keys/private.pem > keys/temp_private.pem
	openssl rsa -in keys/temp_private.pem -outform PEM -pubout | base64 | tr -d '\n' > keys/public.pem
	rm keys/temp_private.pem

vulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...
