package http

import (
	"encoding/json"
	"net/http"
)

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data, err := json.Marshal(h.GetStat())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
