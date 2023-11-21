package iface

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// start CommonBackendCallLessBoilerplate OMIT
func CommonBackendCallLessBoilerplateNOKEEP(ctx context.Context, req Input) (Output, error) { // HLCBC
	// end CommonBackendCallLessSignature OMIT

	// Do the backend-y stuff // HLBackend

	return Output{}, nil
}

// end CommonBackendCallLessBoilerplate OMIT

// start Handle OMIT
func Handle(method string, f BackendFunc) http.HandlerFunc {
	// end HandleSignature OMIT
	return func(w http.ResponseWriter, r *http.Request) {
		var in map[string]interface{} // HL

		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			http.Error(w, fmt.Errorf("json decode: %w", err).Error(), http.StatusBadRequest)
			return
		}

		out, err := f(r.Context(), in) // HL
		if err != nil {                // err handler behaviour? // HL
			http.Error(w, fmt.Errorf("backend call: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			http.Error(w, fmt.Errorf("json encode: %w", err).Error(), http.StatusInternalServerError)
			return
		}
	}
}

// end Handle OMIT
