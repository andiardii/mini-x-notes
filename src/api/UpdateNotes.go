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

    text, hashtags := core.SeparateTextAndHashtags(input.Note)
    // c.JSON(http.StatusInternalServerError, gin.H{"error4": hashtags})
    // return

	// hardcode
	userId := 1

    if input.ID == 0 {
        result, err := core.DB.Exec("INSERT INTO notes (user_id, note) VALUES (?, ?)", userId, text)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
            return
        }

        id, _ := result.LastInsertId()

        for _, hashtag := range hashtags {
            _, err = core.DB.Exec("INSERT INTO tags (notes_id, tag) VALUES (?, ?)", id, hashtag)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
                return
            }
        }

        c.JSON(http.StatusOK, gin.H{"message": "Note created", "id": id})

    } else {
        _, err := core.DB.Exec("UPDATE notes SET user_id = ?, note = ? WHERE id = ?", userId, text, input.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
            return
        }

        core.DB.Exec("DELETE FROM tags WHERE notes_id = ?", input.ID)
        for _, hashtag := range hashtags {
            _, err = core.DB.Exec("INSERT INTO tags (notes_id, tag) VALUES (?, ?)", input.ID, hashtag)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
                return
            }
        }

        c.JSON(http.StatusOK, gin.H{"message": "Note updated"})
    }
}
