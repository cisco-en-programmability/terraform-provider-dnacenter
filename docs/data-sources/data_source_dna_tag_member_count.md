---
page_title: "dna_tag_member_count Data Source - terraform-provider-dnacenter"
subcategory: "Tag"
description: |-
  The tag member data source allows you to count the amount of Cisco DNA Center tag's members.
---

# Data Source dna_tag_member_count

The tag member data source allows you to count the amount of Cisco DNA Center tag's members.

## Example Usage

```hcl
data "dna_tag_member_count" "amount" {
  provider = dnacenter
  id = ""
  member_type = "networkdevice"
}
```

## Argument Reference

- `id` - (Required) Cisco DNA Center tag member id.
- `member_type` - (Optional) Cisco DNA Center tag member type.
- `member_association_type` - (Optional) Cisco DNA Center tag member association type.
- `level` - (Optional) Cisco DNA Center tag member level.

## Attributes Reference

The following attributes are exported.

- `response` - Cisco DNA Center tag's member count.
