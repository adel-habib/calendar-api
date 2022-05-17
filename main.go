package main

import (
	"calendar-api/model"
	"calendar-api/persistance"
	"calendar-api/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
)

func main() {

	dsn := "root:test@tcp(db:3306)/calendar?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := persistance.DBconnection(dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	err = db.AutoMigrate(&model.Region{}, &model.Holiday{})
	if err != nil {
		fmt.Println(err)
	}
	e := echo.New()
	for i := 2000; i < 2030; i++ {
		service.CalculateGermanHolidays(i)
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		reg1 := model.Region{Name: "CusName", ShortName: "CusShortname", ParentRegion: nil}
		db.Create(&reg1)
		fmt.Println("INSERTED INTO DB!")
		reg2 := model.Region{
			Name:         "Some Name",
			ShortName:    "ssn",
			ParentRegion: &reg1,
		}
		db.Create(&reg2)
		fmt.Println("Created again")
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "Something"})
	})

	e.GET("/regions", func(c echo.Context) error {
		var regions []model.Region
		parent := c.QueryParam("parent")
		if parent == "" {
			db.Find(&regions)
		} else {
			var parentEntity model.Region
			db.Where("name = ?", parent).First(&parentEntity)
			fmt.Println(parentEntity)
			db.Where("parent_id = ?", parentEntity.Id).Find(&regions)
		}
		return c.JSON(http.StatusOK, regions)
	})
	e.GET("/geth", func(c echo.Context) error {
		var h []model.Holiday
		persistance.DB.Preload("Regions").Find(&h)
		for _, item := range h {
			fmt.Println(item)
		}
		return c.NoContent(200)
	})

	e.GET("/geth", func(c echo.Context) error {
		h := persistance.GetHolidayById(1)
		return c.JSON(http.StatusOK, h)
	})

	e.GET("/gg", func(c echo.Context) error {
		h := persistance.GetHolidaysByRegionNameAndYear("Deutschland", 2022)
		return c.JSON(http.StatusOK, h)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
