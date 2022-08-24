package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/nguyensinhtu/protoc-gen-druid-spec/pkg/converter"
	"google.golang.org/protobuf/proto"
)

func main() {
	flag.Parse()
	ok := true
	glog.Info("Processing code generator request")
	res, err := converter.ConvertFrom(os.Stdin)
	if err != nil {
		ok = false
		if res == nil {
			message := fmt.Sprintf("Failed to read input: %v", err)
			res = &plugin.CodeGeneratorResponse{
				Error: &message,
			}
		}
	}

	glog.Info("Serializing code generator response")
	data, err := proto.Marshal(res)
	if err != nil {
		glog.Fatal("Cannot marshal response", err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		glog.Fatal("Failed to write response", err)
	}

	if ok {
		glog.Info("Succeeded to process code generator request")
	} else {
		glog.Info("Failed to process code generator but successfully sent the error to protoc")
	}
}
