build-proto:
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
