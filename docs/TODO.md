# fiber-example - TODO

## TODO

* [x] general: update handler for home page to return some static content: an html page with inside the list of published routes via template, a sample css file, a modern favicon ('favicon.svg' file), etc ... wip
* [x] general: add a common route prefix to all API exposed, like '/api' (so that a sample URI could be '/api/hello' or '/api/v1/hello') ... wip
* [x] general: add a 'Makefile' to simplify command lines (to clean, build, create distribution package, run tests, format sources, run the application in debug mode, etc); some tasks (using my standard names) are already included in the 'README.md', so align with it ... wip
* [x] general: add CI jobs using GitHub Actions ... wip
* [x] general: add badges to README ... wip
* [x] general: check Go modules usage (enabled with "on" but by default disabled so ""), if it's done in the right way here (see related env vars in the not versioned '.envrc' file, maybe managed by the installed utility 'direnv') ... note that after updating env vars file, to update even in current shell, do: `direnv allow` ... wip
* [x] general: check if remove indirect (transitive) dependencies from 'go.mod' ... wip
* [x] general: fix editing errors in Visual Studio Code (VSCode) ... wip
* [x] general: perform a simple benchmark and write results here in a txt, just to have some number ... wip
* [x] general: add a 'Dockerfile' (and '.dockerignore', etc), or one for dev and one for run; then put related commands in dedicated make tasks ... wip
* [x] general: add cross-compilation for Windows/Mac/Linux in related (but secondary) make tasks ... wip


---------------


## DONE

* [x] general: create/update initial skeleton (from an existing example) ... ok, follow these articles: [Go Fiber by Examples - How can the Fiber Web Framework be useful ? - DEV Community](https://dev.to/koddr/go-fiber-by-examples-how-can-the-fiber-web-framework-be-useful-487a), [here](https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j), [miguelmota/golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers), [Fiber](https://gofiber.io/), etc
* [x] general: ensure initial version works ... ok, run it in dev/debug mode with: `go run src/main.go` and browse to [localhost:8000](http://127.0.0.1:8000/); to create a dev/debug build with: `go build -o build/fiber-example src/main.go`, finally remove with: `rm ./build/fiber-example*`
* [x] general: create and publish to the remote repository ... ok, but for now in a common 'note' repo, with some docs, scripts, etc
* [x] general: add initial version of 'README.md' with some summary info ... ok, and update/improve later
* [x] general: add features in 'CHANGELOG.md' ... maybe later
* [x] general: add some test sources ... maybe later
* [x] general: add a command-line (cli) flag '-help' to show some help ... maybe later
* [x] general: add a public repo at GitHub (and write here and in the README) and then publish sources there ... ok, created public repo [smartiniOnGitHub/fiber-example](https://github.com/smartiniOnGitHub/fiber-example.git) and published all there; note that (since some time) new repos at GitHub has default branch 'main' (no more 'master'), so I renamed even the local one to 'main'


---------------
