# vol.6

## protobuf validator

https://cipepser.hatenablog.com/entry/gRPC-validation

```shell script
# path/to/proto/user
protoc  \
  --proto_path=${GOPATH}/src \
    --proto_path=. \
    --go_out=plugins=grpc:${GOPATH}/src \
    --govalidators_out=${GOPATH}/src \
    *.proto
```
