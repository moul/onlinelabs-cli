<p align="center"><img width="50%" src="docs/static_files/cli-artwork.png" /></p>

<p align="center">
  <a href="https://circleci.com/gh/scaleway/scaleway-cli/tree/v2"><img src="https://circleci.com/gh/scaleway/scaleway-cli/tree/v2.svg?style=shield" alt="CircleCI" /></a>
  <a href="https://goreportcard.com/report/github.com/scaleway/scaleway-cli"><img src="https://goreportcard.com/badge/scaleway/scaleway-cli" alt="GoReportCard" /></a> <!-- GoReportCard do not support branches. -->
</p>

# Scaleway CLI (v2)

**:warning: This version is under active development, keep in mind that things can break. We advise you to not use this version in production.** 

Scaleway is a single way to create, deploy and scale your infrastructure in the cloud. We help thousands of businesses to run their infrastructures easily.

If you are looking for a stable version, [see the version 1](https://github.com/scaleway/scaleway-sdk-go).

# Installation

<!--- TODO:
## With a Package Manager (Recommended)

A package manager allows to install and upgrade the Scaleway CLI with a single command. We recommend this installation mode for more simplicity and reliability. We support a growing set of package managers to feat your preferences and your platform. Note that some package managers are maintained by our community:

### Homebrew

Install the latest stable release on macOS using [Homebrew](http://brew.sh): _Comming soon..._

```sh
brew install scw
```

### Chocolatey

Install the lastest stable release on Windows using [Chocolatey](https://chocolatey.org/): _Coming soon..._

```powershell
choco install scaleway-cli
```

### Others

TODO: Add other package managers:
- [Chocolate](https://chocolatey.org/packages/scaleway-cli/)
- [AUR](https://aur.archlinux.org/packages/scaleway-cli/)
- [Snap](https://snapcraft.io/)
- [Apt](https://wiki.debian.org/Apt)
-->

## Manually

### Compiled Release Binaries

We provide [static-compiled release binaries](https://github.com/scaleway/scaleway-cli/releases/latest) for darwin (macOS), GNU/Linux, and Windows platforms.

You just have to download the binary compatible with your platform to a directory available in your `PATH`:
```bash
# /usr/local/bin is in my PATH
echo $PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

# With wget
wget "https://github.com/scaleway/scaleway-cli/releases/download/v2.0.0-alpha.1/scw-darwin-amd64" -O /usr/local/bin/scw

# With curl
curl -o /usr/local/bin/scw -L https://github.com/scaleway/scaleway-cli/releases/download/v2.0.0-alpha.1/scw-darwin-amd64
```

For Windows, [this official guide](https://docs.microsoft.com/en-us/previous-versions/office/developer/sharepoint-2010/ee537574(v=office.14)) explains how to add tools to your `PATH`.

### Debian

First, download [the `.deb` file](https://github.com/scaleway/scaleway-cli/releases/latest) compatible with your architecture:

```bash
export ARCH=amd64 # Can be 'amd64', 'arm', 'arm64' or 'i386'
wget "https://github.com/scaleway/scaleway-cli/releases/download/v2.0.0-alpha.1/scw_2.0.0-alpha.1_${ARCH}.deb" -O /tmp/scw.deb
```

Then, run the installation and remove the `.deb` file:
```bash
dpkg -i /tmp/scw.deb && rm -f /tmp/scw.deb
```

<!-- TODO:
## With a Docker Image

For each release, we deliver a tagged image on the [Scaleway Docker Hub](https://hub.docker.com/r/scaleway/cli/tags) so can run `scw` in a sandboxed way: _Coming soon..._

```sh
docker run scaleway/cli version
```
-->

## Build Locally

If you have a >= Go 1.11 environment, you can install the `HEAD` version to test the latest features or to [contribute](CONTRIBUTING.md).
Note that this development version could include bugs, use [tagged releases](https://github.com/scaleway/scaleway-cli/releases/latest) if you need stability.

```bash
go get github.com/scaleway/scaleway-cli/cmd/scw
```

Dependancies: We only use go [Go Modules](https://github.com/golang/go/wiki/Modules) with vendoring.

# Getting Started

Just run the initialization command and let yourself be guided! :dancer:

```bash
scw init
```

It will set up your profile, the authentication, and the auto-completion.

# Examples

TODO: Add a list of examples here.

# Tutorials

TODO: Add a list of tutorials here.

# Development

This repository is at its early stage and is still in active development.
If you are looking for a way to contribute please read [CONTRIBUTING.md](CONTRIBUTING.md).

# Reach Us

We love feedback.
Feel free to reach us on [Scaleway Slack community](https://slack.scaleway.com/), we are waiting for you on [#opensource](https://scaleway-community.slack.com/app_redirect?channel=opensource).
