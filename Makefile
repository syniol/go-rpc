up:
	docker-compose -f ./deploy/docker/docker-compose.yml up -d


down:
	docker-compose -f ./deploy/docker/docker-compose.yml down


run:
	docker-compose -f ./deploy/docker/docker-compose.yml exec -it syniol_rpc_client sh -c "go run ./cmd/client.go"
