---
page_title: "Provider: DNACenter"
subcategory: ""
description: |-
  Terraform provider for interacting with DNACenter SDK.
---

# DNACenter Provider

The DNACenter provider is used to interact with Cisco DNA Center APIs. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

The DNA Center SDK client could be generated with the following parameters:

```terraform
provider "dnacenter" {
  username = "admin"
  password = "admin123"
  base_url = "https://172.168.196.2"
  debug = "false"
  ssl_verify = "false"
}
```

Do not keep your authentication password in HashiCorp for production environments, use Terraform environment variables.

```shell
export DNAC_BASE_URL=https://172.168.196.2
export DNAC_USERNAME=admin
export DNAC_PASSWORD=admin123
export DNAC_DEBUG=false
export DNAC_SSL_VERIFY=false
```

```terraform
provider "dnacenter" {
}
```

## Schema

### Optional

- **username** (String, Optional) Username to authenticate to DNACenter API
- **password** (String, Optional) Password to authenticate to DNACenter API
- **base_url** (String, Optional) DNACenter address
- **debug** (String, Optional) debug flag for DNACenter (defaults to `false`)
- **ssl_verify** (String, Optional) ssl_verify flag for DNACenter (defaults to `true`)
