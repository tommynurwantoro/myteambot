dep:
	go mod tidy

pretty:
	gofmt -s -w .

build: pretty dep
	go build -o tbot app/myteambot/main.go

run: build
	./tbot

deploy:
	sudo cp tbot.service /lib/systemd/system/tbot.service

# Only for development
dev:
	go build -o tbot app/myteambot/main.go
	./tbot