package gofraapp

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gofra/core"
)

func init() {
	core.RegisterApp(&core.AppDefinition{
		Name:        "gofraapp",
		Title:       "GoFra Core App",
		Version:     "0.16.0",
		Publisher:   "Coatlbit",
		Description: "Foundational core application providing authentication, system DocTypes, Desk UI, and RPC APIs",
		License:     "MIT",

		BeforeInstall: func() error {
			log.Println("[gofraapp] Initializing GoFra core application setup...")
			return nil
		},
		AfterInstall: func() error {
			log.Println("[gofraapp] Successfully initialized GoFra core application.")
			return nil
		},
		BeforeMigrate: func() error {
			log.Println("[gofraapp] Preparing schema migrations for GoFra core DocTypes...")
			return nil
		},
		AfterMigrate: func() error {
			log.Println("[gofraapp] Core schema migrations completed successfully.")
			return nil
		},

		DocEvents: map[string]map[string]core.DocHookFunc{
			"*": {
				"after_save": func(doc map[string]interface{}) error {
					return nil
				},
			},
			"User": {
				"validate": func(doc map[string]interface{}) error {
					return nil
				},
			},
			"Activity Log": {
				"after_insert": func(doc map[string]interface{}) error {
					return nil
				},
			},
		},

		RegisterRoutes: func(r fiber.Router) {
			group := r.Group("/api/v1/method/gofraapp")
			group.Get("/ping", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"status":  "ok",
					"app":     "gofraapp",
					"version": "0.16.0",
					"message": "GoFra Core App running",
				})
			})

			group.Get("/info", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"app_name": "gofraapp",
					"title":    "GoFra Core App",
					"version":  "0.16.0",
					"status":   "active",
					"modules": []string{
						"core", "desk", "contacts", "email",
						"workflow", "automation", "integrations",
						"custom", "geo", "printing", "website",
					},
				})
			})

			group.Get("/status", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"app":            "gofraapp",
					"status":         "healthy",
					"doctypes_count": 361,
					"modules_count":  11,
				})
			})
		},
	})
}
