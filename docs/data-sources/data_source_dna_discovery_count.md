---
page_title: "dna_discovery_count Data Source - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_discovery_count data source allows you to count the amount of DNACenter discoveries discovered.
---

# Data Source dna_discovery_count

The dna_discovery_count data source allows you to count the amount of DNACenter discoveries discovered.

## Example Usage

```hcl
data "dna_discovery_count" "amount" {
  provider = dnacenter
  id = var.discovery_id
}
```

## Argument Reference

- `id` - (Required) The id param.
- `task_id` - (Optional) The task id param.

## Attributes Reference

The following attributes are exported.

- `response` - DNACenter discoveries discovered count.
