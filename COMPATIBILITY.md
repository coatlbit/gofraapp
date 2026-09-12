# GoFraApp compatibility notes

GoFraApp provides Frappe-compatible DocType JSON, Desk assets, website assets,
modules, reports, pages, and hook metadata to the GoFra runtime. Install it with:

```bash
gofra-cli install-app gofraapp --site go.local
gofra-cli migrate --site go.local
```

The installer recursively discovers DocType JSON files, copies schemas into the
site registry, synchronizes their tables, and records the app in both
`site_config.json` and `apps.txt`. At runtime, GoFra serves the app's compiled
assets and maps its schemas to Desk metadata and REST resources.

Verification and the current compatibility boundary are documented in
`/Users/xxx/Dev/gofra/VERIFICATION_REPORT.md` and
`/Users/xxx/Dev/gofra/docs/bench-compatibility.md`.
