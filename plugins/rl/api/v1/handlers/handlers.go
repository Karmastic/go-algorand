package handlers

import (
	"encoding/json"
	"github.com/algorand/go-algorand/plugins/rl/teal"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/algorand/go-algorand/daemon/algod/api/server/lib"
	"github.com/algorand/go-algorand/logging"

	"github.com/algorand/go-algorand/plugins/rl/api/spec/v1"
)

// Response is a generic interface wrapping any data returned by the server.
// We wrap every type in a Response type so that we can swagger annotate them.
//
// Each response must have a Body (a payload). We
// write an interface for this because it better mirrors the
// go-swagger annotation style (which requires swagger colon responses
// to have an embedded body struct of the actual data to be sent. of
// course, they can also have headers and the sort.)
// Anything implementing the Response interface will naturally be
// able to be annotated by swagger:response. This also allows package
// functions to naturally unwrap Response types and send the underlying
// Body through another interface (e.g. an http.ResponseWriter)
type Response interface {
	getBody() interface{}
}

func writeJSON(obj interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(obj)
}

// SendJSON is like writeJSON, but it writes to the log instead of returning an error.
// The caller must ensure that no writes to w happen after this function is called.
// Unwraps a Response object and converts it to an HTTP Response.
func SendJSON(obj Response, w http.ResponseWriter, log logging.Logger) {
	w.Header().Set("Content-Type", "application/json")
	err := writeJSON(obj.getBody(), w)
	if err != nil {
		log.Warnf("algod failed to write an object to the response stream: %v", err)
	}
}

// Status is an httpHandler for route GET /v1/status
func Info(ctx lib.ReqContext, w http.ResponseWriter, r *http.Request) {
	var info v1.RLInfo
	info.Details = "RandLabs Extensions to the Algorand REST API"
	response := InfoResponse{&info}
	SendJSON(response, w, ctx.Log)
}

const maxTealBytes = 1000

// DisassembleTeal is an httpHandler for route GET /v1/status
func DisassembleTeal(ctx lib.ReqContext, w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxTealBytes*2)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		lib.ErrorResponse(w, http.StatusBadRequest, err, err.Error(), ctx.Log)
		return
	}
	if len(body) == 0 {
		lib.ErrorResponse(w, http.StatusBadRequest, err, "Program Empty", ctx.Log)
		return
	}

	text, dd, err := teal.DisassembleFile(body)
	if err != nil {
		lib.ErrorResponse(w, http.StatusBadRequest, err, "Program Empty", ctx.Log)
		return
	}

	program := v1.DisassembledTeal{
		Code: text,
		Symbols: *dd,
	}
	response := DisassembleTealResponse{&program}
	SendJSON(response, w, ctx.Log)
}
