package restapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pnkj-kmr/infra-patch-manager/endpoints"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/apidocs")
	// Routes for GET method:
	route.Get("*", swagger.Handler) // get one user by ID
}

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api")

	route.Get(endpoints.APIGetStatus, endpoints.GetStatus)                 // status check
	route.Get(endpoints.APIGetRemotes, endpoints.GetRemoteList)            // remote list
	route.Get(endpoints.APIGetRemote, endpoints.GetRemote)                 // remote search by name
	route.Get(endpoints.APIGetRemotesRights, endpoints.CheckRemotesRights) // remotes rights
	route.Get(endpoints.APIGetRemoteRights, endpoints.CheckRemoteRights)   // remote rights by name
}

// // PrivateRoutes func for describe group of private routes.
// func PrivateRoutes(a *fiber.App) {
// 	// Create routes group.
// 	route := a.Group("/api/v1")

// 	// Routes for POST method:
// 	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook) // create a new book

// 	// Routes for PUT method:
// 	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update one book by ID

// 	// Routes for DELETE method:
// 	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
// }

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		// Anonimus function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, endpoint is not found",
			})
		},
	)
}
