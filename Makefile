dep:
	dep ensure -v -vendor-only

pretty:
	gofmt -s -w .

bin: pretty dep
	go build -o actblbot app/act-bl-bot/main.go

run: bin
	./actblbot

deploy:
	sudo cp actblbot.service /lib/systemd/system/actblbot.service

# Only for development
dev:
	go build -o actblbot app/act-bl-bot/main.go
	./actblbot