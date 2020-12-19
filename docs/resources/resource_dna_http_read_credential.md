---
page_title: "dna_http_read_credential Resource - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_http_read_credential resource allows you to configure a DNACenter HTTP Read credential.
---

# Resource dna_http_read_credential

The dna_http_read_credential resource allows you to configure a DNACenter HTTP Read credential.

## Example Usage

```hcl
resource "dna_http_read_credential" "response" {
  provider = dnacenter
  item {
    username        = "usuario2"
    password        = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
    port            = 23
  }
}
```

## Argument Reference

- `item` - (Required) Item in a DNACenter HTTP Read credential. See [Credential item](#credential-item) below for details.

### Credential item

Each HTTP Read credential item contains the following arguments.

- `comments` - (Optional) The HTTP Read credential's comments.
- `credential_type` - (Optional) The HTTP Read credential's credential type. Available values are "GLOBAL" and "APP".
- `description` - (Optional) The HTTP Read credential's description.
- `id` - (Optional) The HTTP Read credential's id.
- `instance_tenant_id` - (Optional) The HTTP Read credential's instance tenant id.
- `instance_uuid` - (Optional) The HTTP Read credential's instance uuid.
- `password` - (Required) The HTTP Read credential's password.
- `port` - (Required) The HTTP Read credential's port.
- `secure` - (Optional) The HTTP Read credential's secure flag.
- `username` - (Required) The HTTP Read credential's username.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The HTTP Read credential's updated time with format RFC850.
