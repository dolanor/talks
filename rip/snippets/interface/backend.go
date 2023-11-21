package iface

import (
	"context"
	"errors"
	"fmt"
)

type (
	Input struct {
		A int
		B string
	}

	Output struct {
		C int
		D string
	}
)

var ErrBadArgument = errors.New("bad format for the request")
var ErrInternalServer = errors.New("internal server error")

// start BackendFuncSignature OMIT
type BackendFunc func(ctx context.Context, anyIn interface{}) (anyOut interface{}, err error)

// end BackendFuncSignature OMIT

// start BackendWrapper OMIT
func BackendWrapper(ctx context.Context, anyIn interface{}) (anyOut interface{}, err error) { // HLSignature
	// end BackendWrapperSignature OMIT
	mapIn, ok := anyIn.(map[string]interface{}) // HLBackend
	if !ok {                                    // HLBackend
		return Output{}, ErrBadArgument // HLBackend
		// And we need to catch this error in our handler to send a http.StatusBadRequest // HLBackend
	} // HLBackend

	in, err := inputFromMap(mapIn) // HLBackend
	if err != nil {                // HLBackend
		return Output{}, ErrBadArgument // HLBackend
	} // HLBackend

	// Do the backend-y stuff
	out, err := realBackendCall(ctx, in)
	if err != nil {
		return Output{}, ErrInternalServer
	}

	return out, nil
}

// end BackendWrapper OMIT

// start realBackendCall OMIT
func realBackendCall(ctx context.Context, in Input) (Output, error) {
	// end realBackendCallSignature OMIT
	return Output{
		C: in.A + 2,
		D: in.B + "2",
	}, nil
}

// end realBackendCall OMIT

// start inputFromMap OMIT
func inputFromMap(mapIn map[string]interface{}) (Input, error) {
	// end inputFromMapSignature OMIT
	ma, ok := mapIn["A"]
	if !ok {
		return Input{}, errors.New("inputFromMap: no A")
	}
	mb, ok := mapIn["B"]
	if !ok {
		return Input{}, errors.New("inputFromMap: no B")
	}
	fa, ok := ma.(float64)
	if !ok {
		return Input{}, fmt.Errorf("inputFromMap: A is not an int: %T", ma)
	}
	// check the number range
	a := int(fa)
	b, ok := mb.(string)
	if !ok {
		return Input{}, fmt.Errorf("inputFromMap: B is not an int: %T", mb)
	}
	return Input{
		A: a,
		B: b,
	}, nil
}

// end inputFromMap OMIT
