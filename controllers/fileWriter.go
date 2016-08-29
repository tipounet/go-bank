package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/tipounet/go-bank/configuration"
)

func writeResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	conf := configuration.GetConfiguration()

	if conf.Prettyprint == true {
		b, _ := json.MarshalIndent(data, "", "    ")
		w.Write(b)
	} else {
		json.NewEncoder(w).Encode(data)
	}
}
