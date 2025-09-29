package handlers

import "net/http"

// CORSMiddleware maneja las políticas CORS para permitir peticiones desde diferentes orígenes
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Orígenes permitidos
		allowedOrigins := map[string]bool{
			"http://localhost:4200": true, // Angular en desarrollo local
			"http://localhost:80":   true, // Angular en Docker
			"http://localhost":      true, // Angular en Docker (sin puerto explícito)
		}
		
		origin := r.Header.Get("Origin")
		
		// Si el origen está en la lista de permitidos, configurar CORS
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Manejar solicitudes preflight (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// Continuar con el handler normal
		next(w, r)
	}
}