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
```cmd
syntax = "proto2";
package foo;
import "bq_table.proto";
import "bq_field.proto";

message Bar {
  option (gen_druid_spec.druid_opts).data_source = "druid_foo_data_source";

  required string id = 1 [
    (gen_druid_spec.dimention) = {
      name: "event_type" // if not set will use proto name
      create_bitmap_index: false // default is true for string type
      multi_value_handling: "SORTED_ARRAY" // default is SORTED_ARRAY for string
    }
  ];

  required string datetime_field = 2 [
    (gen_druid_spec.timestamp_spec) = {
      name: "time_field"
      format: ""
    }
  ];

  Baz baz = 3 [
    (gen_druid_schema.flatten) = {
      prefix: "baz_" // all nested fields will be flatten with prefix 'baz_'
      type: "jq"
    }
  ];
}

message Baz {
  required int32 a = 1;
}
```

Conceptually, after input data records are read, Druid applies ingestion spec components in a particular order: first flattenSpec (if any), then timestampSpec, then transformSpec, and finally dimensionsSpec and metricsSpec. Keep this in mind when writing your ingestion spec.


## Query filters
Support [Selector filter](https://druid.apache.org/docs/latest/querying/filters.html#selector-filter), [Logical expression filters](https://druid.apache.org/docs/latest/querying/filters.html#logical-expression-filters)