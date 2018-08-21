// Code generated by zanzibar
// @generated
// Checksum : yD4J+1Fe1FCn/GusqpLr4A==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(in *jlexer.Lexer, out *Bar_ArgWithManyQueryParams_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(BarResponse)
				}
				easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(in, &*out.Success)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(out *jwriter.Writer, in Bar_ArgWithManyQueryParams_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(out, *in.Success)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithManyQueryParams_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithManyQueryParams_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithManyQueryParams_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithManyQueryParams_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(l, v)
}
func easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(in *jlexer.Lexer, out *BarResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var IntWithRangeSet bool
	var IntWithoutRangeSet bool
	var MapIntWithRangeSet bool
	var MapIntWithoutRangeSet bool
	var BinaryFieldSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "intWithRange":
			out.IntWithRange = int32(in.Int32())
			IntWithRangeSet = true
		case "intWithoutRange":
			out.IntWithoutRange = int32(in.Int32())
			IntWithoutRangeSet = true
		case "mapIntWithRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithRange = make(map[UUID]int32)
				} else {
					out.MapIntWithRange = nil
				}
				for !in.IsDelim('}') {
					key := UUID(in.String())
					in.WantColon()
					var v1 int32
					v1 = int32(in.Int32())
					(out.MapIntWithRange)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithRangeSet = true
		case "mapIntWithoutRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithoutRange = make(map[string]int32)
				} else {
					out.MapIntWithoutRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 int32
					v2 = int32(in.Int32())
					(out.MapIntWithoutRange)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithoutRangeSet = true
		case "binaryField":
			if in.IsNull() {
				in.Skip()
				out.BinaryField = nil
			} else {
				out.BinaryField = in.Bytes()
			}
			BinaryFieldSet = true
		case "nextResponse":
			if in.IsNull() {
				in.Skip()
				out.NextResponse = nil
			} else {
				if out.NextResponse == nil {
					out.NextResponse = new(BarResponse)
				}
				easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(in, &*out.NextResponse)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !IntWithRangeSet {
		in.AddError(fmt.Errorf("key 'intWithRange' is required"))
	}
	if !IntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'intWithoutRange' is required"))
	}
	if !MapIntWithRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithRange' is required"))
	}
	if !MapIntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithoutRange' is required"))
	}
	if !BinaryFieldSet {
		in.AddError(fmt.Errorf("key 'binaryField' is required"))
	}
}
func easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(out *jwriter.Writer, in BarResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"stringField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.StringField))
	}
	{
		const prefix string = ",\"intWithRange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.IntWithRange))
	}
	{
		const prefix string = ",\"intWithoutRange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.IntWithoutRange))
	}
	{
		const prefix string = ",\"mapIntWithRange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MapIntWithRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.MapIntWithRange {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				out.Int32(int32(v4Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"mapIntWithoutRange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MapIntWithoutRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.MapIntWithoutRange {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.Int32(int32(v5Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"binaryField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.BinaryField)
	}
	if in.NextResponse != nil {
		const prefix string = ",\"nextResponse\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(out, *in.NextResponse)
	}
	out.RawByte('}')
}
func easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(in *jlexer.Lexer, out *Bar_ArgWithManyQueryParams_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var AStrSet bool
	var ABoolSet bool
	var AInt8Set bool
	var AInt16Set bool
	var AInt32Set bool
	var AInt64Set bool
	var AFloat64Set bool
	var AUUIDSet bool
	var AListUUIDSet bool
	var AStringListSet bool
	var AUUIDListSet bool
	var ATsSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "aStr":
			out.AStr = string(in.String())
			AStrSet = true
		case "anOptStr":
			if in.IsNull() {
				in.Skip()
				out.AnOptStr = nil
			} else {
				if out.AnOptStr == nil {
					out.AnOptStr = new(string)
				}
				*out.AnOptStr = string(in.String())
			}
		case "aBool":
			out.ABool = bool(in.Bool())
			ABoolSet = true
		case "anOptBool":
			if in.IsNull() {
				in.Skip()
				out.AnOptBool = nil
			} else {
				if out.AnOptBool == nil {
					out.AnOptBool = new(bool)
				}
				*out.AnOptBool = bool(in.Bool())
			}
		case "aInt8":
			out.AInt8 = int8(in.Int8())
			AInt8Set = true
		case "anOptInt8":
			if in.IsNull() {
				in.Skip()
				out.AnOptInt8 = nil
			} else {
				if out.AnOptInt8 == nil {
					out.AnOptInt8 = new(int8)
				}
				*out.AnOptInt8 = int8(in.Int8())
			}
		case "aInt16":
			out.AInt16 = int16(in.Int16())
			AInt16Set = true
		case "anOptInt16":
			if in.IsNull() {
				in.Skip()
				out.AnOptInt16 = nil
			} else {
				if out.AnOptInt16 == nil {
					out.AnOptInt16 = new(int16)
				}
				*out.AnOptInt16 = int16(in.Int16())
			}
		case "aInt32":
			out.AInt32 = int32(in.Int32())
			AInt32Set = true
		case "anOptInt32":
			if in.IsNull() {
				in.Skip()
				out.AnOptInt32 = nil
			} else {
				if out.AnOptInt32 == nil {
					out.AnOptInt32 = new(int32)
				}
				*out.AnOptInt32 = int32(in.Int32())
			}
		case "aInt64":
			out.AInt64 = int64(in.Int64())
			AInt64Set = true
		case "anOptInt64":
			if in.IsNull() {
				in.Skip()
				out.AnOptInt64 = nil
			} else {
				if out.AnOptInt64 == nil {
					out.AnOptInt64 = new(int64)
				}
				*out.AnOptInt64 = int64(in.Int64())
			}
		case "aFloat64":
			out.AFloat64 = float64(in.Float64())
			AFloat64Set = true
		case "anOptFloat64":
			if in.IsNull() {
				in.Skip()
				out.AnOptFloat64 = nil
			} else {
				if out.AnOptFloat64 == nil {
					out.AnOptFloat64 = new(float64)
				}
				*out.AnOptFloat64 = float64(in.Float64())
			}
		case "aUUID":
			out.AUUID = UUID(in.String())
			AUUIDSet = true
		case "anOptUUID":
			if in.IsNull() {
				in.Skip()
				out.AnOptUUID = nil
			} else {
				if out.AnOptUUID == nil {
					out.AnOptUUID = new(UUID)
				}
				*out.AnOptUUID = UUID(in.String())
			}
		case "aListUUID":
			if in.IsNull() {
				in.Skip()
				out.AListUUID = nil
			} else {
				in.Delim('[')
				if out.AListUUID == nil {
					if !in.IsDelim(']') {
						out.AListUUID = make([]UUID, 0, 4)
					} else {
						out.AListUUID = []UUID{}
					}
				} else {
					out.AListUUID = (out.AListUUID)[:0]
				}
				for !in.IsDelim(']') {
					var v8 UUID
					v8 = UUID(in.String())
					out.AListUUID = append(out.AListUUID, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
			AListUUIDSet = true
		case "anOptListUUID":
			if in.IsNull() {
				in.Skip()
				out.AnOptListUUID = nil
			} else {
				in.Delim('[')
				if out.AnOptListUUID == nil {
					if !in.IsDelim(']') {
						out.AnOptListUUID = make([]UUID, 0, 4)
					} else {
						out.AnOptListUUID = []UUID{}
					}
				} else {
					out.AnOptListUUID = (out.AnOptListUUID)[:0]
				}
				for !in.IsDelim(']') {
					var v9 UUID
					v9 = UUID(in.String())
					out.AnOptListUUID = append(out.AnOptListUUID, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "aStringList":
			if in.IsNull() {
				in.Skip()
				out.AStringList = nil
			} else {
				in.Delim('[')
				if out.AStringList == nil {
					if !in.IsDelim(']') {
						out.AStringList = make(StringList, 0, 4)
					} else {
						out.AStringList = StringList{}
					}
				} else {
					out.AStringList = (out.AStringList)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.AStringList = append(out.AStringList, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
			AStringListSet = true
		case "anOptStringList":
			if in.IsNull() {
				in.Skip()
				out.AnOptStringList = nil
			} else {
				in.Delim('[')
				if out.AnOptStringList == nil {
					if !in.IsDelim(']') {
						out.AnOptStringList = make(StringList, 0, 4)
					} else {
						out.AnOptStringList = StringList{}
					}
				} else {
					out.AnOptStringList = (out.AnOptStringList)[:0]
				}
				for !in.IsDelim(']') {
					var v11 string
					v11 = string(in.String())
					out.AnOptStringList = append(out.AnOptStringList, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "aUUIDList":
			if in.IsNull() {
				in.Skip()
				out.AUUIDList = nil
			} else {
				in.Delim('[')
				if out.AUUIDList == nil {
					if !in.IsDelim(']') {
						out.AUUIDList = make(UUIDList, 0, 4)
					} else {
						out.AUUIDList = UUIDList{}
					}
				} else {
					out.AUUIDList = (out.AUUIDList)[:0]
				}
				for !in.IsDelim(']') {
					var v12 UUID
					v12 = UUID(in.String())
					out.AUUIDList = append(out.AUUIDList, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
			AUUIDListSet = true
		case "anOptUUIDList":
			if in.IsNull() {
				in.Skip()
				out.AnOptUUIDList = nil
			} else {
				in.Delim('[')
				if out.AnOptUUIDList == nil {
					if !in.IsDelim(']') {
						out.AnOptUUIDList = make(UUIDList, 0, 4)
					} else {
						out.AnOptUUIDList = UUIDList{}
					}
				} else {
					out.AnOptUUIDList = (out.AnOptUUIDList)[:0]
				}
				for !in.IsDelim(']') {
					var v13 UUID
					v13 = UUID(in.String())
					out.AnOptUUIDList = append(out.AnOptUUIDList, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "aTs":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ATs).UnmarshalJSON(data))
			}
			ATsSet = true
		case "anOptTs":
			if in.IsNull() {
				in.Skip()
				out.AnOptTs = nil
			} else {
				if out.AnOptTs == nil {
					out.AnOptTs = new(Timestamp)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.AnOptTs).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !AStrSet {
		in.AddError(fmt.Errorf("key 'aStr' is required"))
	}
	if !ABoolSet {
		in.AddError(fmt.Errorf("key 'aBool' is required"))
	}
	if !AInt8Set {
		in.AddError(fmt.Errorf("key 'aInt8' is required"))
	}
	if !AInt16Set {
		in.AddError(fmt.Errorf("key 'aInt16' is required"))
	}
	if !AInt32Set {
		in.AddError(fmt.Errorf("key 'aInt32' is required"))
	}
	if !AInt64Set {
		in.AddError(fmt.Errorf("key 'aInt64' is required"))
	}
	if !AFloat64Set {
		in.AddError(fmt.Errorf("key 'aFloat64' is required"))
	}
	if !AUUIDSet {
		in.AddError(fmt.Errorf("key 'aUUID' is required"))
	}
	if !AListUUIDSet {
		in.AddError(fmt.Errorf("key 'aListUUID' is required"))
	}
	if !AStringListSet {
		in.AddError(fmt.Errorf("key 'aStringList' is required"))
	}
	if !AUUIDListSet {
		in.AddError(fmt.Errorf("key 'aUUIDList' is required"))
	}
	if !ATsSet {
		in.AddError(fmt.Errorf("key 'aTs' is required"))
	}
}
func easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(out *jwriter.Writer, in Bar_ArgWithManyQueryParams_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"aStr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AStr))
	}
	if in.AnOptStr != nil {
		const prefix string = ",\"anOptStr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AnOptStr))
	}
	{
		const prefix string = ",\"aBool\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ABool))
	}
	if in.AnOptBool != nil {
		const prefix string = ",\"anOptBool\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.AnOptBool))
	}
	{
		const prefix string = ",\"aInt8\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(in.AInt8))
	}
	if in.AnOptInt8 != nil {
		const prefix string = ",\"anOptInt8\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(*in.AnOptInt8))
	}
	{
		const prefix string = ",\"aInt16\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int16(int16(in.AInt16))
	}
	if in.AnOptInt16 != nil {
		const prefix string = ",\"anOptInt16\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int16(int16(*in.AnOptInt16))
	}
	{
		const prefix string = ",\"aInt32\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.AInt32))
	}
	if in.AnOptInt32 != nil {
		const prefix string = ",\"anOptInt32\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.AnOptInt32))
	}
	{
		const prefix string = ",\"aInt64\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AInt64))
	}
	if in.AnOptInt64 != nil {
		const prefix string = ",\"anOptInt64\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.AnOptInt64))
	}
	{
		const prefix string = ",\"aFloat64\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.AFloat64))
	}
	if in.AnOptFloat64 != nil {
		const prefix string = ",\"anOptFloat64\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(*in.AnOptFloat64))
	}
	{
		const prefix string = ",\"aUUID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AUUID))
	}
	if in.AnOptUUID != nil {
		const prefix string = ",\"anOptUUID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AnOptUUID))
	}
	{
		const prefix string = ",\"aListUUID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.AListUUID == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.AListUUID {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	if len(in.AnOptListUUID) != 0 {
		const prefix string = ",\"anOptListUUID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.AnOptListUUID {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"aStringList\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.AStringList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.AStringList {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.String(string(v19))
			}
			out.RawByte(']')
		}
	}
	if len(in.AnOptStringList) != 0 {
		const prefix string = ",\"anOptStringList\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.AnOptStringList {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"aUUIDList\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.AUUIDList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v22, v23 := range in.AUUIDList {
				if v22 > 0 {
					out.RawByte(',')
				}
				out.String(string(v23))
			}
			out.RawByte(']')
		}
	}
	if len(in.AnOptUUIDList) != 0 {
		const prefix string = ",\"anOptUUIDList\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v24, v25 := range in.AnOptUUIDList {
				if v24 > 0 {
					out.RawByte(',')
				}
				out.String(string(v25))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"aTs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ATs).MarshalJSON())
	}
	if in.AnOptTs != nil {
		const prefix string = ",\"anOptTs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.AnOptTs).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithManyQueryParams_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithManyQueryParams_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithManyQueryParams_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithManyQueryParams_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(l, v)
}
