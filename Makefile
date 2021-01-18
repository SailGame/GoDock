.PHONY: dock test proto generate clean

dock:
	go build -o build/dock

test:
	go test -cpu 1,4 -timeout 7m github.com/SailGame/GoDock/...

proto:
	mkdir -p pb
	protoc -I proto proto/core/*.proto --go_out=pb/ --go_opt=paths=source_relative --go-grpc_out=pb/ --go-grpc_opt=paths=source_relative

generate:
	go generate ./...

clean:
	rm -rf build/* pb/