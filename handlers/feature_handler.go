package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"personal-app-back/models"
)

// FeaturesHandler maneja las peticiones GET (listar) y POST (crear) para features
func FeaturesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Listar todos los features
		models.FeatureMu.Lock()
		featureList := make([]models.Feature, 0, len(models.Features))
		for _, feature := range models.Features {
			featureList = append(featureList, feature)
		}
		models.FeatureMu.Unlock()

		json.NewEncoder(w).Encode(featureList)

	case "POST":
		// Crear nuevo feature
		var feature models.Feature
		if err := json.NewDecoder(r.Body).Decode(&feature); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validación básica
		if feature.Title == "" {
			http.Error(w, "El título es requerido", http.StatusBadRequest)
			return
		}

		models.FeatureMu.Lock()
		feature.ID = models.FeatureID
		models.FeatureID++
		models.Features[feature.ID] = feature
		models.FeatureMu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(feature)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// FeatureHandler maneja las peticiones GET (obtener), PUT (actualizar) y DELETE (eliminar) para un feature específico
func FeatureHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extraer ID de la URL
	idStr := r.URL.Path[len("/features/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		// Obtener feature específico
		models.FeatureMu.Lock()
		feature, exists := models.Features[id]
		models.FeatureMu.Unlock()

		if !exists {
			http.Error(w, "Feature no encontrado", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(feature)

	case "PUT":
		// Actualizar feature
		models.FeatureMu.Lock()
		_, exists := models.Features[id]
		models.FeatureMu.Unlock()

		if !exists {
			http.Error(w, "Feature no encontrado", http.StatusNotFound)
			return
		}

		var feature models.Feature
		if err := json.NewDecoder(r.Body).Decode(&feature); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validación básica
		if feature.Title == "" {
			http.Error(w, "El título es requerido", http.StatusBadRequest)
			return
		}

		feature.ID = id
		models.FeatureMu.Lock()
		models.Features[id] = feature
		models.FeatureMu.Unlock()

		json.NewEncoder(w).Encode(feature)

	case "DELETE":
		// Eliminar feature
		models.FeatureMu.Lock()
		_, exists := models.Features[id]
		if exists {
			delete(models.Features, id)
		}
		models.FeatureMu.Unlock()

		if !exists {
			http.Error(w, "Feature no encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}