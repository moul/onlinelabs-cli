<p align="center"><img width="50%" src="docs/static_files/cli-artwork.png" /></p>

<p align="center">
  <a href="https://circleci.com/gh/scaleway/scaleway-cli/tree/v2"><img src="https://circleci.com/gh/scaleway/scaleway-cli/tree/v2.svg?style=shield" alt="CircleCI" /></a>
  <a href="https://goreportcard.com/report/github.com/scaleway/scaleway-cli"><img src="https://goreportcard.com/badge/scaleway/scaleway-cli" alt="GoReportCard" /></a> <!-- GoReportCard do not support branches. -->
</p>

# Scaleway CLI (v2)

Scaleway CLI is a tool to help you pilot your Scaleway infrastructure directly from your terminal.

# Installation

## With a Package Manager (Recommended)

A package manager installs and upgrades the Scaleway CLI with a single command.
We recommend this installation mode for more simplicity and reliability:

### Homebrew

Install the [latest stable release](https://formulae.brew.sh/formula/scw) on macOS/Linux using [Homebrew](http://brew.sh):

```sh
brew install scw
```

### Arch Linux

Install the latest stable release on Arch Linux from [official repositories](https://archlinux.org/packages/extra/x86_64/scaleway-cli/).
For instance with `pacman`:

```sh
pacman -S scaleway-cli
```

### Chocolatey

Install the latest stable release on Windows using [Chocolatey](https://chocolatey.org/) ([Package](https://chocolatey.org/packages/scaleway-cli)):

```powershell
choco install scaleway-cli
```

## Manually

### Released Binaries

We provide [static-compiled binaries](https://github.com/scaleway/scaleway-cli/releases/latest) for darwin (macOS), GNU/Linux, and Windows platforms.
You just have to download the binary compatible with your platform to a directory available in your `PATH`:

#### Linux

```bash
curl -s https://raw.githubusercontent.com/scaleway/scaleway-cli/master/scripts/get.sh | sh
```

#### Windows

You can download the last release here: <https://github.com/scaleway/scaleway-cli/releases><br/>
[This official guide](https://docs.microsoft.com/en-us/previous-versions/office/developer/sharepoint-2010/ee537574%28v%3Doffice.14%29) explains how to add tools to your `PATH`.

## Docker Image

You can use the CLI as you would run any Docker image:

```sh
docker run -i --rm scaleway/cli:latest
```

See more in-depth information about running the CLI in Docker [here](./docs/docker.md)

# Getting Started

## Setup your configuration

After you [installed](#Installation) the latest release just run the initialization command and let yourself be guided! :dancer:

```bash
scw init
```

It will set up your profile, the authentication, and the auto-completion.  
NB: you'll need to have an **API-key** (access-key + access-secret), so be sure to create one on the [scaleway web console](https://console.scaleway.com/iam/api-keys).

## Basic commands

```
# Create an instance server
scw instance server create type=DEV1-S image=ubuntu_noble zone=fr-par-1 tags.0="scw-cli"

# List your servers
scw instance server list

# Create a Kubernetes cluster named foo with cilium as CNI, in version 1.17.4 and with a pool named default composed of 3 DEV1-M and with 2 tags
scw k8s cluster create name=foo version=1.17.4 pools.0.size=3 pools.0.node-type=DEV1-M pools.0.name=default tags.0=tag1 tags.1=tag2
```

## Environment

You can configure your config or enable functionalities with environment variables.

Variables to override config are describe in [config documentation](docs/commands/config.md).
To enable beta features, you can set `SCW_ENABLE_BETA=1` in your environment.

# Reference documentation

| Namespace      | Description                             | Documentation                                                                                                     |
|----------------|-----------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| `account`      | User related data                       | [CLI](./docs/commands/account.md) / [API](https://www.scaleway.com/en/developers/api/account/project-api/)        |
| `applesilicon` | Apple silicon API                       | [CLI](./docs/commands/apple-silicon.md) / [API](https://www.scaleway.com/en/developers/api/apple-silicon/)        |
| `autocomplete` | Autocomplete related commands           | [CLI](./docs/commands/autocomplete.md)                                                                            |
| `baremetal`    | Baremetal API                           | [CLI](./docs/commands/baremetal.md) / [API](https://www.scaleway.com/en/developers/api/elastic-metal/)            |
| `billing`      | Billing API                             | [CLI](./docs/commands/billing.md) / [API](https://www.scaleway.com/en/developers/api/billing/)                    |
| `cockpit`      | Cockpit API                             | [CLI](./docs/commands/cockpit.md) / [API](https://www.scaleway.com/en/developers/api/cockpit/)                    |
| `config`       | Config file management                  | [CLI](./docs/commands/config.md)                                                                                  |
| `container`    | Serverless Container API                | [CLI](./docs/commands/container.md) / [API](https://www.scaleway.com/en/developers/api/serverless-containers/)    |
| `documentdb`   | DocumentDB API                          | [CLI](./docs/commands/document-db.md)                                                                             |
| `dns`          | DNS API                                 | [CLI](./docs/commands/dns.md) / [API](https://www.scaleway.com/en/developers/api/domains-and-dns/)                |
| `feedback`     | Send feedback to the Scaleway CLI Team! | [CLI](./docs/commands/feedback.md)                                                                                |
| `flexibleip`   | Flexible IP API                         | [CLI](./docs/commands/fip.md) / [API](https://www.scaleway.com/en/developers/api/elastic-metal-flexible-ip/)      |
| `function`     | Serverless Function API                 | [CLI](./docs/commands/function.md) / [API](https://www.scaleway.com/en/developers/api/serverless-functions/)      |
| `iam`          | IAM API                                 | [CLI](./docs/commands/iam.md) / [API](https://www.scaleway.com/en/developers/api/iam/)                            |
| `info`         | Get info about current settings         | [CLI](./docs/commands/info.md)                                                                                    |
| `init`         | Initialize the config                   | [CLI](./docs/commands/init.md)                                                                                    |
| `instance`     | Instance API                            | [CLI](./docs/commands/instance.md) / [API](https://www.scaleway.com/en/developers/api/instance/)                  |
| `iot`          | IoT API                                 | [CLI](./docs/commands/iot.md) / [API](https://www.scaleway.com/en/developers/api/iot/)                            |
| `ipam`         | IPAM API                                | [CLI](./docs/commands/ipam.md) / [API](https://www.scaleway.com/en/developers/api/ipam/)                          |
| `jobs`         | Serverless Jobs API                     | [CLI](./docs/commands/jobs.md) / [API](https://www.scaleway.com/en/developers/api/serverless-jobs/)               |
| `k8s`          | Kapsule API                             | [CLI](./docs/commands/k8s.md) / [API](https://www.scaleway.com/en/developers/api/kubernetes/)                     |
| `lb`           | Load Balancer API                       | [CLI](./docs/commands/lb.md) / [API](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/)         |
| `marketplace`  | Marketplace API                         | [CLI](./docs/commands/marketplace.md)                                                                             |
| `mnq`          | Messaging and Queueing API              | [CLI](./docs/commands/mnq.md) / [API](https://www.scaleway.com/en/developers/api/messaging-and-queuing/sqs-api/)  |
| `object`       | Object-storage utils                    | [CLI](./docs/commands/object.md) / [API](https://www.scaleway.com/en/docs/object-storage-feature/)                |
| `rdb`          | Database RDB API                        | [CLI](./docs/commands/rdb.md) / [API](https://www.scaleway.com/en/developers/api/managed-database-postgre-mysql/) |
| `redis`        | Redis API                               | [CLI](./docs/commands/redis.md) / [API](https://www.scaleway.com/en/developers/api/managed-database-redis// )     |
| `registry`     | Container registry API                  | [CLI](./docs/commands/registry.md) / [API](https://www.scaleway.com/en/developers/api/registry/)                  |
| `secret`       | Secret manager API                      | [CLI](./docs/commands/secret.md) / [API](https://www.scaleway.com/en/developers/api/secret-manager/)              |
| `shell`        | Start Shell mode                        | [CLI](./docs/commands/shell.md)                                                                                   |
| `tem`          | Transactional Email API                 | [CLI](./docs/commands/tem.md) / [API](https://www.scaleway.com/en/developers/api/transactional-email/)            |
| `vpc-gw`       | VPC Gateway API                         | [CLI](./docs/commands/vpc-gw.md) / [API](https://www.scaleway.com/en/developers/api/public-gateway/)              |
| `vpc`          | VPC API                                 | [CLI](./docs/commands/vpc.md) / [API](https://www.scaleway.com/en/developers/api/vpc/)                            |

## Build it yourself

### Build Locally

If you have a >= Go 1.13 environment, you can install the `HEAD` version to test the latest features or to [contribute](./.github/CONTRIBUTING.md).
Note that this development version could include bugs, use [tagged releases](https://github.com/scaleway/scaleway-cli/releases/latest) if you need stability.

```bash
go install github.com/scaleway/scaleway-cli/v2/cmd/scw@latest
```

Dependencies: We only use go [Modules](https://github.com/golang/go/wiki/Modules) with vendoring.

### Build with Docker

You can build the `scw` CLI with Docker. If you have Docker installed, you can run:

```sh
docker build -t scaleway/cli .
```

Once built, you can then use the CLI as you would run any image:

```sh
docker run -i --rm scaleway/cli
```

See more in-depth information about running the CLI in Docker [here](./docs/docker.md)

# Development

This repository is at its early stage and is still in active development.
If you are looking for a way to contribute please read [CONTRIBUTING.md](./.github/CONTRIBUTING.md).

# Reach Us

We love feedback.
Don't hesitate to open a [Github issue](https://github.com/scaleway/scaleway-cli/issues/new) or
feel free to reach us on [Scaleway Slack community](https://slack.scaleway.com/),
we are waiting for you on [#opensource](https://scaleway-community.slack.com/app_redirect?channel=opensource).
