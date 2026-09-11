# GoFraApp Transition, Architecture & Integration Report

> **Generated:** 2026-09-11  
> **Status:** 100% Integrated, 361 Schemas Synced Cleanly, Pure Go Models & Hooks, Asset Bundles & Desk UI Complete

---

## 1. Executive Summary
`gofraapp` (`/Users/xxx/Dev/gofra-instance/apps/gofraapp`) is the foundational core application for the GoFra framework ecosystem. It mirrors the role of the `frappe` core app in Frappe Framework v16, entirely re-engineered for GoFra:
- **Zero Python Overhead**: All legacy Python runtime files and dynamic metaprogramming have been replaced with statically typed Go structs, compiled hooks, and JSON schema definitions.
- **Full Module Coverage**: 361 DocTypes across 11 modules are generated into type-safe Go models and registered with GoFra's engine.
- **Frontend & Desk Assets**: Prebuilt production bundles (`dist/`), SVG icon symbol caches, and assets manifest (`assets.json`) are self-contained in `public/`.
- **Desk UI Component Stack**: Includes Frappe UI Vue 3 component library (`ui/`), composables, and Socket.IO realtime integration (`realtime/`).

---

## 2. Module & Directory Structure
All modules are organized with their respective DocTypes, pages, reports, workspaces, and UI assets:

```
apps/gofraapp/
├── app.json                  # GoFra Application Manifest
├── hooks.go                  # Statically typed GoFra hooks & API router
├── modules.txt               # Module index (11 modules)
├── patches.txt               # Migration patch index
├── package.json              # Asset build configuration (esbuild, PostCSS, Vue)
├── core/                     # System & Security DocTypes (User, Role, DocType, DocField, DocPerm, etc.)
├── desk/                     # Desk UI Workspaces, Desktop pages, Notification logs
├── contacts/                 # Address and Contact management
├── email/                    # Email Account, Email Queue, and SMTP configuration
├── workflow/                 # Workflow states, transitions, actions
├── automation/               # Automated triggers and background actions
├── integrations/             # Webhooks, OAuth, Social login keys
├── custom/                   # Custom fields, DocPerms, and property setters
├── geo/                      # Countries, Currencies, and Territories
├── printing/                 # Print formats, Letter heads, Print styles
├── website/                  # Web pages, Web forms, Themes, Portals, Blogs
├── realtime/                 # Realtime Socket.IO handler integration
├── ui/                       # Frappe UI Vue 3 modern component library & composables
├── public/                   # Frappe v16 Desk SPA frontend assets, icons, and dist bundles
│   ├── assets.json           # Asset manifest mapping bundle entrypoints
│   ├── dist/                 # Precompiled CSS & JS bundles (desk, form, list, report, etc.)
│   ├── icons/                # Timeless SVG icon sets
│   ├── js/                   # Desk runtime scripts
│   └── scss/                 # SASS stylesheets
├── templates/                # Web templates
└── www/                      # Public portal routes
```

---

## 3. Ported DocTypes & Generated Go Models
361 DocTypes across 11 modules have been indexed, validated, and generated into type-safe Go structs:

| Module | Key DocTypes Ported | Output Models |
|---|---|---|
| **Core** | `User`, `Role`, `HasRole`, `DocType`, `DocField`, `DocPerm`, `SystemSettings`, `ModuleDef`, `File`, `ActivityLog`, `AccessLog`, `Comment`, `Communication` | `core/doctype/<name>/<name>.go` |
| **Desk** | `Workspace`, `NotificationLog`, `Event`, `Dashboard`, `DashboardChart`, `NumberCard`, `RouteHistory`, `SidebarItemGroupLink` | `desk/doctype/<name>/<name>.go` |
| **Contacts** | `Address`, `Contact`, `ContactEmail`, `ContactPhone`, `Gender`, `Salutation` | `contacts/doctype/<name>/<name>.go` |
| **Email** | `EmailAccount`, `EmailQueue`, `EmailQueueRecipient`, `EmailTemplate`, `EmailDomain` | `email/doctype/<name>/<name>.go` |
| **Workflow** | `Workflow`, `WorkflowState`, `WorkflowAction`, `WorkflowTransition`, `WorkflowDocumentState` | `workflow/doctype/<name>/<name>.go` |
| **Automation** | `AutoRepeat`, `Milestone`, `MilestoneTracker`, `AssignmentRule` | `automation/doctype/<name>/<name>.go` |
| **Integrations** | `Webhook`, `OAuthClient`, `SocialLoginKey`, `TokenCache`, `PushNotificationSettings` | `integrations/doctype/<name>/<name>.go` |
| **Custom** | `CustomField`, `PropertySetter`, `CustomDocPerm`, `CustomizeDocType` | `custom/doctype/<name>/<name>.go` |
| **Geo** | `Country`, `Currency`, `Territory` | `geo/doctype/<name>/<name>.go` |
| **Printing** | `PrintFormat`, `LetterHead`, `PrintSettings`, `PrintStyle`, `PrintHeading` | `printing/doctype/<name>/<name>.go` |
| **Website** | `WebPage`, `WebForm`, `WebsiteSettings`, `PortalSettings`, `WebsiteTheme`, `WebsiteSidebar`, `PersonalDataDeletionRequest` | `website/doctype/<name>/<name>.go` |

---

## 4. GoFra Engine Integration & Runtime Features

1. **Statically Typed Hooks (`hooks.go`)**:
   - Registered via `core.RegisterApp(&core.AppDefinition{ ... })`.
   - Full lifecycle hooks: `BeforeInstall`, `AfterInstall`, `BeforeMigrate`, `AfterMigrate`.
   - Type-safe DocEvents dispatcher supporting wildcards `*` and specific DocTypes (`User`, `Activity Log`).
   - Custom API routes registered at `/api/v1/method/gofraapp/*` (`ping`, `info`, `status`).

2. **Asset Resolution & Static Serving**:
   - `gofra/core/desk.go` natively resolves assets from `./apps/gofraapp/public`, `./public/frappe`, and `sites/assets/`.
   - Dynamic asset manifest loading from `apps/gofraapp/public/assets.json`.
   - Full Desk SPA serving under `/desk`, `/app`, `/desk/*`, and `/app/*`.

3. **Database Migration & Schema Sync**:
   - Dialect quoting handles multi-word and case-sensitive DocType table names across SQLite, PostgreSQL, MySQL, and BadgerDB.
   - Non-column layout elements (`Section Break`, `Column Break`, `Heading`) are filtered during DDL generation.
   - Fully synced all 361 schemas into `site.db` on `go.local`.

4. **Frappe RPC Compatibility Layer**:
   - Core endpoints (`frappe.desk.*`, `frappe.client.*`, `frappe.auth.*`, `frappe.handler.*`) seamlessly interact with `gofraapp` schemas and permissions.

