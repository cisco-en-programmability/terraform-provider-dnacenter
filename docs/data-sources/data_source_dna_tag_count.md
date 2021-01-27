---
page_title: "dna_tag_count Data Source - terraform-provider-dnacenter"
subcategory: "Tag"
description: |-
  The tag data source allows you to count the amount of Cisco DNA Center tags.
---

# Data Source dna_tag_count

The tag data source allows you to count the amount of Cisco DNA Center tags.

## Example Usage

```hcl
data "dna_tag_count" "amount" {
  provider = dnacenter
  name = "Tag012"
}
```

## Argument Reference

- `name` - (Optional) Cisco DNA Center tag name.
- `name_space` - (Optional) Cisco DNA Center tag nameSpace.
- `attribute_name` - (Optional) Cisco DNA Center tag attributeName.
- `level` - (Optional) Cisco DNA Center tag level.
- `size` - (Optional) Cisco DNA Center tag size.
- `system_tag` - (Optional) Cisco DNA Center tag systemTag.

## Attributes Reference

The following attributes are exported.

- `response` - Cisco DNA Center tags count.
