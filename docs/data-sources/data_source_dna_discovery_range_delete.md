---
page_title: "dna_discovery_range_delete Data Source - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_discovery_range_delete data source allows you to delete DNACenter discovery by specified range.
---

# Data Source dna_discovery_range_delete

The dna_discovery_range_delete data source allows you to delete DNACenter discovery by specified range.

## Example Usage

```hcl
data "dna_discovery_range_delete" "response" {
  provider = dnacenter
  start_index = 1
  records_to_delete = 2
}
```

## Argument Reference

- `start_index` - (Required) The start_index param.
- `records_to_delete` - (Required) The records_to_delete param.
