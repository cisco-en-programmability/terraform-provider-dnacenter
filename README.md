# DEPRECATION NOTICE
Cisco is consolidating two Catalyst Center Terraform providers. This provider has been deprecated and will no longer be updated. The new provider ([CiscoDevNet/terraform-provider-catalystcenter](https://github.com/CiscoDevNet/terraform-provider-catalystcenter)) is now the official, actively maintained version.

## Why is this change happening? 
The new Terraform provider will provide more efficient operations, ongoing support, new features, and improvements, ensuring a more robust and future-proof experience for all Catalyst Center users. 

## Can I continue using the old provider? 
Yes, but the old provider will no longer receive updates or new features. We encourage you to plan a migration to the new provider.
Bugs will be fixed on a best effort basis until June 30th, 2026.

## What are the recommended migration steps? 
1. Review Documentation: Read the new provider’s documentation carefully. Resource names and attributes are not necesarily the same.
2. Update Configurations: Refactor your .tf files to match the new HCL format and attribute names. 
3. Test Plans: Run terraform plan to identify required changes and resolve any errors. 
4. Handle State: Use terraform import or state manipulation commands to align existing resources with the new provider’s resource model. 
5. Validate Changes: Test in a non-production environment before applying changes in production. 

## What happens if I don’t migrate? 
Your existing setup will continue to function, but no new features will be available for this provider. Over time, compatibility issues may arise as Terraform and Meraki evolve.
Bugs will be fixed on a best effort basis until June 30th, 2026.

- - - 

# terraform-provider-dnacenter

terraform-provider-dnacenter is a Terraform Provider for [Cisco Catalyst Center](https://developer.cisco.com/docs/dna-center/)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x
- [Go](https://golang.org/doc/install) 1.20 (to build the provider plugin)

## Introduction

The terraform-provider-dnacenter provides a Terraform provider for managing and automating your Cisco Catalyst Center environment. It consists of a set of resources and data-sources for performing tasks related to Catalyst Center.

This collection has been tested and supports Cisco Catalyst Center 2.3.5.3.

Other versions of this collection have support for previous Cisco Catalyst Center versions. The recommended versions are listed below on the [Compatibility matrix](#compatibility-matrix).

## Compatibility matrix
The following table shows the supported versions.

| Cisco Catalyst Center version | Terraform "dnacenter" provider version |
|--------------------------|----------------------------------------|
| 2.1.1                    | 0.0.4                                  |
| 2.2.3.3                  | 0.3.0-beta                             |
| 2.2.3.3                  | 0.3.0                                  |
| 2.3.3.0                  | 1.0.19-beta                            |
| 2.3.5.3                  | 1.1.33-beta                            |
| 2.3.7.6                  | 1.2.0-beta                             |
| 2.3.7.9                  | 1.4.0-beta                             |

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
      version = "1.1.33-beta"
    }
  }
}

# Configure provider with your Cisco Catalyst Center SDK credentials
provider "dnacenter" {
  # Cisco Catalyst Center user name
  username = "admin"
  # it can be set using the environment variable DNAC_BASE_URL

  # Cisco Catalyst Center password
  password = "admin123"
  # it can be set using the environment variable DNAC_USERNAME

  # Cisco Catalyst Center base URL, FQDN or IP
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

You can find the documentation online for the previously released versions at [Terraform Registry - Cisco Catalyst Center provider](https://registry.terraform.io/providers/cisco-en-programmability/dnacenter/latest/docs).

# Contributing

Ongoing development efforts and contributions to this provider are tracked as issues in this repository.

We welcome community contributions to this project. If you find problems, need an enhancement or need a new data-source or resource, please open an issue or create a PR against the [Terraform Provider for Cisco Catalyst Center repository](https://github.com/cisco-en-programmability/terraform-provider-dnacenter/issues).

# Change log

All notable changes to this project will be documented in the [CHANGELOG](./CHANGELOG.md) file.

The development team may make additional changes as the library evolves with the Cisco Catalyst Center.

**NOTE**: Consider reviewing the Changelog to review the new features of the 1.4.0-beta version.

## License

This library is distributed under the license found in the [LICENSE](./LICENSE) file.
