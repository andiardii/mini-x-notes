package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "notes/core"
    "strconv"
    "database/sql"
)

func GetNotesByUser(c *gin.Context) {
    idParam := c.Param("userId")
    
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    rows, err := core.DB.Query(`
        SELECT n.id, n.user_id, n.note, GROUP_CONCAT(t.tag ORDER BY t.tag SEPARATOR ',') AS tags
        FROM notes n
        LEFT JOIN tags t ON t.notes_id = n.id
        WHERE n.user_id = ?
        GROUP BY n.id, n.note
    `, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var items []map[string]interface{}
    for rows.Next() {
		var id int
        var user_id int
        var note string
        var tags sql.NullString

        if err := rows.Scan(&id, &user_id, &note, &tags); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        item := map[string]interface{}{
			"id": id,
            "user_id": user_id,
            "notes": note,
            "tags": tags,
        }
        
        if !tags.Valid {
            item["tags"] = nil
        } else {
            item["tags"] = tags.String
        }

        items = append(items, item)
    }

    c.JSON(http.StatusOK, items)
}
