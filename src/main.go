package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "notes/core"
    "notes/api"
	"net/http"
)

func main() {
    core.InitializeDatabase()
    defer core.DB.Close()

    r := gin.Default()

	r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

	r.GET("/getNotesList", api.GetNotesList)
	r.GET("/getNotesByUser/:userId", api.GetNotesByUser)
	r.GET("/getNotesById/:id", api.GetNotesById)
	r.POST("/updateNotes", api.UpdateNotes)
	r.DELETE("/deleteNotes/:id", api.DeleteNotes)

	// r.GET("/getUser/:id", api.GetUser)
	// r.POST("/updateUser/:idUser", api.UpdateUser)
	// r.DELETE("/deleteUser/:id", api.DeleteUser)

	r.GET("/getTagsList", api.GetTagsList)
	r.GET("/getTagsNotes/:notesId", api.GetTagsNotes)

    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
