package classic

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type T struct {
	Field1 string
	Field2 time.Time
}

func classicHandler(w http.ResponseWriter, r *http.Request) {
	// start classic OMIT
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := backendCall(r.Context(), t) // HLinterestingCall
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end classic OMIT
}

func validate(t T) error {
	// no problem
	return nil
}

type Response struct{}

func backendCall(ctx context.Context, t T) (Response, error) {
	log.Println("doing funky stuff")
	return Response{}, nil
}

type T2 struct {
	Field3 int
	Field4 float64
}

func classicHandler2(w http.ResponseWriter, r *http.Request) {
	// start classic2 OMIT
	var t T2                                  // HLclassic2
	err := json.NewDecoder(r.Body).Decode(&t) // HLgarbage2
	if err != nil {                           // HLgarbage2
		http.Error(w, err.Error(), http.StatusBadRequest) // HLgarbage2
		return                                            // HLgarbage2
	} // HLgarbage2

	err = validate2(t) // HLclassic2
	if err != nil {    // HLgarbage2
		http.Error(w, err.Error(), http.StatusBadRequest) // HLgarbage2
		return                                            // HLgarbage2
	} // HLgarbage2

	resp, err := backendCall2(r.Context(), t) // HLclassic2
	if err != nil {                           // HLgarbage2
		http.Error(w, err.Error(), http.StatusInternalServerError) // HLgarbage2
		return                                                     // HLgarbage2
	} // HLgarbage2

	err = json.NewEncoder(w).Encode(resp) // HLgarbage2
	if err != nil {                       // HLgarbage2
		http.Error(w, err.Error(), http.StatusInternalServerError) // HLgarbage2
		return                                                     // HLgarbage2
	} // HLgarbage2
	// end classic2 OMIT
}

func validate2(t T2) error {
	// no problem
	return nil
}

type Response2 struct{}

func backendCall2(ctx context.Context, t T2) (Response, error) {
	log.Println("doing funky stuff")
	return Response{}, nil
}
