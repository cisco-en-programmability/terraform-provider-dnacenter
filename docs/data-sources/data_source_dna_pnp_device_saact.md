---
page_title: "dna_pnp_device_saact Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_device_saact data source allows you to list DNACenter Smart Accounts.
---

# Data Source dna_pnp_device_saact

The dna_pnp_device_saact data source allows you to list DNACenter Smart Accounts.

## Example Usage

```hcl
data "dna_pnp_device_saact" "list" {
  provider = dnacenter
}
```

## Attributes Reference

The following attributes are exported.

- `items` - DNACenter Smart Accounts.
