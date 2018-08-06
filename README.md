# letsgo
An all in one tool suite for go

# Status
Work In Progress

# Features

Not yet decided completely， tell me if you have some ideas.

* go installation
* packages management
* find awesome libraries
* go doc in console
* jump to go package directory in fly


# Design

## Installation

```sh
letsgo init # install all dependent go tools
letsgo install {version} # install specific go version
letsgo list # list all avaliable go versions
letsgo to {version} # switch to certain go version
```

## Get packages

```sh
letsgo get <url> # go get
```

## Dependencies management

```sh
letsgo dep ...  # wrap for dep
```

## Find awesome libraries

```sh
letsgo awesome <keywords>
```

## Doc in console

```sh
letsgo doc <pkg> # render go doc in the console
```

## Jump in go src

```sh
letsgo cd <path> # fuzzy search path that matchs <path> and cd into it
```
