package http

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type cacheHandler struct {
	*Server
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodPut:
		value, _ := ioutil.ReadAll(r.Body)
		if len(value) != 0 {
			if err := h.Set(key, value); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	case http.MethodGet:
		value, err := h.Get(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(value) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(value)
		return
	case http.MethodDelete:
		if err := h.Del(key); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
