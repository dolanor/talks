package iface

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestHandleMoreBoilerplate(t *testing.T) {
	is := is.New(t)

	h := Handle(http.MethodGet, BackendWrapper)

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(Input{A: 1, B: "my input"})
	is.NoErr(err) // encoding input

	req, err := http.NewRequest(http.MethodPost, "application/json", &b)
	is.NoErr(err) // new request

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	b.Reset()
	_, err = io.Copy(&b, rr.Result().Body)
	is.NoErr(err) // create output buffer
	t.Log(b.String())
	is.Equal(200, rr.Result().StatusCode) // error in http

	var out Output
	err = json.NewDecoder(&b).Decode(&out)
	is.NoErr(err) // decoding output

	is.Equal(3, out.C)
	is.Equal("my input2", out.D)
}
