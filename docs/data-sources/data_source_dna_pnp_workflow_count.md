---
page_title: "dna_pnp_workflow_count Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_workflow_count data source allows you to count the amount of DNACenter workflows.
---

# Data Source dna_pnp_workflow_count

The dna_pnp_workflow_count data source allows you to count the amount of DNACenter workflows.

## Example Usage

```hcl
data "dna_pnp_workflow_count" "amount" {
  provider = dnacenter
  name     = ["Workflow 2", "Workflow 1"]
}
```

## Argument Reference

- `name` - (Required) The name params.

## Attributes Reference

The following attributes are exported.

- `response` - DNACenter workflows count.
