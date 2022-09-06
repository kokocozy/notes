package handler

import (
	"net/http"
	"notes-api/helper"
	"notes-api/note"
	"notes-api/user"

	"github.com/gin-gonic/gin"
)

type noteHandler struct {
	noteService note.Service
}

func NewNoteHandler(noteService note.Service) *noteHandler {
	return &noteHandler{noteService}
}

func (h *noteHandler) CreateNote(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	var input note.NewNoteInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Input validation error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newNote, err := h.noteService.CreateNote(currentUser.ID, input)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Error create new note", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := note.FormatNote(newNote)
	response := helper.APIResponse("Success create new note", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *noteHandler) MyNotes(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	notes, err := h.noteService.AllNoteByUserId(currentUser.ID)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Error get note", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := note.FormatNotes(notes)
	response := helper.APIResponse("Success get all notes", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *noteHandler) GetNoteById(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	var input note.GetNoteIdInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Input validation error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	noteFinded, err := h.noteService.FindNote(input.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Error get note", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if noteFinded.UserId != currentUser.ID {
		errors := gin.H{"errors": "Note is not yours"}

		response := helper.APIResponse("Error get note", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := note.FormatNote(noteFinded)
	response := helper.APIResponse("Note found", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *noteHandler) UpdateDataNote(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	var input note.UpdateNoteInput
	var idNote note.GetNoteIdInput

	err := c.ShouldBindUri(&idNote)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Validation uri error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Validation input error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.User = currentUser
	updatedNote, err := h.noteService.UpdateNote(idNote.Id, input)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Error update note", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := note.FormatNote(updatedNote)
	response := helper.APIResponse("Success update note", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *noteHandler) DeleteDataNote(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	var idNote note.GetNoteIdInput

	err := c.ShouldBindUri(&idNote)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Validation uri error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedNote, err := h.noteService.DeleteNote(idNote.Id, currentUser.ID)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Error deleting note", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := note.FormatNote(deletedNote)
	response := helper.APIResponse("Success deleting note", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
