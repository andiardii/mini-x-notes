package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "notes/core"
)

func GetTagsList(c *gin.Context) {
    rows, err := core.DB.Query("SELECT id, notes_id, tag FROM tags")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var items []map[string]interface{}
    for rows.Next() {
        var id int
        var notes_id int
        var tag string
        if err := rows.Scan(&id, &notes_id, &tag); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        item := map[string]interface{}{
            "id": id,
            "notes_id": notes_id,
            "tags": tag,
        }
        items = append(items, item)
    }

    c.JSON(http.StatusOK, items)
}
