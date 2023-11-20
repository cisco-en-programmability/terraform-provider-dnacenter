
# terraform-provider-dnacenter

terraform-provider-dnacenter is a Terraform Provider for [Cisco DNA Center](https://developer.cisco.com/docs/dna-center/)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x
- [Go](https://golang.org/doc/install) 1.20 (to build the provider plugin)

## Introduction

The terraform-provider-dnacenter provides a Terraform provider for managing and automating your Cisco DNA Center environment. It consists of a set of resources and data-sources for performing tasks related to DNA Center.

This collection has been tested and supports Cisco DNA Center 2.3.5.3.

Other versions of this collection have support for previous Cisco DNA Center versions. The recommended versions are listed below on the [Compatibility matrix](#compatibility-matrix).

## Compatibility matrix
The following table shows the supported versions.

| Cisco DNA Center version | Terraform "dnacenter" provider version |
|--------------------------|----------------------------------------|
| 2.1.1                    | 0.0.4                                  |
| 2.2.3.3                  | 0.3.0-beta                             |
| 2.2.3.3                  | 0.3.0                                  |
| 2.3.3.0                  | 1.0.19-beta                            |
| 2.3.5.3                  | 1.1.27-beta                            |

If your SDK, Terraform provider is older please consider updating it first.

## Using the provider

There are two ways to get and use the provider.
1. Downloading & installing it from registry.terraform.io
2. Building it from source

### From build (For test)

Clone this repository to: `$GOPATH/src/github.com/cisco-en-programmability/terraform-provider-dnacenter`

```sh
$ mkdir -p $GOPATH/src/github.com/cisco-en-programmability/
$ cd $GOPATH/src/github.com/cisco-en-programmability/
$ git clone https://github.com/cisco-en-programmability/terraform-provider-dnacenter.git
```

Enter the provider directory and build the provider

! **NOTE**:
It is important to check the architecture of your operating system in the file [MakeFile](./Makefile)

```sh
$ cd $GOPATH/src/github.com/cisco-en-programmability/terraform-provider-dnacenter
$ make developtest
```

If the Makefile values (HOSTNAME, NAMESPACE, NAME, VERSION) were not changed, then the following code could used without changes.
Otherwise change the values accordingly.


To use this provider, copy and paste this code into your Terraform configuration. Then, run terraform init.

```hcl
terraform {
  required_providers {
    dnacenter = {
      source = "cisco-en-programmability/dnacenter"
      version = "1.1.27-beta"
    }
  }
}

# Configure provider with your Cisco DNA Center SDK credentials
provider "dnacenter" {
  # Cisco DNA Center user name
  username = "admin"
  # it can be set using the environment variable DNAC_BASE_URL

  # Cisco DNA Center password
  password = "admin123"
  # it can be set using the environment variable DNAC_USERNAME

  # Cisco DNA Center base URL, FQDN or IP
  base_url = "https://172.168.196.2"
  # it can be set using the environment variable DNAC_PASSWORD

  # Boolean to enable debugging
  debug = "false"
  # it can be set using the environment variable DNAC_DEBUG

  # Boolean to enable or disable SSL certificate verification
  ssl_verify = "false"
  # it can be set using the environment variable DNAC_SSL_VERIFY
}
```

There are several examples of the use of the provider within the folder [samples](./examples/samples)

## Example for_each
```hcl
locals {
 interfaces = {
   "1" = { description = "desc1", interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809" },
   "2" = { description = "desc2", interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486801" },
   "3" = { description = "desc3", interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486802" }
 }
}

resource "dnacenter_interface_update" "example" {
  provider = dnacenter
  for_each = local.interfaces
  parameters {
    description    = each.value.description
    interface_uuid = each.value.interface_uuid
    vlan_id        = each.key
  }
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

In the [docs directory](./docs/) you can find the documentation source for this 1.0.19-beta version.

You can find the documentation online for the previously released versions at [Terraform Registry - Cisco DNA Center provider](https://registry.terraform.io/providers/cisco-en-programmability/dnacenter/latest/docs).

# Contributing

Ongoing development efforts and contributions to this provider are tracked as issues in this repository.

We welcome community contributions to this project. If you find problems, need an enhancement or need a new data-source or resource, please open an issue or create a PR against the [Terraform Provider for Cisco DNA Center repository](https://github.com/cisco-en-programmability/terraform-provider-dnacenter/issues).

# Change log

All notable changes to this project will be documented in the [CHANGELOG](./CHANGELOG.md) file.

The development team may make additional changes as the library evolves with the Cisco DNA Center.

**NOTE**: Consider reviewing the Changelog to review the new features of the 1.0.19-beta version.

## License

This library is distributed under the license found in the [LICENSE](./LICENSE) file.
