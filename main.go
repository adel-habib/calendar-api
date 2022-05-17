package main

import (
	"calendar-api/holidays"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := mysql.Config{
		DSN: "root:test@tcp(db:3306)/calendar?charset=utf8mb4&parseTime=True&loc=Local",
	}
	var err error
	db, err := gorm.Open(mysql.New(cfg), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	err = db.AutoMigrate(&holidays.Region{}, &holidays.Holiday{})
	if err != nil {
		fmt.Println(err)
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		reg1 := holidays.Region{Name: "CusName", ShortName: "CusShortname", ParentRegion: nil}
		db.Create(&reg1)
		fmt.Println("INSERTED INTO DB!")
		reg2 := holidays.Region{
			Name:         "Some Name",
			ShortName:    "ssn",
			ParentRegion: &reg1,
		}
		db.Create(&reg2)
		fmt.Println("Created again")
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "Something"})
	})

	e.GET("/regions", func(c echo.Context) error {
		var regions []holidays.Region
		parent := c.QueryParam("parent")
		if parent == "" {
			db.Find(&regions)
		} else {
			var parentEntity holidays.Region
			db.Where("name = ?", parent).First(&parentEntity)
			fmt.Println(parentEntity)
			db.Where("parent_id = ?", parentEntity.Id).Find(&regions)
		}
		return c.JSON(http.StatusOK, regions)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
