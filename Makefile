DRUID_PLUGIN=bin/protoc-gen-druid-spec
GO_PLUGIN=bin/protoc-gen-go
PROTOC_GEN_GO_PKG=github.com/golang/protobuf/protoc-gen-go
GLOG_PKG=github.com/golang/glog
PROTO_SRC=druid_ingestion.proto druid_spec.proto
PROTO_GENFILES=protos/druid_ingestion.pb.go protos/druid_spec.pb.go
PROTO_PKG=github.com/golang/protobuf/proto
PKGMAP=Mgoogle/protobuf/descriptor.proto=$(PROTOC_GEN_GO_PKG)/descriptor
EXAMPLES_PROTO=examples/foo-proto3.proto
GOLINT=golangci-lint run --timeout 10m

install: $(DRUID_PLUGIN)

$(DRUID_PLUGIN): $(PROTO_GENFILES) goprotobuf glog
	go build -o $@

$(PROTO_GENFILES): $(PROTO_SRC) $(GO_PLUGIN)
	protoc -I. -Ivendor/protobuf --plugin=$(GO_PLUGIN) --go_out=$(PKGMAP):protos $(PROTO_SRC)

goprotobuf:
	go get $(PROTO_PKG)

glog:
	go get $(GLOG_PKG)

$(GO_PLUGIN):
	go get $(PROTOC_GEN_GO_PKG)
	go build -o $@ $(PROTOC_GEN_GO_PKG)

test: $(PROTO_SRC)
	go test ./pkg/converter/... -v

distclean clean:
	go clean
	rm -f $(GO_PLUGIN) $(DRUID_PLUGIN)

realclean: distclean
	rm -f $(PROTO_GENFILES)

examples: $(DRUID_PLUGIN)
	protoc -I. -Ivendor/protobuf --plugin=$(DRUID_PLUGIN) --druid-spec_out=examples $(EXAMPLES_PROTO)

lint:
	$(GOLINT) -v ./...

.PHONY: goprotobuf glog lint