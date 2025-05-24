package main

import (
	"encoding/json"
	"net/http"
)

func makeHandler(reg *Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleAddService(w, r, reg)
		case http.MethodGet:
			handleGetService(w, r, reg)
		case http.MethodDelete:
			handleDeleteService(w, r, reg)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleAddService(w http.ResponseWriter, r *http.Request, reg *Registry) {
	defer r.Body.Close()
	var svc Service
	if err := json.NewDecoder(r.Body).Decode(&svc); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if svc.Name == "" || svc.ID == "" {
		http.Error(w, "Missing service name or ID", http.StatusBadRequest)
		return
	}

	reg.Add(svc)
	w.WriteHeader(http.StatusCreated)
}

func handleGetService(w http.ResponseWriter, r *http.Request, reg *Registry) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing 'name' query parameter", http.StatusBadRequest)
		return
	}

	services := reg.Get(name)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func handleDeleteService(w http.ResponseWriter, r *http.Request, reg *Registry) {
	name := r.URL.Query().Get("name")
	id := r.URL.Query().Get("id")
	if name == "" || id == "" {
		http.Error(w, "Missing 'name' or 'id' query parameter", http.StatusBadRequest)
		return
	}

	reg.Remove(name, id)
	w.WriteHeader(http.StatusNoContent)
}
