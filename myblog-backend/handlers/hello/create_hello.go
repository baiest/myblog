package hello

import (
	"encoding/json"
	"net/http"
)

func (h *HelloHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := h.helloRepository.Hello()
	json.NewEncoder(w).Encode(data)
}
