run:
	go run cmd/api/main.go  

install:
	go mod tidy
	go mod download

docker-build:
	docker build -f build/Dockerfile -t talanakombat .

docker-run:
	docker run -p 8080:8080 --env-file=.env talanakombat  
