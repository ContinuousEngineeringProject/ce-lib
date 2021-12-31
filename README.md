# Continuous Engineering Library

[![License](https://img.shields.io/github/license/ContinuousEngineeringProject/ce-lib)](https://github.com/ContinuousEngineeringProject/ce-lib/blob/master/LICENSE)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
[![DepShield Badge](https://depshield.sonatype.org/badges/ContinuousEngineeringProject/ce-lib/depshield.svg)](https://depshield.github.io)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ContinuousEngineeringProject_ce-lib&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ContinuousEngineeringProject_ce-lib)
[![GoReport](https://goreportcard.com/badge/github.com/ContinuousEngineeringProject/ce-lib)](https://goreportcard.com/report/github.com/ContinuousEngineeringProject/ce-lib)
[![GoDoc](https://godoc.org/github.com/ContinuousEngineeringProject/ce-lib?status.svg)](https://godoc.org/github.com/ContinuousEngineeringProject/ce-lib)
[![GitHub release](https://img.shields.io/github/v/release/ContinuousEngineeringProject/ce-lib?include_prereleases)](https://github.com/ContinuousEngineeringProject/ce-lib/releases/latest)
[![Slack Status](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://continuousengproject.slack.com)

The Continuous Engineering Project is an open source framework that enables continuous engineering best practices through plug & play toolsets.

The Continuous Engineering Library, ce-lib, is a library used by the Continuous Engineering Factory.

## Installation

go-github is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/ContinuousEngineeringProject/ce-lib
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/ContinuousEngineeringProject/ce-lib"
```

and run `go get` without parameters.

Finally, to use the top-of-trunk version of this repo, use the following command:

```bash
go get github.com/ContinuousEngineeringProject/ce-lib@master
```


## Usage

```go
import "github.com/ContinuousEngineeringProject/ce-lib"
```


## Contributing

Read [CONTRIBUTING.md][CONTRIB] for details of all the ways you can contribute to the project.

Also read [CODE_OF_CONDUCT.md][COC] for details on our code of conduct for the project.


## Versioning

We use [SemVer][SEMVER] for versioning. For the versions available, see the [tags on this repository][REPOTAGS].


## License

Licensed under the MIT license - see the [LICENSE][LICENSE] file for details.


[LICENSE]: LICENSE
[SEMVER]: http://semver.org/
[COC]: CODE_OF_CONDUCT.md
[CONTRIB]: CONTRIBUTING.md
[REPOTAGS]: https://github.com/continuousengineeringproject/ce-lib/tags