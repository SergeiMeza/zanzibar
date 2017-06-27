// Code generated by zanzibar
// @generated
// Checksum : OBYKw8jf6+/KhAG4BjY1Xw==
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

func easyjsonC4391290DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders(in *jlexer.Lexer, out *Bar_ArgWithHeaders_Result) {
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
				(*out.Success).UnmarshalEasyJSON(in)
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
func easyjsonC4391290EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders(out *jwriter.Writer, in Bar_ArgWithHeaders_Result) {
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
			(*in.Success).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithHeaders_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC4391290EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithHeaders_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC4391290EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithHeaders_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC4391290DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithHeaders_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC4391290DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders(l, v)
}
func easyjsonC4391290DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders1(in *jlexer.Lexer, out *Bar_ArgWithHeaders_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var NameSet bool
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
		case "name":
			out.Name = string(in.String())
			NameSet = true
		case "userUUID":
			if in.IsNull() {
				in.Skip()
				out.UserUUID = nil
			} else {
				if out.UserUUID == nil {
					out.UserUUID = new(string)
				}
				*out.UserUUID = string(in.String())
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
	if !NameSet {
		in.AddError(fmt.Errorf("key 'name' is required"))
	}
}
func easyjsonC4391290EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders1(out *jwriter.Writer, in Bar_ArgWithHeaders_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if in.UserUUID != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"userUUID\":")
		if in.UserUUID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.UserUUID))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithHeaders_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC4391290EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithHeaders_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC4391290EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithHeaders_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC4391290DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithHeaders_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC4391290DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithHeaders1(l, v)
}