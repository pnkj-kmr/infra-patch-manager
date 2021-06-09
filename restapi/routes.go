package restapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pnkj-kmr/infra-patch-manager/endpoints"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

// Reference url for fiber getstarted
// https://github.com/koddr/tutorial-go-fiber-rest-api

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
	route.Post(endpoints.APIUploadPatch, endpoints.UploadPatch)            // upload file to master server
	route.Post(endpoints.APIApplyPatch, endpoints.ApplyPatch)              // apply patch at desire location
	route.Post(endpoints.APIVerifyPatch, endpoints.VerifyPatch)            // verify patch at desire location
}

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
