---
page_title: "dna_tag Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The tag data source allows you to retrieve information about a particular DNACenter tag.
---

# Data Source dna_tag

The tag data source allows you to retrieve information about a particular DNACenter order.

## Example Usage

```terraform
data "dna_tag" "found" {
  provider = dnacenter
  sort_by = "name"
  order = "des"
}
```

## Argument Reference

- `name` - (Optional) DNACenter tag name.
- `name_space` - (Optional) DNACenter tag nameSpace.
- `attributes` - (Optional) DNACenter tag attributes.
- `level` - (Optional) DNACenter tag level.
- `offset` - (Optional) DNACenter tag offset.
- `limit` - (Optional) DNACenter tag limit.
- `size` - (Optional) DNACenter tag size.
- `field` - (Optional) DNACenter tag field.
- `sort_by` - (Optional) DNACenter tag sortBy.
- `order` - (Optional) DNACenter tag order.
- `system_tag` - (Optional) DNACenter tag systemTag.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `items` - (Computed) Items in a DNACenter tag. See [Tag items](#tag-items) below for details.

### Tag items

Each tag item contains `system_tag`, `description`, `name`, `instance_tenant_id` and the `dynamic_rules` list.

- `system_tag` - (Computed) The tag's systemTag flag.
- `description` - (Computed) The tag's description.
- `name` - (Computed) The tag's name.
- `instance_tenant_id` - (Computed) The tag's instanceTenantId.
- `dynamic_rules` - (Computed) The tag's dynamic rules. See [Dynamic Rules](#dynamic-rules) below for details.

#### Dynamic Rules

- `member_type` - (Computed) The dynamic rule's member type.
- `rules` - (Computed) The dynamic rule's rules definition. See [Rules](#rules) below for details.

##### Rules

- `name` - (Computed) The rule's name.
- `operation` - (Computed) The rule's operation.
- `value` - (Computed) The rule's value.
- `values` - (Computed) The rule's values, a list of strings.
- `items` - (Computed) The rule's items, a list of strings.
