---
page_title: "dna_cli_credential Resource - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_cli_credential resource allows you to configure a DNACenter CLI credential.
---

# Resource dna_cli_credential

The dna_cli_credential resource allows you to configure a DNACenter CLI credential.

## Example Usage

```hcl
resource "dna_cli_credential" "response" {
  provider = dnacenter
  item {
    username        = "usuario2"
    password        = "123456"
    credential_type = "APP"
    id              = "34091dbf-8d55-48b8-aec4-5a572c265370"
  }
}
```

## Argument Reference

- `item` - (Required) Item in a DNACenter CLI credential. See [CLI Credential item](#cli-credential-item) below for details.

### CLI Credential item

Each CLI credential item contains `comments`, `credential_type`, `description`, `enable_password`, `id`, `instance_tenant_id`, `instance_uuid`, `password` and `username`.

- `comments` - (Required) The CLI credential's comments.
- `credential_type` - (Required) The CLI credential's credential type. Available values are "GLOBAL" and "APP".
- `description` - (Required) The CLI credential's description.
- `enable_password` - (Optional) The CLI credential's enable password.
- `id` - (Optional) The CLI credential's id.
- `instance_tenant_id` - (Optional) The CLI credential's instance tenant id.
- `instance_uuid` - (Optional) The CLI credential's instance uuid.
- `password` - (Required) The CLI credential's password.
- `username` - (Required) The CLI credential's username.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The CLI credential's updated time with format RFC850.
- `id` - The CLI credential's id.
- `instance_tenant_id` - The CLI credential's instance tenant id.
- `instance_uuid` - The CLI credential's instance uuid.
