package hello

import (
	"encoding/json"
	"net/http"
)

func (h *HelloHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := h.helloRepository.Hello()
	json.NewEncoder(w).Encode(data)
}
