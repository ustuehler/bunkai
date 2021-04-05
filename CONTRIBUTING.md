# Contributing
We format commit messages according to the [conventional-commits](https://www.conventionalcommits.org/) specification.

## Creating a release
This is how maintainers of this software will create a new release:

1. Run the [`git-chglog` command](https://github.com/git-chglog/git-chglog), and
   review its output.

2. Pick the next [semantic version](https://semver.org/) based on changes listed
   in the **Unreleased** section to substitute for `$version`, below.

3. Update the change log, commit, create a tag, and push.

   ```bash
   git-chglog --next-tag $version -o CHANGELOG.md
   git commit -m "chore(release): $version" CHANGELOG.md
   git tag $version
   git push --follow-tags
   ```
