package main

import (
	"log"
	"net/http"

	"personal-app-back/handlers"
	"personal-app-back/models"
)

func main() {
	// Inicializar datos de ejemplo
	models.InitHeroData()
	models.InitFeatureData()

	// Configurar rutas con middleware CORS
	http.HandleFunc("/hero-sections", handlers.CORSMiddleware(handlers.HeroSectionsHandler))
	http.HandleFunc("/hero-sections/", handlers.CORSMiddleware(handlers.HeroSectionHandler))
	http.HandleFunc("/features", handlers.CORSMiddleware(handlers.FeaturesHandler))
	http.HandleFunc("/features/", handlers.CORSMiddleware(handlers.FeatureHandler))

	log.Println("🚀 Servidor iniciado en http://localhost:8080")
	log.Println("✅ CORS habilitado para múltiples orígenes")
	log.Println("📍 Endpoints disponibles:")
	log.Println("   Hero Sections:")
	log.Println("     - GET    /hero-sections     (Listar todos)")
	log.Println("     - POST   /hero-sections     (Crear nuevo)")
	log.Println("     - GET    /hero-sections/:id (Obtener uno)")
	log.Println("     - PUT    /hero-sections/:id (Actualizar)")
	log.Println("     - DELETE /hero-sections/:id (Eliminar)")
	log.Println("   Features:")
	log.Println("     - GET    /features          (Listar todos)")
	log.Println("     - POST   /features          (Crear nuevo)")
	log.Println("     - GET    /features/:id      (Obtener uno)")
	log.Println("     - PUT    /features/:id      (Actualizar)")
	log.Println("     - DELETE /features/:id      (Eliminar)")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}