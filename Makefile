


.PHONY: run
run:
	go install . && protoc --go_out=./api -I=. --myplugin_out=./api api/test.proto
