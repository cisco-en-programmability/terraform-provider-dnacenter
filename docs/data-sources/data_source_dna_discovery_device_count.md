---
page_title: "dna_discovery_device_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_discovery_device_count data source allows you to count the amount of DNACenter discovery devices jobs.
---

# Data Source dna_discovery_device_count

The dna_discovery_device_count data source allows you to count the amount of DNACenter discovery devices jobs.

## Example Usage

```hcl
data "dna_discovery_device_count" "amount" {
  provider = dnacenter
}
```

## Attributes Reference

The following attributes are exported.

- `response` - DNACenter discovery devices jobs count.
