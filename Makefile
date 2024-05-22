build-proto:
	rm -f order_proto/*;
	protoc \
	--go_out=order_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=order_proto \
	--go-grpc_opt=paths=source_relative \
	order.proto

run-terraform:
	terraform -chdir=infrastructure/terraform init;
	terraform -chdir=infrastructure/terraform plan;
	terraform -chdir=infrastructure/terraform apply;

run-bdd:
	docker build -f ./infrastructure/docker/Dockerfile.go_app_bdd -t hf-order-bdd:latest .;
	docker run --rm --name hf-order-bdd hf-order-bdd:latest
	@docker rmi -f hf-order-bdd >/dev/null 2>&1
	@docker rm $$(docker ps -a -f status=exited -q) -f >/dev/null 2>&1
