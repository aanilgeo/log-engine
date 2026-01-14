# Define the output directory
PROTO_DIR=proto/logging

gen-proto:
	# 1. Clean old generated files
	rm -f $(PROTO_DIR)/*.pb.go
	# 2. Generate new files directly into the correct folder
	protoc --go_out=. --go_opt=module=github.com/aanilgeo/log-engine \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/aanilgeo/log-engine \
	       proto/logging.proto
	@echo "Proto files generated successfully in $(PROTO_DIR)"