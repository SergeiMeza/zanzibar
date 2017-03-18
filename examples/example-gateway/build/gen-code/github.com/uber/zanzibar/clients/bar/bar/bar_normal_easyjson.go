// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package bar

import (
	json "encoding/json"
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

func easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal(in *jlexer.Lexer, out *Bar_Normal_Result) {
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
				easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar(in, &*out.Success)
			}
		case "barException":
			if in.IsNull() {
				in.Skip()
				out.BarException = nil
			} else {
				if out.BarException == nil {
					out.BarException = new(BarException)
				}
				easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar1(in, &*out.BarException)
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
func easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal(out *jwriter.Writer, in Bar_Normal_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar(out, *in.Success)
		}
	}
	if in.BarException != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"barException\":")
		if in.BarException == nil {
			out.RawString("null")
		} else {
			easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar1(out, *in.BarException)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_Normal_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_Normal_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_Normal_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_Normal_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal(l, v)
}
func easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar1(in *jlexer.Lexer, out *BarException) {
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
		case "stringField":
			out.StringField = string(in.String())
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
func easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar1(out *jwriter.Writer, in BarException) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	out.RawByte('}')
}
func easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar(in *jlexer.Lexer, out *BarResponse) {
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
		case "stringField":
			out.StringField = string(in.String())
		case "intWithRange":
			out.IntWithRange = int32(in.Int32())
		case "intWithoutRange":
			out.IntWithoutRange = int32(in.Int32())
		case "mapIntWithRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithRange = make(map[string]int32)
				} else {
					out.MapIntWithRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 int32
					v1 = int32(in.Int32())
					(out.MapIntWithRange)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
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
func easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar(out *jwriter.Writer, in BarResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithRange\":")
	out.Int32(int32(in.IntWithRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithoutRange\":")
	out.Int32(int32(in.IntWithoutRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithRange\":")
	if in.MapIntWithRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v3First := true
		for v3Name, v3Value := range in.MapIntWithRange {
			if !v3First {
				out.RawByte(',')
			}
			v3First = false
			out.String(string(v3Name))
			out.RawByte(':')
			out.Int32(int32(v3Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithoutRange\":")
	if in.MapIntWithoutRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v4First := true
		for v4Name, v4Value := range in.MapIntWithoutRange {
			if !v4First {
				out.RawByte(',')
			}
			v4First = false
			out.String(string(v4Name))
			out.RawByte(':')
			out.Int32(int32(v4Value))
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}
func easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal1(in *jlexer.Lexer, out *Bar_Normal_Args) {
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
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(BarRequest)
				}
				easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar2(in, &*out.Request)
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
func easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal1(out *jwriter.Writer, in Bar_Normal_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	if in.Request == nil {
		out.RawString("null")
	} else {
		easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar2(out, *in.Request)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_Normal_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_Normal_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_Normal_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_Normal_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBarBarNormal1(l, v)
}
func easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar2(in *jlexer.Lexer, out *BarRequest) {
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
		case "stringField":
			out.StringField = string(in.String())
		case "boolField":
			out.BoolField = bool(in.Bool())
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
func easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeGithubComUberZanzibarClientsBarBar2(out *jwriter.Writer, in BarRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"boolField\":")
	out.Bool(bool(in.BoolField))
	out.RawByte('}')
}