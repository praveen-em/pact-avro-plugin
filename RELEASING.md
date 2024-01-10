# Releasing

## Tag and push a new release

Update `VERSION` in [`Makefile`](./Makefile)

Tag the version and push:

```
git tag -a v0.0.1 -m "Initial release"
git push origin v0.0.1
```

Goreleaser will pick up the new tag, and release it to Github.

## Adding to pkg.go.dev
[pkg.go.dev](https://pkg.go.dev) is the central source of information about Go packages and modules. By pushing a new release to a public github repo accessible through a public URL, technically the module is already published. There is no need to upload or publish it anywhere else.

Run the `go list` [command](https://pkg.go.dev/cmd/go#hdr-List_packages_or_modules) to prompt Go to update its index of modules with information about the module. Precede the command with a statement to set the `GOPROXY` environment variable to a Go proxy. 

`GOPROXY=proxy.golang.org go list -m example.com/mymodule@v0.1.0`

This will ensure that your request reaches the proxy, and the module appears on pkg.go.dev site eventually. If you don't see your module or package, you can add it by simply doing any of the steps specified in [pkgsite](https://pkg.go.dev/about#adding-a-package).

## Commit messages

Pact uses the [Conventional Changelog](https://github.com/bcoe/conventional-changelog-standard/blob/master/convention.md)
commit message conventions. Please ensure you follow the guidelines, as they
help us automate our release process.

If you'd like to get some CLI assistance, getting setup is easy:

```shell
npm install commitizen -g
npm i -g cz-conventional-changelog
```

`git cz` to commit and commitizen will guide you.