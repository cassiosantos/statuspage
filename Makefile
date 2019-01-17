dep-vendor:
	docker run -it --rm --name go-dep-vendor -e GO111MODULE=on -v `pwd`:/tmp/statuspage -w="/tmp/statuspage" golang go mod vendor

test-all:
	docker-compose up -d mongo
	for i in `find ./* -maxdepth 1 -type d -not -path "./vendor*"`; do go test $$i -coverprofile /dev/null ; done
	docker-compose down

run: 
	docker-compose down
	docker-compose up -d 
	docker-compose logs -f
