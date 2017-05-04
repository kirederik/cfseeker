package api

import (
	"net/http"

	"github.com/thomasmmitchell/cfseeker/config"
)

//MetaOutput gives meta information about this cfseeker server.
type MetaOutput struct {
	Version string `json:"version" yaml:"version"`
}

func metaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	output := MetaOutput{Version: config.Version}
	w.Write(NewResponse().AttachContents(output).Bytes())
}