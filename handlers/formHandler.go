package handlers

import (
	"context"
	"net/http"

	"example.com/vb/components/form"
	"github.com/gin-gonic/gin"
)

func FormHandler(r *gin.RouterGroup) {
	r.GET("/getform/:id", func(ctx *gin.Context) {
		formType := ctx.Param("id")

		if formType == "userstory" {
			form.UserStoryForm().Render(context.Background(), ctx.Writer)
		} else if formType == "worklogForm" {
			form.UserStoryForm().Render(context.Background(), ctx.Writer)
		} else if formType == "subworklogForm" {
			form.UserStoryForm().Render(context.Background(), ctx.Writer)
		} else {
			ctx.String(http.StatusBadRequest, "Form Not applicable")
		}
	})
}
