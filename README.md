# GoVers

GoVers provides commit message linting, automated semantic version bumps, and changelog generation for git-based code
repositories.

## Overview

GoVers implements the [Conventional Commit](https://conventionalcommits.org/) specification. By following the spec and
adding `feat` and `fix` commit types, GoVers can produce semantic version numbers based on commit messages and generate
changelogs for a given set of commits. GoVers can also enforce conventional commit messages, making it useful as a
pre-commit hook in your git workflow.

## Workflow

Borrowed heavily from [Release-Please](https://github.com/googleapis/release-please)

### Development Flow

Perform your standard workflow, but follow the conventional commits spec when writing commit messages. Specifically:

* `fix` denotes a bugfix and correlates to a SemVer patch release
* `feat` denoes a new feature and correlates to a SemVer minor release
* `fix!` or `feat!` indicates a breaking change and correlates to a SemVer major release
* Other valid commit types can be defined in conventional-commit.toml

### Release Flow

#### Beginning the release cycle

When you are ready to begin the release cycle, switch to the branch being released and run `govers release init`. This
will perform the following steps:

* Find the most recent canonical semver tag
  * If no non-prerelease tag is found, govers will use version 0.0.0 as the previous version
* Collect all commit messages since the last version
  * If `fix!` or `feat!` is present, bump to the next major version
  * If `feat` is present, bump to the next minor version
  * If `fix` is present, bump to the next patch version
  * Otherwise, print an error message and exit
* Create a new branch called `release/<version>`
  * If the release branch already exists, indicate a release is already in progress and exit
* Tag the new branch `<version>-RC.1`

#### Updating the release

Commits can be made against the release branch for feature and scope creep, and for fixes to issues identified in QA,
integration, acceptance or other forms of testing. When these commits are made, run `govers release update` to bump the
release candidate number and apply the tag.

#### Previewing the release

`govers release preview` will cause GoVers to print the updates which will be made to the CHANGELOG.md file for the 
current release branch.

#### Releasing code

To finalize the release, run `govers release`. GoVers will:
* Verify there is an active release branch
* Collect all commit messages on the branch and verify the version is still valid
  * If the commits performed on the branch would increment the version, GoVers will prompt the user to rename the release
  * Alternatively, passing the `-force` switch to `govers release` will skip the user prompt
* Update the CHANGELOG.md file by appending the calculated changeset to the beginning of the file and creating a new commit
* Apply the final version tag to the branch
* Merge the branch to `main`, or to the main branch specified in `govers.toml`

### Prereleases

A prerelease version can be created from any point in the commit history by running `govers pre [-name <value>]`. This
will calculate the semantic version of the current commit using the standard rules for a release. If no name is given,
GoVers will use `-PRE.1` as the first prerelease commit and increment the number each time a new prerelease is build
before a release is performed. If a name is specified, GoVers will attempt to use that value for the tag, and will
display an error if the prerelease tag already exists.

### Reading the current version

GoVers is intended to be called from CI pipelines in order to automaticaly apply the correct version into your build
artifacts. Running `govers` with no arguments will check the current commit for a version tag and display it if present.

GoVers can also be used to generate versions including build metadata. `govers snapshot` will print a version number
using the latest SemVer version plus a build identifier in the form of `<commitish>-YYMMDD.HHMMSS`
