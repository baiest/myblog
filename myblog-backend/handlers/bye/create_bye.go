package bye

import (
	"encoding/json"
	"net/http"
)

func (b *ByeHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := b.byeRepository.Bye(r.Context())
	json.NewEncoder(w).Encode(data)
}
