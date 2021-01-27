---
page_title: "dna_discovery_all_delete Data Source - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_discovery_all_delete data source allows you to delete all Cisco DNA Center discovery.
---

# Data Source dna_discovery_all_delete

The dna_discovery_all_delete data source allows you to delete all Cisco DNA Center discovery.

## Example Usage

```hcl
data "dna_discovery_all_delete" "response" {
  provider = dnacenter
}
```
