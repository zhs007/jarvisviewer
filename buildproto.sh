protoc -I viewerdb/proto/ viewerdb/proto/viewer.proto --go_out=plugins=grpc:viewerdb/proto