
# terraform-provider-dnacenter

terraform-provider-dnacenter is a Terraform Provider for [Cisco DNA Center]()

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x
- [Go](https://golang.org/doc/install) 1.15 (to build the provider plugin)

## Introduction

The terraform-provider-dnacenter provides a Terraform provider for managing and automating your Cisco DNA Center environment. It consists of a set of resources and data-sources for performing tasks related to DNA Center.

This collection has been tested and supports Cisco DNA Center 2.2.3.3.

## Using the provider

There are two ways to get and use the provider.
1. Downloading & installing it from registry.terraform.io
2. Building it from source

### From registry

To install this provider, copy and paste this code into your Terraform configuration. Then, run terraform init. 

```hcl
terraform {
  required_providers {
    dnacenter = {
      source = "cisco-en-programmability/dnacenter"
      version = "2.4.0"
    }
  }
}

provider "dnacenter" {
  # Configuration options
  # More info at https://registry.terraform.io/providers/cisco-en-programmability/dnacenter/latest/docs#example-usage
}
```

### From build

Clone this repository to: `$GOPATH/src/github.com/cisco-en-programmability/terraform-provider-dnacenter`

```sh
$ mkdir -p $GOPATH/src/github.com/cisco-en-programmability/
$ cd $GOPATH/src/github.com/cisco-en-programmability/
$ git clone https://github.com/cisco-en-programmability/{config.names.terraform}}.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/cisco-en-programmability/terraform-provider-dnacenter
$ make build
```

If the Makefile values (HOSTNAME, NAMESPACE, NAME, VERSION) were not changed, then the following code could used without changes.
Otherwise change the values accordingly.

To use this provider, copy and paste this code into your Terraform configuration. Then, run terraform init.

```hcl
terraform {
  required_providers {
    dnacenter = {
      source = "hashicorp.com/edu/dnacenter"
      version = "2.4.0"
    }
  }
}

provider "dnacenter" {
  # Configuration options
  # More info at https://registry.terraform.io/providers/cisco-en-programmability/dnacenter/latest/docs#example-usage
}
```


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed
on your machine (version 1.15+ is _required_). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-dnacenter
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources.

```sh
$ make testacc
```

## Documentation

In the docs directory you can find the documentation.

# Contributing

Ongoing development efforts and contributions to this provider are tracked as issues in this repository.

We welcome community contributions to this project. If you find problems, need an enhancement or need a new data-source or resource, please open an issue or create a PR against the [Terraform Provider for Cisco DNA Center repository](https://github.com/cisco-en-programmability/{config.names.terraform}}/issues).

# Change log

All notable changes to this project will be documented in the [CHANGELOG](./CHANGELOG.md) file.

The development team may make additional changes as the library evolves with the Cisco DNA Center.

## License

This library is distributed under the license found in the [LICENSE](./LICENSE) file.