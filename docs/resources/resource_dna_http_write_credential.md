---
page_title: "dna_http_write_credential Resource - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_http_write_credential resource allows you to configure a DNACenter HTTP Write credential.
---

# Resource dna_http_write_credential

The dna_http_write_credential resource allows you to configure a DNACenter HTTP Write credential.

## Example Usage

```hcl
resource "dna_http_write_credential" "response" {
  provider = dnacenter
  item {
    username        = "usuario2"
    password        = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
    port            = 24
    id              = "5abaa9c9-4470-46c4-90d6-107594164845"
  }
}
```

## Argument Reference

- `item` - (Required) Item in a DNACenter HTTP Write credential. See [Credential item](#credential-item) below for details.

### Credential item

Each HTTP Write credential item contains the following arguments.

- `comments` - (Optional) The HTTP Write credential's comments.
- `credential_type` - (Optional) The HTTP Write credential's credential type. Available values are "GLOBAL" and "APP".
- `description` - (Optional) The HTTP Write credential's description.
- `id` - (Optional) The HTTP Write credential's id.
- `instance_tenant_id` - (Optional) The HTTP Write credential's instance tenant id.
- `instance_uuid` - (Optional) The HTTP Write credential's instance uuid.
- `password` - (Required) The HTTP Write credential's password.
- `port` - (Required) The HTTP Write credential's port.
- `secure` - (Optional) The HTTP Write credential's secure flag.
- `username` - (Required) The HTTP Write credential's username.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The HTTP Write credential's updated time with format RFC850.
