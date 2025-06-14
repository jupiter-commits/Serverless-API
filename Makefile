build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go
deploy_testing: build
	serverless deploy --stage testing