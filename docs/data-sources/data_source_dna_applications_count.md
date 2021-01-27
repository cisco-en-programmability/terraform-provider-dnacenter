---
page_title: "dna_applications_count Data Source - terraform-provider-dnacenter"
subcategory: "Application Policy"
description: |-
  The applications data source allows you to count the amount of Cisco DNA Center applications.
---

# Data Source dna_applications_count

The applications data source allows you to count the amount of Cisco DNA Center applications.

## Example Usage

```hcl
data "dna_applications_count" "amount" {
  provider = dnacenter
}
```

## Attributes Reference

The following attributes are exported.

- `response` - Cisco DNA Center applications count.
