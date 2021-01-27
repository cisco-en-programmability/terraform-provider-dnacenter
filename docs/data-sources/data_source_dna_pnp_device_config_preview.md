---
page_title: "dna_pnp_device_config_preview Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_device_config_preview data source allows you to preview a Cisco DNA Center PnP device config.
---

# Data Source dna_pnp_device_config_preview

The dna_pnp_device_config_preview data source allows you to preview a Cisco DNA Center PnP device config.

## Example Usage

```hcl
data "dna_pnp_device_config_preview" "result" {
  provider = dnacenter
  device_id = var.device_id
  site_id = var.site_id
  type = "Default"
}
```

## Argument Reference

- `device_id` - (Optional) The device id argument.
- `site_id` - (Optional) The site id argument.
- `type` - (Optional) The type argument.

## Attributes Reference

The following attributes are exported.

- `item` - The config_preview's item. See [Item](#item) below for details.

### Items

- `response` - The item's response. See [Config preview response](#config-preview-response) below for details.
- `version` - The item's response.

#### Config preview response

- `complete` - The config preview's complete flag.
- `config` - The config preview's config.
- `error` - The config preview's error flag.
- `error_message` - The config preview's error message.
- `expired_time` - The config preview's expired time.
- `rf_profile` - The config preview's rf profile.
- `sensor_profile` - The config preview's sensor profile.
- `site_id` - The config preview's site id.
- `start_time` - The config preview's start time.
- `task_id` - The config preview's task id.
