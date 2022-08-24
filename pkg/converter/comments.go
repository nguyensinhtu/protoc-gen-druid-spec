package converter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Comments map[string]string

func ParseComments(fd *descriptor.FileDescriptorProto) Comments {
	comments := make(Comments)

	for _, loc := range fd.GetSourceCodeInfo().GetLocation() {
		if !hasComment(loc) {
			continue
		}

		path := loc.GetPath()
		key := make([]string, len(path))
		for idx, p := range path {
			key[idx] = strconv.FormatInt(int64(p), 10)
		}

		comments[strings.Join(key, ".")] = buildComment(loc)
	}

	return comments
}

func (c Comments) Get(path string) string {
	if val, ok := c[path]; ok {
		fmt.Println("path ", path, ", comment: ", val)
		return val
	}

	return ""
}

func hasComment(loc *descriptor.SourceCodeInfo_Location) bool {
	if loc.GetLeadingComments() == "" && loc.GetTrailingComments() == "" {
		return false
	}

	return true
}

func buildComment(loc *descriptor.SourceCodeInfo_Location) string {
	comment := strings.TrimSpace(loc.GetLeadingComments()) + "\n\n" + strings.TrimSpace(loc.GetTrailingComments())
	return strings.Trim(comment, "\n")
}
