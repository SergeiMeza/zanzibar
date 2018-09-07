// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package apprentice

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Apprentice_GetRequest_Args represents the arguments for the Apprentice.getRequest function.
//
// The arguments for getRequest are sent and received over the wire as this struct.
type Apprentice_GetRequest_Args struct {
	TraceID           string  `json:"traceID,required"`
	ClientServiceName *string `json:"clientServiceName,omitempty"`
	ClientMethod      *string `json:"clientMethod,omitempty"`
}

// ToWire translates a Apprentice_GetRequest_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Apprentice_GetRequest_Args) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.TraceID), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.ClientServiceName != nil {
		w, err = wire.NewValueString(*(v.ClientServiceName)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.ClientMethod != nil {
		w, err = wire.NewValueString(*(v.ClientMethod)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Apprentice_GetRequest_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Apprentice_GetRequest_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Apprentice_GetRequest_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Apprentice_GetRequest_Args) FromWire(w wire.Value) error {
	var err error

	traceIDIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.TraceID, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				traceIDIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.ClientServiceName = &x
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.ClientMethod = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !traceIDIsSet {
		return errors.New("field TraceID of Apprentice_GetRequest_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Apprentice_GetRequest_Args
// struct.
func (v *Apprentice_GetRequest_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("TraceID: %v", v.TraceID)
	i++
	if v.ClientServiceName != nil {
		fields[i] = fmt.Sprintf("ClientServiceName: %v", *(v.ClientServiceName))
		i++
	}
	if v.ClientMethod != nil {
		fields[i] = fmt.Sprintf("ClientMethod: %v", *(v.ClientMethod))
		i++
	}

	return fmt.Sprintf("Apprentice_GetRequest_Args{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Apprentice_GetRequest_Args match the
// provided Apprentice_GetRequest_Args.
//
// This function performs a deep comparison.
func (v *Apprentice_GetRequest_Args) Equals(rhs *Apprentice_GetRequest_Args) bool {
	if !(v.TraceID == rhs.TraceID) {
		return false
	}
	if !_String_EqualsPtr(v.ClientServiceName, rhs.ClientServiceName) {
		return false
	}
	if !_String_EqualsPtr(v.ClientMethod, rhs.ClientMethod) {
		return false
	}

	return true
}

// GetTraceID returns the value of TraceID if it is set or its
// zero value if it is unset.
func (v *Apprentice_GetRequest_Args) GetTraceID() (o string) { return v.TraceID }

// GetClientServiceName returns the value of ClientServiceName if it is set or its
// zero value if it is unset.
func (v *Apprentice_GetRequest_Args) GetClientServiceName() (o string) {
	if v.ClientServiceName != nil {
		return *v.ClientServiceName
	}

	return
}

// GetClientMethod returns the value of ClientMethod if it is set or its
// zero value if it is unset.
func (v *Apprentice_GetRequest_Args) GetClientMethod() (o string) {
	if v.ClientMethod != nil {
		return *v.ClientMethod
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "getRequest" for this struct.
func (v *Apprentice_GetRequest_Args) MethodName() string {
	return "getRequest"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Apprentice_GetRequest_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Apprentice_GetRequest_Helper provides functions that aid in handling the
// parameters and return values of the Apprentice.getRequest
// function.
var Apprentice_GetRequest_Helper = struct {
	// Args accepts the parameters of getRequest in-order and returns
	// the arguments struct for the function.
	Args func(
		traceID string,
		clientServiceName *string,
		clientMethod *string,
	) *Apprentice_GetRequest_Args

	// IsException returns true if the given error can be thrown
	// by getRequest.
	//
	// An error can be thrown by getRequest only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for getRequest
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// getRequest into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by getRequest
	//
	//   value, err := getRequest(args)
	//   result, err := Apprentice_GetRequest_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from getRequest: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(string, error) (*Apprentice_GetRequest_Result, error)

	// UnwrapResponse takes the result struct for getRequest
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if getRequest threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Apprentice_GetRequest_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Apprentice_GetRequest_Result) (string, error)
}{}

func init() {
	Apprentice_GetRequest_Helper.Args = func(
		traceID string,
		clientServiceName *string,
		clientMethod *string,
	) *Apprentice_GetRequest_Args {
		return &Apprentice_GetRequest_Args{
			TraceID:           traceID,
			ClientServiceName: clientServiceName,
			ClientMethod:      clientMethod,
		}
	}

	Apprentice_GetRequest_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *OperationError:
			return true
		default:
			return false
		}
	}

	Apprentice_GetRequest_Helper.WrapResponse = func(success string, err error) (*Apprentice_GetRequest_Result, error) {
		if err == nil {
			return &Apprentice_GetRequest_Result{Success: &success}, nil
		}

		switch e := err.(type) {
		case *OperationError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Apprentice_GetRequest_Result.Err")
			}
			return &Apprentice_GetRequest_Result{Err: e}, nil
		}

		return nil, err
	}
	Apprentice_GetRequest_Helper.UnwrapResponse = func(result *Apprentice_GetRequest_Result) (success string, err error) {
		if result.Err != nil {
			err = result.Err
			return
		}

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Apprentice_GetRequest_Result represents the result of a Apprentice.getRequest function call.
//
// The result of a getRequest execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Apprentice_GetRequest_Result struct {
	// Value returned by getRequest after a successful execution.
	Success *string         `json:"success,omitempty"`
	Err     *OperationError `json:"err,omitempty"`
}

// ToWire translates a Apprentice_GetRequest_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Apprentice_GetRequest_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueString(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.Err != nil {
		w, err = v.Err.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Apprentice_GetRequest_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _OperationError_Read(w wire.Value) (*OperationError, error) {
	var v OperationError
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Apprentice_GetRequest_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Apprentice_GetRequest_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Apprentice_GetRequest_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Apprentice_GetRequest_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Err, err = _OperationError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.Err != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Apprentice_GetRequest_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Apprentice_GetRequest_Result
// struct.
func (v *Apprentice_GetRequest_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	if v.Err != nil {
		fields[i] = fmt.Sprintf("Err: %v", v.Err)
		i++
	}

	return fmt.Sprintf("Apprentice_GetRequest_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Apprentice_GetRequest_Result match the
// provided Apprentice_GetRequest_Result.
//
// This function performs a deep comparison.
func (v *Apprentice_GetRequest_Result) Equals(rhs *Apprentice_GetRequest_Result) bool {
	if !_String_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	if !((v.Err == nil && rhs.Err == nil) || (v.Err != nil && rhs.Err != nil && v.Err.Equals(rhs.Err))) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Apprentice_GetRequest_Result) GetSuccess() (o string) {
	if v.Success != nil {
		return *v.Success
	}

	return
}

// GetErr returns the value of Err if it is set or its
// zero value if it is unset.
func (v *Apprentice_GetRequest_Result) GetErr() (o *OperationError) {
	if v.Err != nil {
		return v.Err
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "getRequest" for this struct.
func (v *Apprentice_GetRequest_Result) MethodName() string {
	return "getRequest"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Apprentice_GetRequest_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
