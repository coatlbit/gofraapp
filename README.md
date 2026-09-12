# GoFraApp (`gofraapp`)

<div align="center">
  <img src="public/images/frappe-framework-logo.svg" width="80" height="80" alt="GoFraApp Logo"/>
  <h3>Foundational Core Application for GoFra Framework</h3>
  <p><strong>Pure Go, High-Performance, Statically Typed Metadata Engine & Desk SPA</strong></p>
</div>

<div align="center">

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go)](https://golang.org)
[![Author](https://img.shields.io/badge/Author-Coatlbit-orange.svg)](https://coatlbit.org)
[![Website](https://img.shields.io/badge/Website-coatlbit.org-green.svg)](https://coatlbit.org)

</div>

---

## 1. Overview

**GoFraApp (`gofraapp`)** is the foundational core application for the **GoFra Framework** ecosystem. It serves as the direct Go-native counterpart to the `frappe` core application in Frappe Framework v16, engineered to eliminate Python runtime bottlenecks, Redis process overhead, and dynamic duck-typing vulnerabilities.

GoFraApp provides:
- **361 System & Domain DocTypes** across 11 modules with strongly typed Go models.
- **Complete Desk SPA Frontend**: Precompiled production bundles (`dist/`), SVG icon symbol caches, and dynamic asset manifests.
- **Frappe UI Component System**: Integrated Vue 3 component library (`ui/`) and composables.
- **Statically Typed Hooks (`hooks.go`)**: Compile-time lifecycle events, DocEvents, and custom API method routing.
- **Realtime Pub/Sub**: Native Go WebSockets and Socket.IO handler integration (`realtime/`).

---

## 2. Module Hierarchy & Ported DocTypes

GoFraApp organizes the entire system into 11 specialized modules:

| Module | Scope & Capabilities | Key Ported DocTypes |
|---|---|---|
| **Core** | Authentication, Users, Roles, Permissions, Security, Files | `User`, `Role`, `HasRole`, `DocType`, `DocField`, `DocPerm`, `File`, `ActivityLog`, `SystemSettings` |
| **Desk** | Workspace layouts, Dashboards, Notification logs, Shortcuts | `Workspace`, `NotificationLog`, `Event`, `Dashboard`, `DashboardChart`, `NumberCard`, `RouteHistory` |
| **Contacts** | Enterprise address book, Phone/Email indexing, Salutations | `Address`, `Contact`, `ContactEmail`, `ContactPhone`, `Gender`, `Salutation` |
| **Email** | SMTP/IMAP configurations, Email Queue, Templates, Domains | `EmailAccount`, `EmailQueue`, `EmailQueueRecipient`, `EmailTemplate`, `EmailDomain` |
| **Workflow** | Document approval states, transitions, actions, and tasks | `Workflow`, `WorkflowState`, `WorkflowAction`, `WorkflowTransition`, `WorkflowDocumentState` |
| **Automation** | Scheduled actions, milestones, automated assignment rules | `AutoRepeat`, `Milestone`, `MilestoneTracker`, `AssignmentRule` |
| **Integrations** | Webhooks, OAuth2 SSO, Social Login Providers, Push Notifications | `Webhook`, `OAuthClient`, `SocialLoginKey`, `TokenCache`, `PushNotificationSettings` |
| **Custom** | Runtime custom fields, property setters, custom permissions | `CustomField`, `PropertySetter`, `CustomDocPerm`, `CustomizeDocType` |
| **Geo** | Standard geographical entities, currencies, territorial trees | `Country`, `Currency`, `Territory` |
| **Printing** | Letterheads, print formats, print styles, and page layouts | `PrintFormat`, `LetterHead`, `PrintSettings`, `PrintStyle`, `PrintHeading` |
| **Website** | Web portal CMS, Web forms, Themes, Blogs, Privacy portals | `WebPage`, `WebForm`, `WebsiteSettings`, `PortalSettings`, `WebsiteTheme`, `WebsiteSidebar` |

---

## 3. Directory Layout

```
apps/gofraapp/
├── app.json                  # GoFra Application Manifest
├── hooks.go                  # Statically typed GoFra hooks & API router
├── modules.txt               # Module registry index (11 modules)
├── patches.txt               # Migration patch index
├── package.json              # Asset build toolchain (esbuild, PostCSS, Vue)
├── core/                     # Core system & security DocTypes
├── desk/                     # Desk UI Workspaces & Desktop pages
├── contacts/                 # Address and Contact management
├── email/                    # Email Account, Queue & SMTP
├── workflow/                 # Workflow states & transitions
├── automation/               # Automated triggers & background actions
├── integrations/             # Webhooks, OAuth & Social Login
├── custom/                   # Custom fields & Property setters
├── geo/                      # Countries & Currencies
├── printing/                 # Print formats & Letterheads
├── website/                  # Web pages, Forms, Portal & Themes
├── realtime/                 # Socket.IO & WebSocket event handlers
├── ui/                       # Frappe UI Vue 3 component library
├── public/                   # Frontend Desk assets, icons, and dist bundles
│   ├── assets.json           # Dynamic asset manifest
│   ├── dist/                 # Prebuilt CSS & JS production bundles
│   ├── icons/                # SVG icon sets
│   ├── js/                   # Desk runtime scripts
│   └── scss/                 # SASS stylesheets
├── templates/                # Web templates
└── www/                      # Public portal routes
```

---

## 4. Statically Typed Hooks (`hooks.go`)

GoFra replaces Python's string-based dictionaries with compiled type safety:

```go
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
			log.Println("[gofraapp] Initializing GoFra core setup...")
			return nil
		},
		AfterInstall: func() error {
			log.Println("[gofraapp] Successfully initialized GoFra core.")
			return nil
		},
		BeforeMigrate: func() error {
			log.Println("[gofraapp] Preparing schema migrations...")
			return nil
		},
		AfterMigrate: func() error {
			log.Println("[gofraapp] Core schema migrations completed.")
			return nil
		},

		DocEvents: map[string]map[string]core.DocHookFunc{
			"*": {
				"after_save": func(doc map[string]interface{}) error {
					return nil
				},
			},
		},

		RegisterRoutes: func(r fiber.Router) {
			group := r.Group("/api/v1/method/gofraapp")
			group.Get("/ping", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{"status": "ok", "app": "gofraapp", "version": "0.16.0"})
			})
			group.Get("/info", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{"app": "gofraapp", "author": "Coatlbit", "status": "active"})
			})
		},
	})
}
```

---

## 5. Quickstart & Installation

To install `gofraapp` into a GoFra site workspace:

```bash
# 1. Create a new site
gofra-cli new-site site1.local --db-type sqlite --admin-password admin

# 2. Install gofraapp
gofra-cli install-app gofraapp --site site1.local

# 3. Synchronize all 361 schemas
gofra-cli migrate --site site1.local

# 4. Set as active site and launch server
gofra-cli use site1.local
gofra-cli serve --port 8080
```

Open `http://localhost:8080/desk` to log in and use GoFra Desk.

---

## 6. Author & Licensing

- **Author**: **Coatlbit**
- **Website**: [coatlbit.org](https://coatlbit.org)
- **Contact Email**: [hola@coatlbit.org](mailto:hola@coatlbit.org)
- **License**: [MIT License](LICENSE) &copy; 2026 Coatlbit.
