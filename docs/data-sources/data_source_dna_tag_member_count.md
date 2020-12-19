---
page_title: "dna_tag_member_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The tag member data source allows you to count the amount of DNACenter tag's members.
---

# Data Source dna_tag_member_count

The tag member data source allows you to count the amount of DNACenter tag's members.

## Example Usage

```hcl
data "dna_tag_member_count" "amount" {
  provider = dnacenter
  id = ""
  member_type = "networkdevice"
}
```

## Argument Reference

- `id` - (Required) DNACenter tag member id.
- `member_type` - (Optional) DNACenter tag member type.
- `member_association_type` - (Optional) DNACenter tag member association type.
- `level` - (Optional) DNACenter tag member level.

## Attributes Reference

The following attributes are exported.

- `response` - DNACenter tag's member count.
