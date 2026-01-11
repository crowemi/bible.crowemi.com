# bible.crowemi.com

A small Go web app that serves curated scripture passages and summaries using Go's html/template with embedded templates.

## 🚀 Quickstart

Prerequisites:
- Go 1.20+ (or your preferred recent Go toolchain)
- (Optional) Docker

Clone and run locally:

```bash
git clone https://github.com/crowemi/bible.crowemi.com.git
cd bible.crowemi.com
go mod tidy
go run ./cmd/bible
# open http://localhost:8080
```

Build binary:

```bash
go build -o bin/bible ./cmd/bible
./bin/bible
```

Docker (build + run):

```bash
docker build -t bible .
docker run -p 8080:8080 bible
```

## ⚙️ Configuration

The app loads configuration with `config.LoadConfig(path)` and expects a JSON config file (e.g. `.secret/config.json`) passed at startup. See `cmd/bible/main.go` for the current usage.

## 📁 Templates

Templates are embedded via `templates.TemplateFS` (Go `embed.FS`). The config code parses HTML files and stores them in `config.Templates` for rendering.

Important notes:
- Templates are identified by `{{define "name"}}`, not by filename. If multiple files define the same name (e.g. `content`) the **last parsed** definition overrides earlier ones.
- The `base` layout uses `{{ block "content" . }} ... {{ end }}` — each page should provide its own content via a `block` or define a unique template name and call `{{template "base" .}}` inside that page's `define` block.

Common templates:
- `home` — main landing page
- `page` — generic page
- `base` — layout

## 💡 Passing data / links

- You cannot send an HTTP body with a plain `<a href="...">` link. Use query strings, path parameters, or forms/fetch POSTs.
- Example (query string): `<a href="/read?passage=Genesis%201-2">Read</a>` and in Go: `r.URL.Query().Get("passage")`.
- Example (form): submit a hidden form or use `fetch()` to POST JSON to the server.

## 🧰 Static assets / Tailwind

This project can use Tailwind via CDN in `templates/layout/base.html` or you can include a minimized local stylesheet (e.g. `static/css/tailwind-lite.css`) if you prefer not to use the CDN.

## ✅ Testing & Development tips

- Adding logging in `config.LoadConfig` helps inspect which templates were parsed and in what order.
- If a page shows the wrong content, check for duplicate `{{define "content"}}` templates and prefer scoping content with `{{define "home"}} {{block "content" .}}...{{end}} {{template "base" .}}`.

## Contributing

Open a PR with concise changes. Add tests where applicable and keep commits focused.

## License

This project is available under the MIT License. See `LICENSE` for details.

---

If you'd like, I can:
- Add the `README` to the repo (done),
- Add a sample `.secret/config.json.example`, or
- Add a short Developer Guide section detailing template conventions and common gotchas.

Tell me which you'd like next.
