---
page_title: "dna_pnp_device_unclaim Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_device_unclaim data source allows you to unclaim a Cisco DNA Center PnP device.
---

# Data Source dna_pnp_device_unclaim

The dna_pnp_device_unclaim data source allows you to unclaim a Cisco DNA Center PnP device.

## Example Usage

```hcl
data "dna_pnp_device_unclaim" "response" {
  provider       = dnacenter
  device_id_list = []
}
```

## Argument Reference

- `device_id_list` - (Optional) The device id list.

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

- `json_array_response` - The item's json array response.
- `json_response` - The item's json response.
- `message` - The item's message.
- `status_code` - The item's status code.
