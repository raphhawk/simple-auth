package main

import (
	"fmt"
	"os"
	"simple-auth/internal/models"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Setup Database
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		os.Exit(1)
	}
	// defer db.Close()

	db.AutoMigrate(&models.User{})

	// Setup Echo
	e := echo.New()

	// Routes
	e.POST("/signup", SignUp(db))
	e.POST("/signin", SignIn(db))
	e.POST("/refresh", RefreshToken(db))
	e.POST("/revoke", RevokeToken)

	e.Logger.Fatal(e.Start(":8080"))
}
