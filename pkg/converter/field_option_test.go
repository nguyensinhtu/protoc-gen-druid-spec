package converter

import (
	"testing"
)

func TestIgnore(t *testing.T) {
	testConvert(t, `
		file_to_generate: "foo.proto"
		proto_file <
			name: "foo.proto"
			package: "example_package"
			message_type <
				name: "FooProto"
				field <
					name: "i1"
					number: 1
					type: TYPE_INT32
					label: LABEL_OPTIONAL
				>
				field <
					name: "i2"
					number: 2
					type: TYPE_INT32
					label: LABEL_OPTIONAL
					options <
						[gen_druid_spec.spec] <
							ignore: true
						>
					>
				>
				options <
					[gen_druid_spec.druid_opts]: <
					data_source_name: "foo_table"
					>
				>
			>
		>
	`, map[string]string{
		"example_package/foo_table.ingestion": `{
			"spec": {
				"dataSchema": {
					"dataSource": "foo_table",
					"dimensionsSpec": {
						"dimensions": [
							{
								"type": "long",
								"name": "i1"
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

func TestMetric(t *testing.T) {
	testConvert(t, `
	file_to_generate: "foo.proto"
	proto_file <
		name: "foo.proto"
		package: "example_package"
		message_type <
			name: "FooProto"
			field <
				name: "i1"
				number: 1
				type: TYPE_INT32
				label: LABEL_OPTIONAL
			>
			field <
				name: "i2"
				number: 2
				type: TYPE_STRING
				label: LABEL_OPTIONAL
				options <
					[gen_druid_spec.spec] <
						metric <
							metric_name: "i2_theta_sketch"
							size: 16384
						>
					>
				>
			>
			options <
				[gen_druid_spec.druid_opts]: <
				data_source_name: "foo_table"
				>
			>
		>
	>
	`, map[string]string{
		"example_package/foo_table.ingestion": `{
			"spec": {
				"dataSchema": {
					"dataSource": "foo_table",
					"dimensionsSpec": {
						"dimensions": [
							{
								"type": "long",
								"name": "i1"
							},
							{
								"type": "string",
								"name": "i2",
								"multiValueHandling": "SORTED_ARRAY",
								"createBitmapIndex": true
							}
						]
					},
					"metricsSpec": [
						{
							"type": "thetaSketch",
							"name": "i2_theta_sketch",
							"fieldName": "i2",
							"size": 16384,
							"isInputThetaSketch": false
						}
					] 
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

func TestFlatten(t *testing.T) {
	testConvert(t, `
	file_to_generate: "foo.proto"
	proto_file <
		name: "foo.proto"
		package: "example_package"
		message_type <
			name: "FooProto"
			field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
			field < 
				name: "bar" 
				number: 2 
				type: TYPE_MESSAGE 
				label: LABEL_OPTIONAL 
				type_name: "BarProto" 
				options <
					[gen_druid_spec.spec] <
						flatten <
							prefix: "baz_"
						>
					>
				>
			>
			options < [gen_druid_spec.druid_opts] <data_source_name: "foo_table"> >
		>
		message_type <
			name: "BarProto"
			field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
			field < name: "i2" number: 1 type: TYPE_STRING label: LABEL_OPTIONAL >
			field <
				name: "zoo_in_bar" 
				number: 2 
				type: TYPE_MESSAGE 
				label: LABEL_OPTIONAL 
				type_name: "ZooProto" 
				options <
					[gen_druid_spec.spec] <
						flatten <
						>
					>
				>
			>
		>

		message_type < 
			name: "ZooProto"
			field < name: "name" number: 1 type: TYPE_STRING label: LABEL_OPTIONAL > 
			field < name: "no" number: 2 type: TYPE_INT64 label: LABEL_OPTIONAL >
			field < name: "dump" number: 3 type: TYPE_INT64 label: LABEL_OPTIONAL 
				options < [gen_druid_spec.spec] < ignore: true > >	
			> 
		>
	>`, map[string]string{
		"example_package/foo_table.ingestion": `{
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
								"name": "baz__i1",
								"type": "long"
							},
							{
								"createBitmapIndex": true,
								"multiValueHandling": "SORTED_ARRAY",
								"name": "baz__i2",
								"type": "string"
							},
							{
								"createBitmapIndex": true,
								"multiValueHandling": "SORTED_ARRAY",
								"name": "baz__zoo_in_bar__name",
								"type": "string"
							},
							{
								"name": "baz__zoo_in_bar__no",
								"type": "long"
							}
						]
					}
				},
				"ioConfig": {
					"inputFormat": {
						"flattenSpec": {
							"fields": [
								{
									"expr": ".bar.i1",
									"name": "baz__i1",
									"type": "jq"
								},
								{
									"expr": ".bar.i2",
									"name": "baz__i2",
									"type": "jq"
								},
								{
									"expr": ".bar.zoo_in_bar.name",
									"name": "baz__zoo_in_bar__name",
									"type": "jq"
								},
								{
									"expr": ".bar.zoo_in_bar.no",
									"name": "baz__zoo_in_bar__no",
									"type": "jq"
								}
							],
							"useFieldDiscovery": false
						},
						"type": "json"
					}
				}
			}
		}`,
	})
}

func TestDimension(t *testing.T) {
	testConvert(t, `
		file_to_generate: "foo.proto"
		proto_file <
			name: "foo.proto"
			package: "example_package"
			message_type <
				name: "FooProto"
				field <
					name: "i1"
					number: 1
					type: TYPE_STRING
					label: LABEL_OPTIONAL
					options <
						[gen_druid_spec.spec] <
						dimension < multi_value_handling: "SORTED_SET" >
						>
					>
				>
				options < [gen_druid_spec.druid_opts] <data_source_name: "foo_table"> > 
			>
		>
	`, map[string]string{
		"example_package/foo_table.ingestion": `{
			"spec": {
				"dataSchema": {
					"dataSource": "foo_table",
					"dimensionsSpec": {
						"dimensions": [
							{
								"name": "i1",
								"type": "string",
								"multiValueHandling": "SORTED_SET",
								"createBitmapIndex": true
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

func TestIgnoreNested(t *testing.T) {
	testConvert(t, `
	file_to_generate: "foo.proto"
	proto_file <
		name: "foo.proto"
		package: "example_package"
		message_type <
			name: "FooProto"
			field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
			field < 
				name: "bar" 
				number: 2 
				type: TYPE_MESSAGE 
				label: LABEL_OPTIONAL 
				type_name: "BarProto" 
				options <
					[gen_druid_spec.spec] <
						flatten <
							prefix: "baz_"
						>
					>
				>
			>
			options < [gen_druid_spec.druid_opts] <data_source_name: "foo_table"> >
		>
		message_type <
			name: "BarProto"
			field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL 
				options < [gen_druid_spec.spec] < ignore: true > > 
			>
			field <
				name: "zoo_in_bar" 
				number: 2 
				type: TYPE_MESSAGE 
				label: LABEL_OPTIONAL 
				type_name: "ZooProto" 
				options <
					[gen_druid_spec.spec] <
						flatten <
						>
					>
				>
			>
		>

		message_type < 
			name: "ZooProto"
			field < name: "dump" number: 3 type: TYPE_INT64 label: LABEL_OPTIONAL 
				options < [gen_druid_spec.spec] < ignore: true > >	
			> 
		>
	>`, map[string]string{
		"example_package/foo_table.ingestion": `{
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
