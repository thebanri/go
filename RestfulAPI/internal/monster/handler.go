package monster

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Just Allowed GET method", http.StatusMethodNotAllowed)
		return 
	}
	


}
