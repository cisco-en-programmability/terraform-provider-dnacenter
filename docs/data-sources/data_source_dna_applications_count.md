---
page_title: "dna_applications_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The applications data source allows you to count the amount of DNACenter applications.
---

# Data Source dna_applications_count

The applications data source allows you to count the amount of DNACenter applications.

## Example Usage

```hcl
data "dna_applications_count" "amount" {
  provider = dnacenter
}
```

## Attributes Reference

The following attributes are exported.

- `response` - DNACenter applications count.
