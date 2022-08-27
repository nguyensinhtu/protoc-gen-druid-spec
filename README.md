## protoc-gen-druid-spec

# Installation
```cmd
go install github.com/nguyensinhtu/protoc-gen-druid-spec@latest
```

# Usage
protoc --druid-spect_out=path/to/outdir foo.proto
```cmd
protoc --druid-spec_out=path/to/out/dir foo.proto --proto_path=. --proto_path=<path_to_google_proto_folder>/src

```

# Example
```protobuf
syntax = "proto3";

package foo;

import "druid_ingestion.proto";
import "druid_spec.proto";

import "google/protobuf/wrappers.proto";
import "google/protobuf/timestamp.proto";

message Bar {
  option (gen_druid_spec.druid_opts) = {
    data_source_name: "bar_proto3_table"
    segment_granularity: "day"
    query_granularity: "day"
  };

  string client_id = 1 [ (gen_druid_spec.spec) = {
    dimension : {
      multi_value_handling : "SORTED_SET"
      create_bitmap_index : true
    }
    metric : {
      metric_name : "client_id_sketch"
      size : 16384
      type : "thetaSketch"
    }
  } ];

  Baz baz = 2 
    [ (gen_druid_spec.spec).flatten = {prefix : "baz_" } ];;

  TimeSpec time_spec = 4[ (gen_druid_spec.spec).flatten = { prefix : "time_spec_" } ];
}

message Baz {
  int32 a = 1 [ (gen_druid_spec.spec).metric = {
    metric_name: "a_metric"
    size: 16384 
    type: "thetaSketch"
  }];
}

message TimeSpec {
  string date_key = 1[ (gen_druid_spec.spec).timestamp = {} ];
}

```
Output
```json
{
 "spec": {
  "dataSchema": {
   "dataSource": "bar_proto3_table",
   "timestampSpec": {
    "column": "time_spec__date_key"
   },
   "dimensionsSpec": {
    "dimensions": [
     {
      "name": "client_id",
      "type": "string",
      "multiValueHandling": "SORTED_SET",
      "createBitmapIndex": true
     },
     {
      "name": "baz__a",
      "type": "long"
     }
    ],
    "dimensionExclusions": [
     "time_spec__date_key"
    ]
   },
   "metricsSpec": [
    {
     "name": "client_id_sketch",
     "type": "thetaSketch",
     "isInputThetaSketch": false,
     "fieldName": "client_id",
     "size": 16384
    },
    {
     "name": "a_metric",
     "type": "thetaSketch",
     "isInputThetaSketch": false,
     "fieldName": "baz__a",
     "size": 16384
    }
   ],
   "granularitySpec": {
    "type": "uniform",
    "segmentGranularity": "day",
    "queryGranularity": "day",
    "rollup": true,
    "intervals": []
   }
  },
  "ioConfig": {
   "inputFormat": {
    "type": "json",
    "flattenSpec": {
     "fields": [
      {
       "type": "jq",
       "name": "baz__a",
       "expr": ".baz.a"
      },
      {
       "type": "jq",
       "name": "time_spec__date_key",
       "expr": ".time_spec.date_key"
      }
     ],
     "useFieldDiscovery": false
    }
   }
  }
 }
}
```

Conceptually, after input data records are read, Druid applies ingestion spec components in a particular order: first flattenSpec (if any), then timestampSpec, then transformSpec, and finally dimensionsSpec and metricsSpec. Keep this in mind when writing your ingestion spec.

## flattenSpec
Parent field separate with nested field by double underscore '`__`'
- if you dont set prefix, default will be protobuf field name
```protobuf
message Bar {
   Foo foo = 4[ (gen_druid_spec.spec).flatten = {} ];
}

message Foo {
  int32 i1 = 1;
}
```

Output
```json
{
  "fields": [
    {
      "type": "jq",
      "name": "foo__i1",
      "expr": ".foo.i1"
    }
  ],
  "useFieldDiscovery": false
}
```
- if you want to remove parent name set empty prefix  
```protobuf
message Bar {
   Foo foo = 4[ (gen_druid_spec.spec).flatten = {prefix: ""} ];
}

message Foo {
  int32 i1 = 1;
}
```
Output
```json
{
  "fields": [
    {
      "type": "jq",
      "name": "i1",
      "expr": ".foo.i1"
    }
  ],
  "useFieldDiscovery": false
}
```

- if you want to use diffirent prefix name
```protobuf
message Bar {
   Foo foo = 4[ (gen_druid_spec.spec).flatten = {prefix: "my_custom_name"} ];
}

message Foo {
  int32 i1 = 1;
}
```
Output
```json
{
  "fields": [
    {
      "type": "jq",
      "name": "my_custom_name__i1",
      "expr": ".foo.i1"
    }
  ],
  "useFieldDiscovery": false
}
```


## Query filters
Support [Selector filter](https://druid.apache.org/docs/latest/querying/filters.html#selector-filter), [Logical expression filters](https://druid.apache.org/docs/latest/querying/filters.html#logical-expression-filters)