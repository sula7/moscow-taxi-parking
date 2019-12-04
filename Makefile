mod:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --remove-orphans
