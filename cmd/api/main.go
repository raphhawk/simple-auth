package main

import (
	"fmt"
	"os"
	"simple-auth/internal/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func main() {
	// Setup Database
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Failed to connect to database")
		os.Exit(1)
	}
	defer db.Close()

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
