package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/Pwawazi/lms-saas/internal/user"
    "github.com/Pwawazi/lms-saas/config"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    cfg := config.LoadConfig()

    db := config.InitDB(cfg)
    userRepo := user.NewRepository(db)
    userService := user.NewService(userRepo)
    userHandler := user.NewHandler(userService)

    e.POST("/users", userHandler.CreateUser)
    // Other routes...

    e.Logger.Fatal(e.Start(":8080"))
}
