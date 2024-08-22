package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "notes/core"
)

type NoteInput struct {
    ID    int    `json:"id,omitempty"`
    Note  string `json:"note" binding:"required"`
}

func UpdateNotes(c *gin.Context) {
    var input NoteInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// hardcode
	userId := 1

    if input.ID == 0 {
        result, err := core.DB.Exec("INSERT INTO notes (user_id, note) VALUES (?, ?)", userId, input.Note)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        id, _ := result.LastInsertId()
        c.JSON(http.StatusOK, gin.H{"message": "Note created", "id": id})

    } else {
        _, err := core.DB.Exec("UPDATE notes SET user_id = ?, note = ? WHERE id = ?", userId, input.Note, input.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Note updated"})
    }
}
