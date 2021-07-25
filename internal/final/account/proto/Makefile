GOSRC=$(GOPATH)/src

DIR=acount

proto:
	[ -e $(DIR) ]||mkdir $(DIR)
	protoc --proto_path=.:$(GOPATH)/bin:/usr/include:/usr/local/include \
		--go_out=plugins=grpc:./$(DIR) \
		"acount.proto"
