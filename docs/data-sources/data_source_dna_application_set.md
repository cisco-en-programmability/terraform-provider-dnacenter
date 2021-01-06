---
page_title: "dna_application_set Data Source - terraform-provider-dnacenter"
subcategory: "Application Policy"
description: |-
  The dna_application_set data source allows you to retrieve information about a particular DNACenter application set.
---

# Data Source dna_application_set

The dna_application_set data source allows you to retrieve information about a particular DNACenter application set.

## Example Usage

```hcl
data "dna_application_set" "list" {
  provider = dnacenter
  offset = 0
  limit = 4
}
```

## Argument Reference

- `name` - (Optional) DNACenter application set's name.
- `offset` - (Optional) DNACenter application set's offset.
- `limit` - (Optional) DNACenter application set's limit.

## Attributes Reference

The following attributes are exported.

- `items` - Items in a DNACenter app set. See [Application Set items](#application-set-items) below for details.

### Application Set items

Each application set item contains the following attributes.

- `name` - The application set's name.
- `id` - The application set's id.
- `identity_source` - The application set's identity source. See [Identity Source](#identity-source) below for details.

### Identity Source

Each application set's identity source contains `id` and `type`.

- `id` - The identity source's id.
- `type` - The identity source's type.
