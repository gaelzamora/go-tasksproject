package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gaelzamora/go-rest-crud/internal/application"
	"github.com/gaelzamora/go-rest-crud/internal/domain"
	"github.com/gorilla/mux"
)

type TaskHandler struct {
	service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	tasks, err := h.service.GetAllTasksByUser(userID)
	if err != nil {
		http.Error(w, "Error obteniendo tareas", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var task domain.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	createdTask, err := h.service.CreateTask(userID, &task) // Pasar un puntero
	if err != nil {
		http.Error(w, "Error creando tarea", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Created Task: %+v\n", createdTask) // Depuración

	w.Write([]byte("Tarea creada"))
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	existingTask, err := h.service.GetTaskByID(uint(id))
	if err != nil {
		http.Error(w, "Tarea no encontrada", http.StatusNotFound)
		return
	}

	var updatedData domain.Task
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	existingTask.Name = updatedData.Name
	existingTask.Content = updatedData.Content

	err = h.service.UpdateTask(existingTask)
	if err != nil {
		http.Error(w, "Error actualizando la tarea", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingTask)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTaskByID(uint(id))
	if err != nil {
		fmt.Println("Tarea no encontrada")
		http.Error(w, "Tarea no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(uint(id))
	if err != nil {
		http.Error(w, "Error eliminando tarea", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
