package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	mux *gin.Engine
}

func InitHandlers(route *gin.Engine) error {
	router := Router{
		mux: route,
	}

	authgrp := router.mux.Group("/")
	authgrp.Use(AuthenticateUserInfo())

	router.HandleRoot()
	router.LoginHandler()
	router.SignUpHandler()
	router.ResourceHandler()
	router.CreateUser()

	BaseHTMLHandler(authgrp)
	AddUserStoryHandler(authgrp)
	FormHandler(authgrp)
	return nil
}

func (r Router) HandleRoot() {
	r.mux.GET("/", func(ctx *gin.Context) {

		// _, err := auth.VerifyToken(ctx.Cookie(""))
		// if err != nil {
		// 	ctx.Redirect(http.StatusTemporaryRedirect, "/login.html")
		// }

		ctx.Redirect(http.StatusTemporaryRedirect, "/home")
	})
}

func (r Router) ResourceHandler() {
	r.mux.GET("/resources/:file", func(ctx *gin.Context) {
		fname := ctx.Param("file")
		if fname == "image-1.png" {
			nByte, err := os.ReadFile("/Users/vignesh0/practice/resources/image-1.png")
			if err != nil {
				log.Println(err)
			}

			ctx.Writer.Write(nByte)
		} else {
			ctx.Status(http.StatusNotFound)
		}
	})
}
