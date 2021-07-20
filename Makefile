.PHONY: init
init:
	mkdir out

.PHONY: run
run:
	go install . && protoc --go_out=./out -I=. --myplugin_out=./out api/test.proto
