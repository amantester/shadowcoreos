package handlers

import (
	"github.com/amantester/shadowcoreos/apps/core-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	ProjectService *services.ProjectService
}

func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		ProjectService: projectService,
	}
}

func (h *ProjectHandler) Create(c *fiber.Ctx) error {
	type Request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var body Request

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	project, err := h.ProjectService.Create(
		body.Name,
		body.Description,
		c.Locals("user_id").(string),
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(project)
}

func (h *ProjectHandler) List(c *fiber.Ctx) error {
	projects, err := h.ProjectService.List()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(projects)
}
