# AGENTS.md

Go client for the Crowdin API v2 and Crowdin Enterprise API v2 (`github.com/crowdin/crowdin-api-client-go`).

Go 1.22. The library has zero non-stdlib runtime dependencies (testify is test-only) — keep it that way.

## Layout

Two packages, both under `crowdin/`:

- `crowdin/` — `package crowdin`: `crowdin.go` (the `Client`, `NewClient`, HTTP verbs, `ToPtr`) plus one `<resource>.go` service file per API resource
- `crowdin/model/` — `package model`: request/response structs, list options, `Validate()` methods; shared `errors.go`, `pagination.go`, `update.go`
- Tests sit next to the code (`<resource>_test.go`); fixtures are inline raw-string JSON — there is no testdata directory

## Commands

- Build: `go build -v ./...`
- Test (all): `go test -v ./...`; one test: `go test -run 'TestBranchesService_List' ./crowdin/`
- Lint (what CI runs): `golangci-lint run` (CI pins v1.57.2; config in `.golangci.yml`)

## Adding or changing an endpoint

Fetch the endpoint spec first (see Crowdin API reference below). Then:

1. Models in `crowdin/model/<resource>.go`: the entity struct with `json` tags; wrappers `XResponse{Data *X}` and `XListResponse{Data []*XResponse}`; list options embed `ListOptions` and implement `Values() (url.Values, bool)`; request structs implement `Validate() error` beginning with `if r == nil { return ErrNilRequest }` — `Post`/`Put`/`Patch` call it automatically, and a request type without `Validate()` silently skips validation. Optional request booleans are `*bool` with `omitempty` (so an explicit `false` transmits); optional request strings/ints are plain values with `omitempty`; nullable response fields are pointers. Initialisms stay upper-case: `ID`, `URL`, `IDs`.
2. Service methods in `crowdin/<resource>.go` on `type <X>Service struct { client *Client }`: the first parameter is always `ctx context.Context`; return `(payload, *Response, error)`, or `(*Response, error)` for deletes; the godoc comment ends with the developer.crowdin.com operation URL (enterprise URL for Enterprise-only endpoints) and a period.
3. For a new service, register it with two edits in `crowdin/crowdin.go`: a field on `Client` (alphabetical, gofmt-aligned) and an init line in `NewClient` under `// Initialize services.`.
4. Tests: `client, mux, teardown := setupClient(); defer teardown()`, then `mux.HandleFunc(path, ...)` using the shared helpers from `crowdin/crowdin_test.go` (`testMethod`, `testURL`, `testJSONBody`) and an inline JSON reply; assert with testify (`require.NoError`, `assert.Equal`) — the current house style; name tests `Test<X>Service_<Method>`. Model option/validation tests are table-driven in `crowdin/model/<resource>_test.go` and use the package-local lowercase `toPtr` (importing `crowdin.ToPtr` there is an import cycle).

A complete new service touches 4–5 files: the service file, its model file, `crowdin/crowdin.go`, and the test files.

## Lint rules that shape the code

`.golangci.yml` enables 46 linters at their defaults, on tests too, and the tree contains zero `//nolint` — write code that satisfies the linters rather than suppressing them:

- `godot`: every declaration comment ends with a period, struct-field comments included.
- `prealloc`: list methods preallocate — `list := make([]*model.X, 0, len(res.Data))` before the append loop.
- `exhaustive`: a switch over the typed string enums covers every constant — the house idiom is `case A, B: // valid` then `default: return errors.New(...)`.
- `gochecknoglobals`: no new package-level vars (`ErrNilRequest` passes only via its `Err` prefix).
- `gci`/`goimports`: import blocks are the stdlib group, a blank line, then this module's packages.

## Crowdin API reference

Before implementing or changing any endpoint, fetch its spec from the llms.txt indexes (pick by environment, then project type):

- https://support.crowdin.com/_llms-txt/api/crowdin/file-based.txt — Crowdin API, file-based projects (start here)
- https://support.crowdin.com/_llms-txt/api/crowdin/string-based.txt — Crowdin API, string-based projects
- https://support.crowdin.com/_llms-txt/api/enterprise/file-based.txt — Crowdin Enterprise API, file-based projects
- https://support.crowdin.com/_llms-txt/api/enterprise/string-based.txt — Crowdin Enterprise API, string-based projects

Each index links one spec file per route (e.g. `.../api.projects.strings.get.txt`) with the exact request and response shapes.

## Conventions

- Conventional Commits for commit messages and PR titles; CI lints PR titles.
- PRs target `main`.
- Keep the public API backward compatible.
- Never touch the `userAgent` constant in `crowdin/crowdin.go` — the library version lives there and release automation bumps it.

## PR checklist

A change is ready when:

1. `go build -v ./...` compiles,
2. `golangci-lint run` is clean,
3. `go test -v ./...` passes, and
4. every new or changed endpoint method has service tests (plus model tests for options and validation) and a godoc comment ending in its operation link.
