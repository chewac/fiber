package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	// => http://localhost:3000/hello.html
	// => http://localhost:3000/js/jquery.js
	// => http://localhost:3000/css/style.css

	// serve from multiple directories
	app.Static("/", "./files")

	app.Listen(":3000")
}
