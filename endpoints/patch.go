package endpoints

import (
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/pnkj-kmr/infra-patch-manager/actions"
	"github.com/pnkj-kmr/infra-patch-manager/utility"
)

// UploadPatch method helps to upload the patch to given remote server
// @Description uploading the patch to server
// @Summary uploading file to server
// @Tags Patch
// @Accept json
// @Produce json
// ---Success 200 {string} string "{"ok":"string", "msg": "string", "data": ""}"
// ---Success 200 {array} jsn.Remote status "ok"
// @Param name path string true "Remote Name"
// @Router /api/upload/patch/{name} [post]
func UploadPatch(c *fiber.Ctx) error {
	// getting remote name from params
	remote := c.Params("name")

	// Get first file from form field "document":
	file, err := c.FormFile("file_upload")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
	}

	// Default upload location of file
	patchpath := filepath.Join(utility.AssetsDirectory, file.Filename)

	// Save the files to disk:
	if err := c.SaveFile(file, patchpath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
	}

	log.Println("File upload api request with data:", remote, patchpath)
	action, err := actions.NewAction()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
	}

	if remote == "all" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":   true,
			"msg":  "File Upload",
			"data": action.PatchFileUploadToAll(patchpath),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"msg":  "File Upload",
		"data": action.PatchFileUploadTo(remote, patchpath),
	})
}

// ApplyPatch help to apply patch at given servers/server
// @Description apply patch to remote server app(s)
// @Summary apply patch to remote server app(s)
// @Tags Patch
// @Accept json
// @Produce json
// ---Success 200 {string} string "{"ok":"string", "msg": "string", "data": ""}"
// ---Success 200 {array} jsn.Remote status "ok"
// @Router /api/apply/patch [post]
func ApplyPatch(c *fiber.Ctx) error {
	b := new(patchbody)
	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(NewErrResponse(err))
	}

	log.Println("Apply patch with data filter", b)
	action, err := actions.NewAction()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
	}

	if b.Remote == "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":   true,
			"msg":  "Patch Applied",
			"data": action.ApplyPatchToAll(b.Apptype),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"msg":  "Patch Applied",
		"data": action.ApplyPatchTo(b.Remote, b.Apptype),
	})
}

// VerifyPatch method helps to verify the applied patch
// @Description verifying the applied server app patch
// @Summary verifying the applied server app patch
// @Tags Patch
// @Accept json
// @Produce json
// ---Success 200 {string} string "{"ok":"string", "msg": "string", "data": ""}"
// ---Success 200 {array} jsn.Remote status "ok"
// @Router /api/verify/patch [post]
func VerifyPatch(c *fiber.Ctx) error {
	b := new(patchbody)
	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(NewErrResponse(err))
	}

	log.Println("Verify patch with data filter", b)
	action, err := actions.NewAction()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
	}

	if b.Remote == "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":   true,
			"msg":  "Patch Verified",
			"data": action.VerifyPatchToAll(b.Apptype),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"msg":  "Patch Verified",
		"data": action.VerifyPatchTo(b.Remote, b.Apptype),
	})
}
