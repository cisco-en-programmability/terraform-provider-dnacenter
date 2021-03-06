---
page_title: "dna_tag Data Source - terraform-provider-dnacenter"
subcategory: "Tag"
description: |-
  The tag data source allows you to retrieve information about a particular Cisco DNA Center tag.
---

# Data Source dna_tag

The tag data source allows you to retrieve information about a particular Cisco DNA Center order.

## Example Usage

```hcl
data "dna_tag" "found" {
  provider = dnacenter
  sort_by = "name"
  order = "des"
}
```

## Argument Reference

- `name` - (Optional) Cisco DNA Center tag name.
- `name_space` - (Optional) Cisco DNA Center tag nameSpace.
- `attributes` - (Optional) Cisco DNA Center tag attributes.
- `level` - (Optional) Cisco DNA Center tag level.
- `offset` - (Optional) Cisco DNA Center tag offset.
- `limit` - (Optional) Cisco DNA Center tag limit.
- `size` - (Optional) Cisco DNA Center tag size.
- `field` - (Optional) Cisco DNA Center tag field.
- `sort_by` - (Optional) Cisco DNA Center tag sortBy.
- `order` - (Optional) Cisco DNA Center tag order.
- `system_tag` - (Optional) Cisco DNA Center tag systemTag.

## Attributes Reference

The following attributes are exported.

- `items` - Items in a Cisco DNA Center tag. See [Tag items](#tag-items) below for details.

### Tag items

Each tag item contains `system_tag`, `description`, `name`, `instance_tenant_id` and the `dynamic_rules` list.

- `system_tag` - The tag's systemTag flag.
- `description` - The tag's description.
- `name` - The tag's name.
- `instance_tenant_id` - The tag's instanceTenantId.
- `dynamic_rules` - The tag's dynamic rules. See [Dynamic Rules](#dynamic-rules) below for details.

#### Dynamic Rules

- `member_type` - The dynamic rule's member type.
- `rules` - The dynamic rule's rules definition. See [Rules](#rules) below for details.

##### Rules

- `name` - The rule's name.
- `operation` - The rule's operation.
- `value` - The rule's value.
- `values` - The rule's values, a list of strings.
- `items` - The rule's items, a list of strings.
