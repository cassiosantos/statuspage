dep-vendor:
	docker run -it --rm --name go-dep-vendor -e GO111MODULE=on -v `pwd`:/tmp/statuspage -w="/tmp/statuspage" golang go mod vendor

test-all: dep-vendor
	docker-compose down
	docker-compose up -d mongo
	for i in `find ./* -maxdepth 1 -type d -not -path "./vendor*" -not -path "./docs*"`; do go test $$i -coverprofile /dev/null ; done
	docker-compose down

build:
	docker-compose build statuspage

run: 
	docker-compose down
	docker-compose up -d 
	docker-compose logs -f
