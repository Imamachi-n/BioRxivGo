build:
	dep ensure -v
	go build -o "bin/batch4RSS"

run: build
	bin/batch4RSS

dev:
	go run main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock
