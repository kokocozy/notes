package main

import (
	"log"
	"net/http"
	"notes-api/auth"
	"notes-api/handler"
	"notes-api/helper"
	"notes-api/note"
	"notes-api/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:halo123@tcp(127.0.0.1:3306)/db_notes?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	noteRepository := note.NewRepository(db)
	userRepository := user.NewRepository(db)

	authService := auth.NewService()
	noteService := note.NewService(noteRepository)
	userService := user.NewService(userRepository)

	noteHandler := handler.NewNoteHandler(noteService)
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	routerUser := api.Group("user/")
	routerNote := api.Group("note/")

	routerUser.POST("create", noteHandler.CreateNote)
	routerUser.POST("register", userHandler.RegisterUser)
	routerUser.POST("login", userHandler.LoginUser)
	routerUser.GET("profile", authMiddleware(authService, userService), userHandler.Profile)
	routerUser.POST("profile", authMiddleware(authService, userService), userHandler.UpdateDataProfile)
	routerUser.POST("change-password", authMiddleware(authService, userService), userHandler.UpdatePasswordProfile)

	routerNote.POST("create", authMiddleware(authService, userService), noteHandler.CreateNote)
	routerNote.GET("mynote", authMiddleware(authService, userService), noteHandler.MyNotes)
	routerNote.GET("mynote/:id", authMiddleware(authService, userService), noteHandler.GetNoteById)
	routerNote.PUT("mynote/:id", authMiddleware(authService, userService), noteHandler.UpdateDataNote)
	routerNote.DELETE("mynote/:id", authMiddleware(authService, userService), noteHandler.DeleteDataNote)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenString = tokenArray[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
