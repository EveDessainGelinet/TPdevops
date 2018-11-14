run:
	go run ./main.go

build:
	rm -rf ./build
	go build

release:
	GOOS=windows GOARCH=386 go build

publish: build
	heroku push origin master
