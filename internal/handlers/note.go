package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"tg_bot/internal/service"
	"tg_bot/pkg/lib"
)

type NoteHandler struct {
	noteService service.Note
	logger      lib.Logger
	http        chi.Router
}

func NewNoteRoutes(r chi.Router, noteService service.Note, logger lib.Logger) {
	nh := &NoteHandler{noteService: noteService, logger: logger, http: r}
	r.Get("/notes", nh.GetAll)
	r.Post("/notes", nh.Create)
	r.Delete("/notes/{id}", nh.Delete)
	r.Put("/notes/{id}", nh.Update)
}

// Create godoc
//
// @Summary Create note
// @Description Create note
// @ID create-note
// @Tags Notes
// @Success 201
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /api/notes [post]
func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateNoteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}
	err := h.noteService.Create(r.Context(), input.Title, input.Description, input.RemindsAt)
	if err != nil {
		h.logger.Error("error while creating note: %w", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAll godoc
//
// @Summary Get all notes
// @Description Get all notes except deleted
// @ID get-all-note
// @Tags Notes
// @Produce json
// @Success 200 {object} NotesResponse
// @Failure 500 {object} error
// @Router /api/notes [get]
func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	notes, err := h.noteService.GetAll(r.Context())
	if err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	resp := NotesResponse{Notes: notes}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

// Delete godoc
//
// @Summary Delete note
// @Description Delete note by id
// @Param id path int true "Note id"
// @ID delete-note
// @Tags Notes
// @Success 204
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /api/notes/{id} [delete]
func (h *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var idStr string
	if idStr = chi.URLParam(r, "id"); idStr == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	}

	err = h.noteService.Delete(r.Context(), id)
	if err != nil {
		h.logger.Error(err)
		if err == service.ErrNoteNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

// Update godoc
//
// @Summary Update note
// @Description Update note by id
// @Param id path int true "Note id"
// @Param request body CreateNoteInput true "Input data to update note"
// @ID update-note
// @Tags Notes
// @Produce json
// @Success 200 {object} UpdatedNoteResponse
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /api/notes/{id} [put]
func (h *NoteHandler) Update(w http.ResponseWriter, r *http.Request) {
	var idStr string
	if idStr = chi.URLParam(r, "id"); idStr == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	}

	var input CreateNoteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	note, err := h.noteService.Update(r.Context(), id, input.Title, input.Description, input.RemindsAt)
	if err != nil {
		h.logger.Error(err)
		if err == service.ErrNoteNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	noteResp := UpdatedNoteResponse{note}
	if err = json.NewEncoder(w).Encode(noteResp); err != nil {

		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}
