---
page_title: "Provider: Cisco DNA Center"
subcategory: ""
description: |-
  Terraform provider for interacting with Cisco DNA Center SDK.
---

# Cisco DNA Center Provider

The Cisco DNA Center provider is used to interact with Cisco DNA Center APIs. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

The DNA Center SDK client could be generated with the following parameters:

```hcl
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

Do not keep your authentication password in HashiCorp for production environments, use environment variables.

## Argument Reference

- **username** - (Optional) - Username to authenticate to Cisco DNA Center API
- **password** - (Optional) - Password to authenticate to Cisco DNA Center API
- **base_url** - (Optional) - Cisco DNA Center address
- **debug** - (Optional) - debug flag for Cisco DNA Center (defaults to `false`)
- **ssl_verify** - (Optional) - ssl_verify flag for Cisco DNA Center (defaults to `true`)
