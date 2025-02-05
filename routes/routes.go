package routes

import (
	"github.com/AKAZJAYA/blogbackend/controller"
	"github.com/AKAZJAYA/blogbackend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	// Add CORS middleware with secure configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",  // Specify exact origin
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	// Public routes
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	// Protected routes group
	api := app.Group("/api")
	api.Use(middleware.IsAuthenticate)

	// Protected endpoints
	app.Post("/api/uploads-image", controller.Upload)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/allpost/:id", controller.DetailPost)
	app.Get("/api/uniquepost", controller.UniqePost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)

	// Static file serving
	app.Static("/api/uploads", "./uploads")
}