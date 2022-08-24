package converter

import (
	"encoding/json"
	"reflect"
	"testing"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/encoding/prototext"
)

type schema map[string]interface{}

func joinNames(targets map[string]*schema) (result string) {
	sep := ""
	for name := range targets {
		result += sep + name
		sep = ", "
	}
	return
}

func testConvert(t *testing.T, input string, expectedOutputs map[string]string, extras ...func(request *plugin.CodeGeneratorRequest)) {
	req := plugin.CodeGeneratorRequest{}
	if err := prototext.Unmarshal([]byte(input), &req); err != nil {
		t.Fatal("Failed to parse test input: ", err)
	}

	// apply custom transformations, if any
	for _, extra := range extras {
		extra(&req)
	}

	expectedSchema := make(map[string]*schema)
	for filename, data := range expectedOutputs {
		parsed := new(schema)
		if err := json.Unmarshal([]byte(data), parsed); err != nil {
			t.Fatalf("Failed to parse an expectation: %s: %v", data, err)
		}
		expectedSchema[filename] = parsed
	}

	res, err := Convert(&req)
	if err != nil {
		t.Fatal("Conversion failed. ", err)
	}
	if res.Error != nil {
		t.Fatal("Conversion failed. ", res.Error)
	}

	actualSchema := make(map[string]*schema)
	for _, file := range res.GetFile() {
		s := &schema{}
		// fmt.Println("content res: ", file.GetContent())
		if err := json.Unmarshal([]byte(file.GetContent()), s); err != nil {
			t.Fatalf("Expected to be a valid JSON, but wasn't %s: %v", file.GetContent(), err)
		}
		actualSchema[file.GetName()] = s
	}

	if len(actualSchema) != len(expectedSchema) {
		t.Errorf("Expected %d files generated, but actually %d files:\nExpectation: %s\n Actual: %s",
			len(expectedSchema), len(actualSchema), joinNames(expectedSchema), joinNames(actualSchema))
	}

	for name, actual := range actualSchema {
		expected, ok := expectedSchema[name]
		if !ok {
			t.Error("Unexpected file generated: ", name)
		}
		if !reflect.DeepEqual(expected, actual) {
			expectedJson, err := json.Marshal(expected)
			if err != nil {
				t.Error("Failed to parse expected ", err)
			}

			actualJson, _ := json.Marshal(actual)
			t.Errorf("Expected the content of %s to be \"%v\" but got \"%v\"", name, string(expectedJson), string(actualJson))
		}
	}
}

func TestIgnoreNonTargetMessage(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.nested"
				message_type <
					name: "FooProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
				>
				message_type <
					name: "BarProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					options < [gen_druid_spec.druid_opts] <data_source_name: "bar_table"> >
				>
				message_type <
					name: "BazProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
				>
			>
		`,
		map[string]string{
			"example_package/nested/bar_table.ingestion": `{
				"spec": {
					"dataSchema": {
						"dataSource": "bar_table",
						"dimensionsSpec": {
							"dimensions": [
								{
									"name": "i1",
									"type": "long"
								}
							]
						}
					},
					"ioConfig": {
						"inputFormat": {
							"type": "json"
						}
					}
				}
			}`,
		})
}

// TestIgnoreNonTargetFile checks if the generator ignores messages in non target files.
func TestIgnoreNonTargetFile(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.nested"
				message_type <
					name: "FooProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					options < [gen_druid_spec.druid_opts] <data_source_name: "foo_table"> >
				>
			>
			proto_file <
				name: "bar.proto"
				package: "example_package.nested"
				message_type <
					name: "BarProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					options < [gen_druid_spec.druid_opts] <data_source_name: "bar_table"> >
				>
			>
		`,
		map[string]string{
			"example_package/nested/foo_table.ingestion": `{
				"spec": {
					"dataSchema": {
						"dataSource": "foo_table",
						"dimensionsSpec": {
							"dimensions": [
								{
									"name": "i1",
									"type": "long"
								}
							]
						}
					},
					"ioConfig": {
						"inputFormat": {
							"type": "json"
						}
					}
				}
			}`,
		})
}

func TestStopsAtRecursiveMessage(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.recursive"
				message_type <
					name: "FooProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field <
                        name: "bar" number: 2 type: TYPE_MESSAGE label: LABEL_OPTIONAL
                        type_name: "BarProto" >
					options < [gen_druid_spec.druid_opts] <data_source_name: "foo_table"> >
				>
				message_type <
					name: "BarProto"
					field < name: "i2" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field <
                        name: "foo" number: 2 type: TYPE_MESSAGE label: LABEL_OPTIONAL
                        type_name: "FooProto" >
				>
			>
		`,
		map[string]string{
			"example_package/recursive/foo_table.ingestion": `{
				"spec": {
					"dataSchema": {
						"dataSource": "foo_table",
						"dimensionsSpec": {
							"dimensions": [
								{
									"name": "i1",
									"type": "long"
								},
								{
									"name": "bar",
									"type": "string"
								}
							]
						}
					},
					"ioConfig": {
						"inputFormat": {
							"type": "json"
						}
					}
				}
			}`,
		})
}

func TestTypes(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.nested"
				message_type <
					name: "FooProto"
					field < name: "i32" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "i64" number: 2 type: TYPE_INT64 label: LABEL_OPTIONAL >
					field < name: "ui32" number: 3 type: TYPE_UINT32 label: LABEL_OPTIONAL >
					field < name: "ui64" number: 4 type: TYPE_UINT64 label: LABEL_OPTIONAL >
					field < name: "si32" number: 5 type: TYPE_SINT32 label: LABEL_OPTIONAL >
					field < name: "si64" number: 6 type: TYPE_SINT64 label: LABEL_OPTIONAL >
					field < name: "ufi32" number: 7 type: TYPE_FIXED32 label: LABEL_OPTIONAL >
					field < name: "ufi64" number: 8 type: TYPE_FIXED64 label: LABEL_OPTIONAL >
					field < name: "sfi32" number: 9 type: TYPE_SFIXED32 label: LABEL_OPTIONAL >
					field < name: "sfi64" number: 10 type: TYPE_SFIXED64 label: LABEL_OPTIONAL >
					field < name: "d" number: 11 type: TYPE_DOUBLE label: LABEL_OPTIONAL >
					field < name: "f" number: 12 type: TYPE_FLOAT label: LABEL_OPTIONAL >
					field < name: "bool" number: 16 type: TYPE_BOOL label: LABEL_OPTIONAL >
					field < name: "str" number: 13 type: TYPE_STRING label: LABEL_OPTIONAL >
					field < name: "bytes" number: 14 type: TYPE_BYTES label: LABEL_OPTIONAL >
					field <
						name: "enum1" number: 15 type: TYPE_ENUM label: LABEL_OPTIONAL
						type_name: ".example_package.nested.FooProto.Enum1"
					>
					field <
						name: "enum2" number: 16 type: TYPE_ENUM label: LABEL_OPTIONAL
						type_name: "FooProto.Enum1"
					>
					field <
						name: "grp1" number: 17 type: TYPE_GROUP label: LABEL_OPTIONAL
						type_name: ".example_package.nested.FooProto.Group1"
					>
					field <
						name: "grp2" number: 18 type: TYPE_GROUP label: LABEL_OPTIONAL
						type_name: "FooProto.Group1"
					>
					field <
						name: "msg1" number: 19 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".example_package.nested.FooProto.Nested1"
					>
					field <
						name: "msg2" number: 20 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: "FooProto.Nested1"
					>
					field <
						name: "msg3" number: 21 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".example_package.nested2.BarProto"
					>
					field <
						name: "msg4" number: 22 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: "nested2.BarProto"
					>
					field <
						name: "arr" number: 23 type: TYPE_INT32 label: LABEL_REPEATED
					>
					field <
						name: "t" number: 11 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.Timestamp"
					>
					nested_type <
						name: "Group1"
						field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					>
					nested_type <
						name: "Nested1"
						field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					>
					nested_type <
						name: "EmptyNested1"
					>
					enum_type < name: "Enum1" value < name: "E1" number: 1 > value < name: "E2" number: 2 > >
					options < [gen_druid_spec.druid_opts] <data_source_name: "foo_table"> >
				>
			>
			proto_file <
				name: "bar.proto"
				package: "example_package.nested2"
				message_type <
					name: "BarProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "i2" number: 2 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "i3" number: 3 type: TYPE_INT32 label: LABEL_OPTIONAL >
				>
			>
		`,
		map[string]string{
			"example_package/nested/foo_table.ingestion": `{
				"spec": {
					"dataSchema": {
						"dataSource": "foo_table",
						"dimensionsSpec": {
							"dimensions": [
								{
									"name": "i32",
									"type": "long"
								},
								{
									"name": "i64",
									"type": "long"
								},
								{
									"name": "ui32",
									"type": "long"
								},
								{
									"name": "ui64",
									"type": "long"
								},
								{
									"name": "si32",
									"type": "long"
								},
								{
									"name": "si64",
									"type": "long"
								},
								{
									"name": "ufi32",
									"type": "long"
								},
								{
									"name": "ufi64",
									"type": "long"
								},
								{
									"name": "sfi32",
									"type": "long"
								},
								{
									"name": "sfi64",
									"type": "long"
								},
								{
									"name": "d",
									"type": "double"
								},
								{
									"name": "f",
									"type": "double"
								},
								{
									"name": "bool",
									"type": "long"
								},
								{
									"name": "str",
									"type": "string",
									"createBitmapIndex": true,
                  "multiValueHandling": "SORTED_ARRAY"
								},
								{
									"name": "bytes",
									"type": "string"
								},
								{
									"name": "enum1",
									"type": "string"
								},
								{
									"name": "enum2",
									"type": "string"
								},
								{
									"name": "grp1",
									"type": "string"
								},
								{
									"name": "grp2",
									"type": "string"
								},
								{
									"name": "msg1",
									"type": "string"
								},
								{
									"name": "msg2",
									"type": "string"
								},
								{
									"name": "msg3",
									"type": "string"
								},
								{
									"name": "msg4",
									"type": "string"
								},
								{
									"name": "arr",
									"type": "string",
									"createBitmapIndex": true,
									"multiValueHandling": "SORTED_ARRAY"
								},
								{
									"name": "t",
									"type": "string"
								}
							]
						}
					},
					"ioConfig": {
						"inputFormat": {
							"type": "json"
						}
					}
				}
			}`,
		})
}

// TestWellKnownTypes tests the generator with various well-known message types
// which have custom JSON serialization.
func TestWellKnownTypes(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package"
				message_type <
					name: "FooProto"
					field <
						name: "i32" number: 1 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.Int32Value"
					>
					field <
						name: "i64" number: 2 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.Int64Value"
					>
					field <
						name: "ui32" number: 3 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.UInt32Value"
					>
					field <
						name: "ui64" number: 4 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.UInt64Value"
					>
					field <
						name: "d" number: 5 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.DoubleValue"
					>
					field <
						name: "f" number: 6 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.FloatValue"
					>
					field <
						name: "bool" number: 7 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.BoolValue"
					>
					field <
						name: "str" number: 8 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.StringValue"
					>
					field <
						name: "bytes" number: 9 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.BytesValue"
					>
					field <
						name: "du" number: 10 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.Duration"
					>
					field <
						name: "t" number: 11 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.Timestamp"
					>
					options < [gen_druid_spec.druid_opts] <data_source_name: "foo_table"> >
				>
			>
		`,
		map[string]string{
			"example_package/foo_table.ingestion": `{
				"spec": {
					"dataSchema": {
						"dataSource": "foo_table",
						"dimensionsSpec": {
							"dimensions": [
								{
									"name": "i32",
									"type": "long"
								},
								{
									"name": "i64",
									"type": "long"
								},
								{
									"name": "ui32",
									"type": "long"
								},
								{
									"name": "ui64",
									"type": "long"
								},
								{
									"name": "d",
									"type": "double"
								},
								{
									"name": "f",
									"type": "double"
								},
								{
									"name": "bool",
									"type": "long"
								},
								{
									"createBitmapIndex": true,
                  "multiValueHandling": "SORTED_ARRAY",
									"name": "str",
									"type": "string"
								},
								{
									"name": "bytes",
									"type": "string"
								},
								{
									"name": "du",
									"type": "string"
								},
								{
									"name": "t",
									"type": "string"
								}
							]
						}
					},
					"ioConfig": {
						"inputFormat": {
							"type": "json"
						}
					}
				}
			}`,
		})
}
