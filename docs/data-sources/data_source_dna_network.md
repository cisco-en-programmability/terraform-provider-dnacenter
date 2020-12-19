---
page_title: "dna_network Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_network data source allows you to retrieve information about a particular DNACenter networks.
---

# Data Source dna_network

The dna_network data source allows you to retrieve information about a particular DNACenter networks.

## Example Usage

```hcl
data "dna_network" "response" {
  provider = dnacenter
  credential_sub_type = "CLI"
}
```

## Argument Reference

- `site_id` - (Optional) The site_id param.

## Attributes Reference

The following attributes are exported.

- `items` - Items in a DNACenter network. See [Items](#items) below for details.

### Items

Each item contains the following attributes.

- `group_uuid` - The item's group uuid.
- `inherited_group_name` - The item's inherited group name.
- `inherited_group_uuid` - The item's inherited group uuid.
- `instance_type` - The item's instance type.
- `instance_uuid` - The item's instance uuid.
- `key` - The item's key.
- `namespace` - The item's namespace.
- `type` - The item's type.
- `value` - The item's value.
- `version` - The item's version.
