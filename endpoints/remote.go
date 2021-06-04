package endpoints

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pnkj-kmr/infra-patch-manager/task"
)

// GetRemoteList method gives avaiable remotes in system
// @Description getting all remotes
// @Summary getting all available remotes
// @Tags Patch
// @Accept json
// @Produce json
// @Success 200 {string} string "{"ok":"string", "msg": "string", "data": ""}"
// ----Success 200 {array} jsn.Remote status "ok"
// @Router /api/remotes [get]
func GetRemoteList(c *fiber.Ctx) error {

	task, err := task.NewPatchTask()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
	}

	// Calling ping to all remotes with ping msg
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"msg":  "Remotes",
		"data": task.PingToAll("ping"),
	})
}

// GetRemote finds the remote in system
// @Description searching remote by name
// @Summary searching remote by name
// @Tags Patch
// @Accept json
// @Produce json
// @Success 200 {string} string "{"ok":"string", "msg": "string", "data": ""}"
// ----Success 200 {object} jsn.Remote status "ok"
// @Param name path string true "Remote Name"
// @Router /api/remote/{name} [get]
func GetRemote(c *fiber.Ctx) error {
	// getting remote name from params
	remote := c.Params("name")

	task, err := task.NewPatchTask()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
	}

	// Calling ping to all remotes with ping msg
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"msg":  "Remote Detail",
		"data": task.PingTo(remote, "ping"),
	})
}

// CheckRemotesRights to check the remote server rights
// @Description getting remotes rights of read/write
// @Summary getting remotes rights of read/write
// @Tags Patch
// @Accept json
// @Produce json
// @Success 200 {string} string "{"ok":"string", "msg": "string", "data": ""}"
// ----Success 200 {array} jsn.Remote status "ok"
// @Router /api/remotes/rights 	[get]
func CheckRemotesRights(c *fiber.Ctx) error {
	task, err := task.NewPatchTask()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
	}

	// Calling ping to all remotes with ping msg
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"msg":  "Remotes Rights",
		"data": task.RightsCheckForAll(),
	})
}

// CheckRemoteRights to check the remote server rights
// @Description getting remote rights of read/write
// @Summary getting remote rights of read/write
// @Tags Patch
// @Accept json
// @Produce json
// @Success 200 {string} string "{"ok":"string", "msg": "string", "data": ""}"
// ----Success 200 {array} jsn.Remote status "ok"
// @Param name path string true "Remote Name"
// @Param apps query []string false "Remote Apps"
// @Router /api/remote/{name}/rights 	[get]
func CheckRemoteRights(c *fiber.Ctx) error {
	// getting remote name from params
	remote := c.Params("name")

	rQuery := strings.Split(c.Query("apps"), ",")
	// rQuery := new(remoteQuery)
	// if err := c.QueryParser(rQuery); err != nil {
	// 	return err
	// }
	// log.Println("Remote name >>>", remote)
	// log.Println("Query string >>>", rQuery)

	task, err := task.NewPatchTask()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
	}

	// Calling ping to all remotes with ping msg
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"msg":  "Remote Rights",
		"data": task.RightsCheckFor(remote, rQuery),
	})
}
