# Contributing
We format commit messages according to the [conventional-commits](https://www.conventionalcommits.org/) specification.

## Creating a release
This is how maintainers of this software will create a new release:

1. Run `git tag <version>`
2. Run `git-chglog -o CHANGELOG.md`
3. Run `git push`
4. Run `git push --tags`
