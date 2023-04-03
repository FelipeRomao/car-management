package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Name  string `json:"name"`
	Model string `json:"model"`
}

// Creates a server REST (GET, POST) request - START

var cars []Car

func createCars() {
	cars = append(cars, Car{Name: "Ferrari", Model: "Ferrari"})
	cars = append(cars, Car{Name: "Onix", Model: "Chevrolet"})
	cars = append(cars, Car{Name: "Argo", Model: "Fiat"})
}

func main() {
	createCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	car := new(Car)

	if err := c.Bind(car); err != nil {
		return err
	}

	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

// Creates a server REST (GET, POST) request - END

// Connection on database MySQL for creates car - START
func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")

	if err != nil {
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, model) VALUES ($1, $2)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name, car.Model)

	if err != nil {
		return err
	}

	return nil
}

// Connection on database MySQL for creates car - END
