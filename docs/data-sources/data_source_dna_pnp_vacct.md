---
page_title: "dna_pnp_vaact Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_vaact data source allows you to list Cisco DNA Center Virtual Accounts.
---

# Data Source dna_pnp_vaact

The dna_pnp_vaact data source allows you to list Cisco DNA Center Virtual Accounts.

## Example Usage

```hcl
data "dna_pnp_vaact" "list" {
  provider = dnacenter
  domain = var.domain
}
```

## Argument Reference

- `domain` - (Optional) Smart Account Domain param.

## Attributes Reference

The following attributes are exported.

- `items` - Cisco DNA Center Virtual Accounts.
