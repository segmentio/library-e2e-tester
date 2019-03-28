# Releasing

Binaries are automatially published on tags from our CI using [`github-release`](https://github.com/aktau/github-release).

## git-extras

If you have [`git-extras`](https://github.com/tj/git-extras) installed, you can simply run `git release --semver major/minor/patch`. See the [documentation](https://github.com/tj/git-extras/blob/master/Commands.md#git-release) for details.

## manual

1. Tag a release with `git tag -a x.y.z -m "Version x.y.z"` where `x.y.z` is the version you are releasing.
2. Publish your tags with `git push --tags`.
