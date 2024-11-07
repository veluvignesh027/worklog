package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r Router) HandleLogout() {
	r.mux.POST("/logout", func(ctx *gin.Context) {
		ctx.String(http.StatusAccepted, "Logged out sucessfully!")
	})
}

func (r Router) HandleHelpPage() {
	r.mux.GET("/help", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Documents related to the apps")
	})
}
