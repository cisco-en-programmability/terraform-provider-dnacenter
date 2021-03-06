---
page_title: "dna_pnp_device_claim_site Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_device_claim_site data source allows you to claim a Cisco DNA Center PnP device to a site.
---

# Data Source dna_pnp_device_claim_site

The dna_pnp_device_claim_site data source allows you to unclaim a Cisco DNA Center PnP device to a site.

## Example Usage

```hcl
data "dna_pnp_device_claim_site" "response" {
  provider       = dnacenter
  device_id = var.device_id
  site_id = var.site_id
  type = "Default"
}
```

## Argument Reference

- `device_id` - The device_id argument.
- `site_id` - The site_id argument.
- `type` - The type argument.

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

- `response` - The item's response.
- `version` - The item's version.
