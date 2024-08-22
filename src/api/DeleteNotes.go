package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "notes/core"
)

func DeleteNotes(c *gin.Context) {
    id := c.Param("id")

    result, err := core.DB.Exec("DELETE FROM notes WHERE id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Note successfully deleted"})
}