package main

import (
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/menuser/go-fiber-crm-basic/database"
	"github.com/menuser/go-fiber-crm-basic/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	database.Connect()
	database.GetDB()
	database.DBConn.AutoMigrate(&lead.Lead{})
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close() //The key word 'defer' usually used to close certain resources
}
