dep:
	dep ensure -v -vendor-only

pretty:
	gofmt -s -w .

bin: pretty dep
	go build -o bot app/act-bl-bot/main.go

run: bin
	./bot

# Only for development
dev:
	go build -o bot app/act-bl-bot/main.go
	./bot