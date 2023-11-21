package generic

import (
	"context"
	"net/http"
)

// start BackendFuncSignature OMIT
type BackendFunc[Input, Output any] func(ctx context.Context, input Input) (output Output, err error)

// end BackendFuncSignature OMIT

func Handle[Input, Output any](method string, f BackendFunc[Input, Output]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
