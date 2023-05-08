package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	app := fiber.New()
	SetupRoutes(app)
	SetupDatabase()
	db.AutoMigrate(&Data{})
	app.Listen(":8080")

}

func SetupDatabase() error {
	dbUri := "postgres://a:a@localhost:5432/a" // used for this project else should be used in .env
	var err error
	db, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func SetupRoutes(a *fiber.App) {
	a.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "online"})
	})
	a.Get("/getAllData", func(c *fiber.Ctx) error {
		var AllData []Data
		db.Find(&AllData)
		return c.JSON(AllData)
	})
	a.Post("/addData", func(c *fiber.Ctx) error {
		d := new(Data)
		if err := c.BodyParser(d); err != nil {
			c.Status(500).JSON(fiber.Map{"error": "wrongformat"})
		}
		db.Create(&d)
		return c.JSON(d)
	})
	a.Put("/updateData/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		d := new(Data)
		db.First(&d, id)
		return c.JSON(d)

	})
	a.Delete("/deleteData/:id", func(c *fiber.Ctx) error {
		d := new(Data)
		id := c.Params("id")
		db.Delete(&d, id)
		return c.JSON(fiber.Map{
			"status": "deleted",
		})
	})

}
