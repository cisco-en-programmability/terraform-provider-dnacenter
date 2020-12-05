---
page_title: "dna_tag_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The tag data source allows you to count the amount of DNACenter tags.
---

# Data Source dna_tag_count

The tag data source allows you to count the amount of DNACenter tags.

## Example Usage

```terraform
data "dna_tag_count" "amount" {
  provider = dnacenter
  name = "Tag012"
}
```

## Argument Reference

- `name` - (Optional) DNACenter tag name.
- `name_space` - (Optional) DNACenter tag nameSpace.
- `attribute_name` - (Optional) DNACenter tag attributeName.
- `level` - (Optional) DNACenter tag level.
- `size` - (Optional) DNACenter tag size.
- `system_tag` - (Optional) DNACenter tag systemTag.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `response` - (Computed) DNACenter tags count.
