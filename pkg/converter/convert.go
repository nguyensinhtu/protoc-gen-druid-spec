package converter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/nguyensinhtu/protoc-gen-druid-spec/protos"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

var (
	typeFromWKT = map[string]string{
		".google.protobuf.Int32Value":  "long",
		".google.protobuf.Int64Value":  "long",
		".google.protobuf.UInt32Value": "long",
		".google.protobuf.UInt64Value": "long",
		".google.protobuf.DoubleValue": "double",
		".google.protobuf.FloatValue":  "double",
		".google.protobuf.BoolValue":   "long",
		".google.protobuf.StringValue": "string",
		".google.protobuf.BytesValue":  "string",
		".google.protobuf.Duration":    "string",
		".google.protobuf.Timestamp":   "string",
	}
	typeFromFieldType = map[descriptor.FieldDescriptorProto_Type]string{
		descriptor.FieldDescriptorProto_TYPE_DOUBLE: "double",
		descriptor.FieldDescriptorProto_TYPE_FLOAT:  "double",

		descriptor.FieldDescriptorProto_TYPE_INT64:    "long",
		descriptor.FieldDescriptorProto_TYPE_UINT64:   "long",
		descriptor.FieldDescriptorProto_TYPE_INT32:    "long",
		descriptor.FieldDescriptorProto_TYPE_UINT32:   "long",
		descriptor.FieldDescriptorProto_TYPE_FIXED64:  "long",
		descriptor.FieldDescriptorProto_TYPE_FIXED32:  "long",
		descriptor.FieldDescriptorProto_TYPE_SFIXED32: "long",
		descriptor.FieldDescriptorProto_TYPE_SFIXED64: "long",
		descriptor.FieldDescriptorProto_TYPE_SINT32:   "long",
		descriptor.FieldDescriptorProto_TYPE_SINT64:   "long",

		descriptor.FieldDescriptorProto_TYPE_STRING: "string",
		descriptor.FieldDescriptorProto_TYPE_BYTES:  "string",
		descriptor.FieldDescriptorProto_TYPE_ENUM:   "string",

		descriptor.FieldDescriptorProto_TYPE_BOOL: "long",

		descriptor.FieldDescriptorProto_TYPE_GROUP:   "record",
		descriptor.FieldDescriptorProto_TYPE_MESSAGE: "record",
	}

	modeFromFieldLabel = map[descriptor.FieldDescriptorProto_Label]string{
		descriptor.FieldDescriptorProto_LABEL_OPTIONAL: "NULLABLE",
		descriptor.FieldDescriptorProto_LABEL_REQUIRED: "REQUIRED",
		descriptor.FieldDescriptorProto_LABEL_REPEATED: "REPEATED",
	}

	supportedMultiValueHanldingOpts = map[string]bool{
		"SORTED_SET":   true,
		"ARRAY":        true,
		"SORTED_ARRAY": true,
	}

	supportedMetricAggregators = map[string]bool{
		"count":       true,
		"longSum":     true,
		"floatSum":    true,
		"doubleSum":   true,
		"doubleMin":   true,
		"doubleMax":   true,
		"floatMin":    true,
		"floatMax":    true,
		"longMin":     true,
		"longMax":     true,
		"doubleMean":  true,
		"thetaSketch": true,
	}

	supportedGranularities = map[string]bool{
		"none": true, "all": true, "second": true, "minute": true, "fifteen_minute": true, "thirty_minute": true,
		"day": true, "week": true, "month": true, "quarter": true, "year": true,
	}

	supportedIngestionType = map[string]bool{
		"index_parallel": true,
		"index":          true,
		"index_hadoop":   true,
		"kafka":          true,
	}
)

type DimensionField struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	MultiValueHandling string `json:"multiValueHandling,omitempty"`
	CreateBitmapIndex  bool   `json:"createBitmapIndex,omitempty"`
}

type DimensionSpec struct {
	Dimensions          []*DimensionField `json:"dimensions,omitempty"`
	DimensionExclusions []string          `json:"dimensionExclusions,omitempty"`
}

func (spec *DimensionSpec) merge(other *DimensionSpec) {
	if spec.Dimensions == nil {
		spec.Dimensions = other.Dimensions
	} else if other.Dimensions != nil {
		spec.Dimensions = append(spec.Dimensions, other.Dimensions...)
	}
	if spec.DimensionExclusions == nil {
		spec.DimensionExclusions = other.DimensionExclusions
	} else if other.DimensionExclusions != nil {
		spec.DimensionExclusions = append(spec.DimensionExclusions, other.DimensionExclusions...)
	}
}

type TimestampField struct {
	FieldName    string `json:"column,omitempty"`
	Format       string `json:"format,omitempty"`
	MissingValue string `json:"missingValue,omitempty"`
}

type MetricField struct {
	MetricFieldName    string `json:"name,omitempty"`
	ApproxType         string `json:"type,omitempty"`
	IsInputThetaSketch bool   `json:"isInputThetaSketch"`
	FieldName          string `json:"fieldName,omitempty"`
	Size               int64  `json:"size,omitempty"`
}

type FlattenField struct {
	Type       string `json:"type,omitempty"`
	Name       string `json:"name,omitempty"`
	Expression string `json:"expr,omitempty"`
}

type FlattenSpec struct {
	FlattenFields     []*FlattenField `json:"fields,omitempty"`
	UseFieldDiscovery bool            `json:"useFieldDiscovery"`
}

type GranularitySpec struct {
	Type      string    `json:"type,omitempty"`
	Segment   string    `json:"segmentGranularity,omitempty"`
	Query     string    `json:"queryGranularity,omitempty"`
	Rollup    bool      `json:"rollup"`
	Intervals []*string `json:"intervals"`
}

func registerType(pkgName *string, msg *descriptor.DescriptorProto, comments Comments, path string) {
	pkg := globalPkg
	if pkgName != nil {
		for _, node := range strings.Split(*pkgName, ".") {
			if pkg == globalPkg && node == "" {
				// Skips leading "."
				continue
			}

			child, ok := pkg.children[node]
			if !ok {
				child = &ProtoPackage{
					name:     pkg.name + "." + node,
					parent:   pkg,
					children: make(map[string]*ProtoPackage),
					types:    make(map[string]*descriptor.DescriptorProto),
					comments: make(map[string]Comments),
					path:     make(map[string]string),
				}
				pkg.children[node] = child
			}
			pkg = child
		}
	}

	pkg.types[msg.GetName()] = msg
	pkg.comments[msg.GetName()] = comments
	pkg.path[msg.GetName()] = path
}

func emptyResultWithError(err error) ([]*FlattenField, *DimensionSpec, []*MetricField, *TimestampField, error) {
	return nil, nil, nil, nil, err
}

func getRealType(desc *descriptor.FieldDescriptorProto) (string, error) {
	fieldType, ok := typeFromFieldType[desc.GetType()]
	if !ok {
		return "", fmt.Errorf("unrecognized field type: %s", desc.GetType().String())
	}

	wkt, ok := typeFromWKT[desc.GetTypeName()]
	if ok && fieldType == "record" {
		return wkt, nil
	}
	return fieldType, nil
}

func isTruthlyDruidStringType(desc *descriptor.FieldDescriptorProto) bool {
	if descriptor.FieldDescriptorProto_TYPE_STRING == desc.GetType() ||
		desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED ||
		desc.GetTypeName() == ".google.protobuf.StringValue" {
		return true
	}
	return false
}

func convertField(
	curPkg *ProtoPackage,
	desc *descriptor.FieldDescriptorProto,
	msgOpts *protos.DruidIngestionOptions,
	parentMessages map[*descriptor.DescriptorProto]bool,
	prefixName string,
	comments Comments,
	path string) (flattenFields []*FlattenField, dimensionSpec *DimensionSpec, metricFields []*MetricField, timestampField *TimestampField, err error) {

	dimensionSpec = &DimensionSpec{Dimensions: []*DimensionField{}}
	fieldName := desc.GetName()
	fieldType, err := getRealType(desc)
	if err != nil {
		return emptyResultWithError(err)
	}
	flattenedFieldName := fieldName
	// in case set prefix = ""
	if len(prefixName) > 0 {
		flattenedFieldName = fmt.Sprintf("%s_%s", prefixName, fieldName)
	}

	fieldMode, ok := modeFromFieldLabel[desc.GetLabel()]
	if !ok {
		return emptyResultWithError(fmt.Errorf("unrecognized field label: %s", desc.GetLabel().String()))
	}

	isFlattened := false
	if len(path) > 0 {
		isFlattened = true
	}

	dimensionField := &DimensionField{
		Type: fieldType,
		Name: fieldName,
	}
	flattenField := &FlattenField{
		Expression: fmt.Sprintf("%s.%s", path, fieldName),
		Name:       fieldName,
	}
	if isFlattened {
		dimensionField.Name = flattenedFieldName
		flattenField.Name = flattenedFieldName
	}

	if fieldType == "record" || fieldMode == "REPEATED" {
		dimensionField.Type = "string"
	}

	// truthly string or array or well know type are string
	if isTruthlyDruidStringType(desc) {
		dimensionField.CreateBitmapIndex = true
		dimensionField.MultiValueHandling = "SORTED_ARRAY"
	}

	opts := desc.GetOptions()
	if opts == nil || !proto.HasExtension(opts, protos.E_Spec) {
		if isFlattened {
			flattenFields = append(flattenFields, flattenField)
		}
		dimensionSpec.Dimensions = append(dimensionSpec.Dimensions, dimensionField)
		return
	}

	opt := proto.GetExtension(opts, protos.E_Spec).(*protos.DruidSpecFieldOptions)
	if opt.Ignore {
		return
	}

	if opt.Timestamp != nil && fieldType == "record" {
		return emptyResultWithError(fmt.Errorf("can not apply timestamp opts for record field '%s'", fieldName))
	}

	if opt.Timestamp != nil && opt.Dimension != nil {
		return emptyResultWithError(fmt.Errorf("can not apply timestamp opts with dimension opts for one field, %s", fieldName))
	}

	if opt.Timestamp != nil {
		timestampField = &TimestampField{
			FieldName: prefixName + fieldName,
		}
		if len(opt.Timestamp.Format) > 0 {
			timestampField.Format = opt.Timestamp.Format
		}
	}

	if isFlattened && timestampField != nil {
		timestampField.FieldName = flattenedFieldName
	}

	if opt.Flatten != nil && (opt.Dimension != nil || opt.Metric != nil || opt.Timestamp != nil) {
		return emptyResultWithError(fmt.Errorf("can not apply flatten opts with dimension/metric/timestamp opts for one field '%s'", desc.GetName()))
	}

	if opt.Dimension != nil {
		if opt.Dimension.CreateBitmapIndex != nil {
			dimensionField.CreateBitmapIndex = opt.Dimension.CreateBitmapIndex.Value
		}
		if len(opt.Dimension.MultiValueHandling) > 0 {
			if _, exists := supportedMultiValueHanldingOpts[opt.Dimension.MultiValueHandling]; exists {
				dimensionField.MultiValueHandling = opt.Dimension.MultiValueHandling
			} else {
				return emptyResultWithError(fmt.Errorf("unsupported multi_value_handling option, get '%s'", opt.Dimension.MultiValueHandling))
			}
		}
		dimensionSpec.Dimensions = append(dimensionSpec.Dimensions, dimensionField)
	} else if opt.Flatten == nil && opt.Timestamp == nil {
		dimensionSpec.Dimensions = append(dimensionSpec.Dimensions, dimensionField)
	}

	if opt.Timestamp != nil {
		dimensionSpec.DimensionExclusions = []string{dimensionField.Name}
	}

	if opt.Metric != nil {
		if len(opt.Metric.MetricName) <= 0 {
			return emptyResultWithError(fmt.Errorf("metric field name must be set"))
		}
		metricField := &MetricField{
			MetricFieldName:    opt.Metric.MetricName,
			ApproxType:         "thetaSketch",
			IsInputThetaSketch: opt.Metric.IsInputThetaSketch,
			Size:               16384,
			FieldName:          prefixName + fieldName,
		}
		if isFlattened {
			metricField.FieldName = flattenedFieldName
		}
		if len(opt.Metric.Type) > 0 {
			if _, ok = supportedMetricAggregators[opt.Metric.Type]; ok {
				metricField.ApproxType = opt.Metric.Type
			} else {
				return emptyResultWithError(fmt.Errorf("not supported metric type %s", opt.Metric.Type))
			}
		}
		if opt.Metric.Size > 0 {
			metricField.Size = opt.Metric.Size
		}
		if len(opt.Metric.Type) > 0 {
			metricField.ApproxType = opt.Metric.Type
		}
		metricFields = append(metricFields, metricField)
	}

	if opt.Flatten != nil {
		nestedPrefixName := fieldName
		if len(opt.Flatten.Prefix) > 0 {
			nestedPrefixName = opt.Flatten.Prefix
		}

		if opt.Flatten.IgnoreName {
			nestedPrefixName = prefixName
		} else if len(opt.Flatten.OutputName) > 0 {
			nestedPrefixName = fmt.Sprintf("%s%s", prefixName, opt.Flatten.OutputName)
		} else {
			nestedPrefixName = fmt.Sprintf("%s_", nestedPrefixName)
			if isFlattened && len(prefixName) > 0 {
				nestedPrefixName = fmt.Sprintf("%s_%s", prefixName, nestedPrefixName)
			}
		}

		newPath := fmt.Sprintf("%s.%s", path, fieldName)

		// flatten primitive field
		if fieldType != "record" {
			newDimensionField := dimensionField
			newDimensionField.Name = nestedPrefixName
			dimensionSpec.Dimensions = append(dimensionSpec.Dimensions, newDimensionField)

			flattenField = &FlattenField{Type: "jq", Name: nestedPrefixName, Expression: newPath}
			if len(opt.Flatten.Type) > 0 && opt.Flatten.Type != "jq" {
				return emptyResultWithError(fmt.Errorf("unsupproted flatten type, got %s", opt.Flatten.Type))
			}

			if len(opt.Flatten.Type) > 0 {
				flattenField.Type = opt.Flatten.Type
			}

			flattenFields = append(flattenFields, flattenField)

			return
		}

		// flatten message field
		nestedFlattendFields, nestedDimensionSpec, nestedMetricFields, nestedTimestampField, err := flattenFieldsForType(curPkg,
			desc.GetTypeName(),
			nestedPrefixName,
			parentMessages,
			opt.Flatten, newPath)
		if err != nil {
			return emptyResultWithError(err)
		}

		if timestampField != nil && nestedTimestampField != nil {
			return emptyResultWithError(fmt.Errorf("detected multiple timestamp spec in %s", desc.GetTypeName()))
		}
		timestampField = nestedTimestampField

		dimensionSpec.Dimensions = append(dimensionSpec.Dimensions, nestedDimensionSpec.Dimensions...)
		dimensionSpec.DimensionExclusions = append(dimensionSpec.DimensionExclusions, nestedDimensionSpec.DimensionExclusions...)
		flattenFields = append(flattenFields, nestedFlattendFields...)
		metricFields = append(metricFields, nestedMetricFields...)
	} else if isFlattened {
		flattenFields = append(flattenFields, flattenField)
	}

	return
}

func flattenFieldsForType(curPkg *ProtoPackage,
	typeName string,
	prefixName string,
	parentMessages map[*descriptor.DescriptorProto]bool,
	flattenOpts *protos.DruidFlattenFieldOptions,
	path string,
) ([]*FlattenField, *DimensionSpec, []*MetricField, *TimestampField, error) {
	recordType, ok, comments, _ := curPkg.lookupType(typeName)
	if !ok {
		return emptyResultWithError(fmt.Errorf("no such type named %s", typeName))
	}

	fieldMsgOpts, err := getDruidOpts(recordType)
	if err != nil {
		return emptyResultWithError(err)
	}

	flattenFields, dimensionSpec, metricFields, timestampField, err := convertMessageType(curPkg, recordType, fieldMsgOpts, parentMessages, prefixName, comments, path)
	if err != nil {
		return emptyResultWithError(err)
	}
	flattenType := "jq"
	if len(flattenOpts.Type) > 0 {
		if flattenOpts.Type != "jq" {
			return emptyResultWithError(fmt.Errorf("unsupproted flatten type, got %s", flattenOpts.Type))
		}
		flattenType = flattenOpts.Type
	}

	for _, fopts := range flattenFields {
		fopts.Type = flattenType
	}
	return flattenFields, dimensionSpec, metricFields, timestampField, err
}

func convertMessageType(
	curPkg *ProtoPackage,
	msg *descriptor.DescriptorProto,
	opts *protos.DruidIngestionOptions,
	parentMessages map[*descriptor.DescriptorProto]bool,
	prefixName string,
	comments Comments,
	path string) (flattenFields []*FlattenField, dimensionSpec *DimensionSpec, metricFields []*MetricField, timestampField *TimestampField, err error) {

	if parentMessages[msg] {
		glog.Infof("Detected recursion for message %s, ignoring subfields", *msg.Name)
		return
	}

	if glog.V(4) {
		glog.Info("Converting message: ", prototext.Format(msg))
	}

	parentMessages[msg] = true
	timestampField = nil
	dimensionSpec = &DimensionSpec{}
	for _, fieldDesc := range msg.GetField() {
		extractedFalttenFields, extractedDimensionSpec, extractedMetricFields, extractedTimestampField, err := convertField(curPkg,
			fieldDesc,
			opts,
			parentMessages,
			prefixName,
			comments,
			path)

		if err != nil {
			glog.Errorf("Failed to convert field %s in %s: %v", fieldDesc.GetName(), msg.GetName(), err)
			return emptyResultWithError(err)
		}

		if extractedDimensionSpec != nil {
			dimensionSpec.merge(extractedDimensionSpec)
		}

		if extractedFalttenFields != nil {
			flattenFields = append(flattenFields, extractedFalttenFields...)
		}

		if extractedMetricFields != nil {
			metricFields = append(metricFields, extractedMetricFields...)
		}
		if extractedTimestampField != nil && timestampField != nil {
			return emptyResultWithError(fmt.Errorf("mulitple timestamp options found in one message %s at field %s", msg.GetName(), fieldDesc.GetName()))
		}
		if timestampField == nil {
			timestampField = extractedTimestampField
		}
	}

	parentMessages[msg] = false

	return
}

func convertFile(file *descriptor.FileDescriptorProto) ([]*plugin.CodeGeneratorResponse_File, error) {
	name := path.Base(file.GetName())
	pkg, ok := globalPkg.relativelyLookupPackage(file.GetPackage())
	if !ok {
		return nil, fmt.Errorf("no such package found: %s", file.GetPackage())
	}

	response := []*plugin.CodeGeneratorResponse_File{}
	comments := ParseComments(file)
	for _, msg := range file.GetMessageType() {
		opts, err := getDruidOpts(msg)
		if err != nil {
			return nil, err
		}
		if opts == nil {
			continue
		}

		dataSourceName := opts.GetDataSourceName()
		if len(dataSourceName) == 0 {
			continue
		}

		glog.V(2).Info("Generating ingestion for a message type ", msg.GetName())
		prefix := ""
		path := ""
		flattenFields, dimensionSpec, metricFields, timestampField, err := convertMessageType(pkg, msg, opts, make(map[*descriptor.DescriptorProto]bool), prefix, comments, path)
		if err != nil {
			glog.Errorf("Failed to convert %s: %v", name, err)
			return nil, err
		}

		var flattenSpec *FlattenSpec
		if len(flattenFields) > 0 {
			flattenSpec = &FlattenSpec{FlattenFields: flattenFields}
		}
		if opts.IoConfig != nil {
			flattenSpec.UseFieldDiscovery = opts.IoConfig.UseFieldDiscovery
		}

		granularitySpec := &GranularitySpec{Type: "uniform", Rollup: true, Query: "none", Segment: "day", Intervals: make([]*string, 0)}
		if granularityOpts := opts.Granularity; granularityOpts != nil {
			if len(granularityOpts.QueryGranularity) > 0 {
				if _, exist := supportedGranularities[strings.ToLower(granularityOpts.QueryGranularity)]; !exist {
					return nil, fmt.Errorf("unsupported queryGranularity, got %s", granularityOpts.QueryGranularity)
				}
				granularitySpec.Query = strings.ToLower(granularityOpts.QueryGranularity)
			}

			if len(granularityOpts.SegmentGranularity) > 0 {
				if _, exist := supportedGranularities[strings.ToLower(granularityOpts.SegmentGranularity)]; !exist {
					return nil, fmt.Errorf("unsupported segmentGranularity, got %s", granularityOpts.SegmentGranularity)
				}
				granularitySpec.Segment = strings.ToLower(granularityOpts.SegmentGranularity)
			}
			granularitySpec.Rollup = granularityOpts.Rollup
		}

		ioConfig := map[string]interface{}{
			"inputFormat": struct {
				Type        string       `json:"type,omitempty"`
				FlattenSpec *FlattenSpec `json:"flattenSpec,omitempty"`
			}{"json", flattenSpec},
		}

		if configuredIOConfig := opts.IoConfig; configuredIOConfig != nil {
			if len(configuredIOConfig.Topic) > 0 {
				ioConfig["topic"] = configuredIOConfig.Topic
			}
			ioConfig["useEarliestOffset"] = configuredIOConfig.UseEarliestOffset

			if len(configuredIOConfig.BootstrapServers) > 0 {
				ioConfig["consumerProperties"] = map[string]interface{}{
					"bootstrap.servers": configuredIOConfig.BootstrapServers,
				}
			}

			if len(configuredIOConfig.Type) > 0 {
				ioConfig["Type"] = configuredIOConfig.Type
			}
		}

		ingestion := map[string]interface{}{
			"spec": map[string]interface{}{
				"dataSchema": struct {
					DataSource      string           `json:"dataSource,omitempty"`
					TimestampSpec   *TimestampField  `json:"timestampSpec,omitempty"`
					DimensionsSpec  *DimensionSpec   `json:"dimensionsSpec,omitempty"`
					MetricsSpec     []*MetricField   `json:"metricsSpec,omitempty"`
					GranularitySpec *GranularitySpec `json:"granularitySpec,omitempty"`
				}{opts.GetDataSourceName(), timestampField, dimensionSpec, metricFields, granularitySpec},
				"ioConfig": ioConfig,
			},
		}

		if len(opts.IngestionType) > 0 {
			if _, exists := supportedIngestionType[opts.IngestionType]; !exists {
				return nil, fmt.Errorf("unsupported ingestion type, got %s", opts.IngestionType)
			}
			ingestion["type"] = opts.IngestionType
		}

		ingestionJson, err := json.MarshalIndent(ingestion, "", " ")
		if err != nil {
			glog.Error("Failed to encode schema", err)
			return nil, err
		}

		resFile := &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(fmt.Sprintf("%s/%s.ingestion", strings.Replace(file.GetPackage(), ".", "/", -1), dataSourceName)),
			Content: proto.String(string(ingestionJson)),
		}
		response = append(response, resFile)
	}

	return response, nil
}

func getDruidOpts(msg *descriptor.DescriptorProto) (*protos.DruidIngestionOptions, error) {
	options := msg.GetOptions()
	if options == nil {
		return nil, nil
	}

	if !proto.HasExtension(options, protos.E_DruidOpts) {
		return nil, nil
	}

	return proto.GetExtension(options, protos.E_DruidOpts).(*protos.DruidIngestionOptions), nil
}

func handleSingleMessageOpt(file *descriptor.FileDescriptorProto, requestParam string) {
	if !strings.Contains(requestParam, "single-message") || len(file.GetMessageType()) == 0 {
		return
	}
	file.MessageType = file.GetMessageType()[:1]
	message := file.GetMessageType()[0]
	message.Options = &descriptor.MessageOptions{}
	fileName := file.GetName()
	proto.SetExtension(message.GetOptions(), protos.E_DruidOpts, &protos.DruidIngestionOptions{
		DataSourceName: fileName[strings.LastIndexByte(fileName, '/')+1 : strings.LastIndexByte(fileName, '.')],
	})
}

func Convert(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	generateTargets := make(map[string]bool)
	for _, file := range req.GetFileToGenerate() {
		generateTargets[file] = true
	}

	res := &plugin.CodeGeneratorResponse{}
	for _, file := range req.GetProtoFile() {
		for _, msg := range file.GetMessageType() {
			glog.V(1).Infof("Loading a message type %s from package %s", msg.GetName(), file.GetPackage())
			registerType(file.Package, msg, ParseComments(file), "")
		}
	}
	for _, file := range req.GetProtoFile() {
		if _, ok := generateTargets[file.GetName()]; ok {
			glog.V(1).Info("Converting ", file.GetName())
			handleSingleMessageOpt(file, req.GetParameter())
			converted, err := convertFile(file)
			if err != nil {
				res.Error = proto.String(fmt.Sprintf("Failed to convert %s: %v", file.GetName(), err))
				return res, err
			}
			res.File = append(res.File, converted...)
		}
	}
	return res, nil
}

// ConvertFrom converts input from protoc to a CodeGeneratorRequest and starts conversion
// Returning a CodeGeneratorResponse containing either an error or the results of converting the given proto
func ConvertFrom(rd io.Reader) (*plugin.CodeGeneratorResponse, error) {
	glog.V(1).Info("Reading code generation request")
	input, err := ioutil.ReadAll(rd)
	if err != nil {
		glog.Error("Failed to read request:", err)
		return nil, err
	}
	req := &plugin.CodeGeneratorRequest{}
	err = proto.Unmarshal(input, req)
	if err != nil {
		glog.Error("Can't unmarshal input:", err)
		return nil, err
	}

	glog.V(1).Info("Converting input")
	return Convert(req)
}
