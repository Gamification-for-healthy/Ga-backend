package routes

import 
(
"net/http"
"Ga-backend/config"
)
func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}
}