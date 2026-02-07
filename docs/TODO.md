# fiber-example - TODO

## TODO

* [x] general: add a Makefile to simplify commands ... wip
* [x] general: update to latest Fiber v3 (just released) ... wip
* [x] general: add another command-line (cli) application (with its main function), so by Go best practices, in a subfolder 'cmd/healthcheck/' ... wip
* [x] general: move handler/s in its own sub-module ('handler/' subfolder); maybe with a subfolder for 'api/' and other subfolders for other content ... wip
* [x] general: update handler for home page to return some static content: an html page with inside the list of published routes via template, a sample css file, a modern favicon ('favicon.svg' file), etc ... wip
* [x] general: add a common route prefix to all API exposed, like '/api' (so that a sample URI could be '/api/hello' or '/api/v1/hello') ... wip
* [x] general: add a 'Makefile' to simplify command lines (to clean, build, create distribution package, run tests, format sources, run the application in debug mode, etc); some tasks (using my standard names) are already included in the 'README.md', so align with it ... wip
* [x] general: update 'Dockerfile' to add an HEALTHCHECK (on a route published by the application for example), important because this is a web application (webapp) ... wip
* [x] general: check if add a full-static frontend (maybe with sources if a folder 'frontend/' and its binary distribution in another to check later) and related commands to simplify usage; note that it could even live completely in another repository and be copied in a distribution folder to use here ... wip
* [x] general: check if use Go/golang embedded file system features to compile into main executable even static assets like static content, a full-static frontend, etc ... wip
* [x] general: add CI jobs using GitHub Actions ... wip
* [x] general: add badges to README ... wip
* [x] general: check Go modules usage (enabled with "on" but by default disabled so ""), if it's done in the right way here (see related env vars in the not versioned '.envrc' file, maybe managed by the installed utility 'direnv') ... note that after updating env vars file, to update even in current shell, do: `direnv allow` ... wip
* [x] general: fix editing errors in Visual Studio Code (VSCode); note that this requires to move sources in the right folder/s and probably even a fix in Go/golang project file 'go.mod' (import statements both remote and local with rewrite or other similar) ... wip
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
* [x] general: add a source ('version.go' but check if in 'main' module or in a 'version' one) with a 'version' (or 'Version' if need to be public) struct and some variables inside (with default values like "unset") to store some build-related informations (like: branch if given, commit/short commit hash, tag if given, and even others like application version like "1.0.0-snapshot", etc) that usually are injected from outside (like a CI pipeline); update 'Makefile' to pass to build phase, both normal compilation and even in Docker containers; print those values at startup ... to simplify local tests/development, added some variable/s to the usual '.envrc' file (in a real env it could be set in the build environment by the CI pipeline or other build script/facility to generate/increase it automatically); update make file, ensure 'build-optimized' is good, then add even to 'build-normal', 'run', and maybe other tasks for DEV env (with non-optimized code) ... maybe later, it's too complex for this project initial stage and to be managed in the right way it requires a build (shell) script and a 'Makefile' to inject variables
* [x] general: print detailed version (and source control related informations) when executed from command-line (cli) with the flag '--version'/'-v'; all this would be very useful to fully identify the code/commit for the generated executable (otherwise with only release number it's not enough or could be outdated); for example look at `kubectl version` output, like: `{Client Version: version.Info{Major:"1", Minor:"22", GitVersion:"v1.22.2", GitCommit:"8b5a19147530eaac9476b0ab82980b4088bbc1b2", GitTreeState:"clean", BuildDate:"2021-09-15T21:38:50Z", GoVersion:"go1.16.8", Compiler:"gc", Platform:"linux/amd64"}` ... maybe later, it's too early here now
* [x] general: update dependencies and remove indirect (transitive) dependencies from 'go.mod', because they are re-added; remove old dependencies with `go mod tidy`, and repeat often ... ok, update requirements to latest Go/golang 1.18 and to latest Fiber (currently 'v2.34.0') and updated even all indirect dependencies to latest
* [x] general: move main sources from the 'src/' folder in project root (the 'src/' folder is not really useful in Go/Golang projects); then update (if any): 'README.md', 'Makefile', 'Dockerfile', build scripts, etc ... ok, removed folder 'src/' and 'test/' from the project and updated other files
* [x] general: update to latest Fiber 2.x (and Go/golang 1.25.x, even if not strictly required, but it will be required for Fiber v3) ... ok, changes done, ensured all works; then add a git tag and create a maintenance branch '2.x'


---------------
