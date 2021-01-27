---
page_title: "dna_tag_member_type Data Source - terraform-provider-dnacenter"
subcategory: "Tag"
description: |-
  The tag member type data source allows you to list Cisco DNA Center tag member types.
---

# Data Source dna_tag_member_type

The tag member type data source allows you to list Cisco DNA Center tag member types.

## Example Usage

```hcl
data "dna_tag_member_type" "list" {
  provider = dnacenter
}
```

## Attributes Reference

The following attributes are exported.

- `response` - Cisco DNA Center tag member types.
