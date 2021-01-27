---
page_title: "dna_tag_member Data Source - terraform-provider-dnacenter"
subcategory: "Tag"
description: |-
  The dna_tag_member data source allows you to retrieve information about a particular Cisco DNA Center tag member.
---

# Data Source dna_tag_member

The dna_tag_member data source allows you to retrieve information about a particular Cisco DNA Center tag member.

## Example Usage

```hcl
data "dna_tag_member" "response" {
  provider    = dnacenter
  tag_id      = dna_tag.data.id
  member_type = "networkdevice"
}
```

## Argument Reference

- `tag_id` - (Required) The tag_id param.
- `member_type` - (Required) The member_type param. A value of dna_tag_member_type.
- `member_association_type` - (Optional) The member_association_type param.
- `level` - (Optional) The level param.
- `offset` - (Optional) The offset param.
- `limit` - (Optional) The limit param.

## Attributes Reference

The following attributes are exported.

- `items` - The items response. This is a JSON response.
