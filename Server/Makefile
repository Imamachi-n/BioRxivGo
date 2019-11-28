build:
	dep ensure -v
	go build -o "bin/goServer"

run: build
	bin/goServer

dev:
	go run main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

setup:
	go get -u github.com/golang/dep/cmd/dep

docker: setup build

dbuild:
	docker build . -t biorxivgoback:dev
