---
page_title: "dna_service_provider_profile Data Source - terraform-provider-dnacenter"
subcategory: "Network Settings"
description: |-
  The dna_service_provider_profile data source allows you to retrieve information about a particular DNACenter service provider profiles.
---

# Data Source dna_service_provider_profile

The dna_service_provider_profile data source allows you to retrieve information about a particular DNACenter service provider profiles.

## Example Usage

```hcl
data "dna_service_provider_profile" "response" {
  provider = dnacenter
}
```

## Attributes Reference

The following attributes are exported.

- `items` - Items in a DNACenter service provider profile. See [Items](#items) below for details.

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
- `value` - The item's value. See [value](#value) below for details.
- `version` - The item's version.

#### value

- `sla_profile_name` - The value of sla_profile_name.
- `sp_profile_name` - The value of sp_profile_name.
- `wan_provider` - The value of wan_provider.
