package handlers

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"example.com/vb/components/base"
	"example.com/vb/database"
	"github.com/gin-gonic/gin"
)

func BaseHTMLHandler(r *gin.RouterGroup) {
	r.GET("/home", func(ctx *gin.Context) {
		title := os.Getenv("HTML_TITLE")
		icon := os.Getenv("ICON_PATH")
		base.Base(title, icon).Render(context.Background(), ctx.Writer)
	})
}

func AddUserStoryHandler(r *gin.RouterGroup) {
	r.POST("/add-user-story", func(ctx *gin.Context) {

		title := ctx.Param("title")
		assignee := ctx.Param("assignee")

		tmp := database.UserStory{
			Number:     database.UserStoryCounter,
			Title:      title,
			Assignee:   assignee,
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}

		database.UserStoryCounter++

		contentDiv := `<div class="task-list">
					    <h2> <a href="#"> US-0` + strconv.Itoa(tmp.Number) + "\t:\t" + tmp.Title +
			`</div>`
		ctx.Request.Header.Add("createdAt", tmp.CreatedAt.String())
		ctx.Request.Header.Add("modifiedAt", tmp.ModifiedAt.String())
		ctx.Request.Header.Add("assignedTo", tmp.Assignee)
		ctx.String(http.StatusOK, contentDiv)
	})
}
