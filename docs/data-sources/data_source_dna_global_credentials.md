---
page_title: "dna_global_credentials Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_global_credentials data source allows you to retrieve information about a particular DNACenter global credentials.
---

# Data Source dna_global_credentials

The dna_global_credentials data source allows you to retrieve information about a particular DNACenter global credentials.

## Example Usage

```hcl
data "dna_global_credentials" "response" {
  provider = dnacenter
  credential_sub_type = "CLI"
}
```

## Argument Reference

- `credential_sub_type` - (Optional) The credential_sub_type param. Available values are "CLI", "SNMPV2_READ_COMMUNITY", "SNMPV2_WRITE_COMMUNITY", "SNMPV3", "HTTP_WRITE", "HTTP_READ" and "NETCONF".
- `sort_by` - (Optional) The sort_by param.
- `order` - (Optional) The order param.

## Attributes Reference

The following attributes are exported.

- `items` - Items in a DNACenter global credential. See [Items](#items) below for details.

### Items

Each item contains the following attributes.

- `comments` - The item's comments.
- `credential_type` - The item's credential_type.
- `description` - The item's description.
- `id` - The item's id.
- `instance_tenant_id` - The item's instance_tenant_id.
- `instance_uuid` - The item's instance_uuid.
