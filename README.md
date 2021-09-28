# Terraform Provider for Cisco DNA Center

- Website: https://www.terraform.io

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x
- [Go](https://golang.org/doc/install) 1.17 (to build the provider plugin)

## Introduction

The terraform-provider-dnacenter provides a Terraform provider for managing and automating your Cisco DNA Center environment. It consists of a set of resources and data-sources for performing tasks related to DNA Center.

This collection has been tested and supports Cisco DNA Center 2.1.1.

## Using the provider

If you are building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory, run `terraform init` to initialize it.

ex.

```hcl
  # Configure provider with your Cisco DNA Center SDK credentials
  provider "dnacenter" {
    # Cisco DNA Center user name
    # username = "admin"
    # it can be set using the environment variable DNAC_BASE_URL

    # Cisco DNA Center password
    # password = "admin123"
    # it can be set using the environment variable DNAC_USERNAME

    # Cisco DNA Center base URL, FQDN or IP
    # base_url = "https://172.168.196.2"
    # it can be set using the environment variable DNAC_PASSWORD

    # Boolean to enable debugging
    # debug = "false"
    # it can be set using the environment variable DNAC_DEBUG

    # Boolean to enable or disable SSL certificate verification
    # ssl_verify = "false"
    # it can be set using the environment variable DNAC_SSL_VERIFY
  }

  # Configure CLI credential
  resource "dna_cli_credential" "response" {
    provider = dnacenter
    item {
      username = "${var.username}"
      password = "${var.password}"
      credential_type = "APP"
    }
  }
  output "dna_cli_credential_response" {
    value = dna_cli_credential.response
  }

  # Retrieve project's templates
  data "dna_template_project" "response" {
    provider = dnacenter
    name     = "Cloud DayN Templates"
  }

  data "dna_template" "response" {
    provider   = dnacenter
    project_id = data.dna_template_project.response.items.0.id
  }
  output "dna_template_response" {
    value = data.dna_template.response
  }
```

In the examples directory you can find more.

## Building The Provider

Clone this repository to: `$GOPATH/src/github.com/cisco-en-programmability/terraform-provider-dnacenter`

```sh
$ mkdir -p $GOPATH/src/github.com/cisco-en-programmability
$ cd $GOPATH/src/github.com/cisco-en-programmability
$ git clone https://github.com/cisco-en-programmability/terraform-provider-dnacenter.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/cisco-en-programmability/terraform-provider-dnacenter
$ make build
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed
on your machine (version 1.15+ is _required_). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary
in the `$GOPATH/bin` directory.

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

We welcome community contributions to this project. If you find problems, need an enhancement or need a new data-source or resource, please open an issue or create a PR against the [Terraform Provider for Cisco DNA Center repository](https://github.com/cisco-en-programmability/terraform-provider-dnacenter/issues).

## License

This library is distributed under the license found in the [LICENSE](./LICENSE) file.
