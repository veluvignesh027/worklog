package handlers

import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"example.com/vb/auth"
	"example.com/vb/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

func (r Router) LoginHandler() {
	r.mux.GET("/login.html", func(ctx *gin.Context) {
		nByte, err := os.ReadFile("/Users/vignesh0/practice/components/HTML/login.html")
		if err != nil {
			log.Println(err)
		}
		ctx.Header("Content-type", "text/html")
		ctx.Writer.Write(nByte)
	})
}

func AuthenticateUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth.VerifyToken("email-placeholder")
		ctx.Next()
	}
}

func (r Router) SignUpHandler() {
	r.mux.GET("/signup.html", func(ctx *gin.Context) {
		nByte, err := os.ReadFile("/Users/vignesh0/practice/components/HTML/signup.html")
		if err != nil {
			log.Println(err)
		}
		ctx.Header("Content-type", "text/html")
		ctx.Writer.Write(nByte)
	})
}

func (r Router) CreateUser() {
	r.mux.POST("/create-user", func(ctx *gin.Context) {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Println(err)
		}

		params := GetFromForm(string(body))

		tmp := database.User{
			FName:      params["firstname"],
			LName:      params["lastname"],
			Password:   params["password"],
			Email:      params["username"],
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
			CustomerId: string(rand.New),
		}

	})
}

func GetFromForm(str string) map[string]string {
	params := make(map[string]string)

	splitstrs := strings.Split(str, "&")

	for _, i := range splitstrs {
		before, mid, _ := strings.Cut(i, "=")
		params[before] = mid

	}
	log.Println(params)
	return params
}
