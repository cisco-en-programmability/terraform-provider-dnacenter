---
page_title: "dna_pnp_device_config_preview Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_pnp_device_config_preview data source allows you to preview a DNACenter PnP device config.
---

# Data Source dna_pnp_device_config_preview

The dna_pnp_device_config_preview data source allows you to preview a DNACenter PnP device config.

## Example Usage

```hcl
data "dna_pnp_device_config_preview" "result" {
  provider = dnacenter
  device_id = "<device_id>"
  site_id = "<site_id>"
  type = "Default"
}
```

## Argument Reference

- `device_reset_list` - (Optional) The device reset list argument. See [device_reset_list](#device_reset_list) below for details.
- `project_id` - (Optional) The project id argument.
- `workflow_id` - (Optional) The workflow id argument.

### device_reset_list

- `config_list` - (Optional) The device reset's config list. See [config_list](#config_list) below for details.
- `device_id` - (Optional) The device reset's device id.
- `license_level` - (Optional) The device reset's license level.
- `license_type` - (Optional) The device reset's license type.
- `top_of_stack_serial_number` - (Optional) The device reset's top of stack serial number.

#### config_list

- `config_parameters` - (Optional) The config's parameters. See [config_parameters](#config_parameters) below for details.

##### config_parameters

- `key` - (Optional) The config param's key.
- `value` - (Optional) The config param's value.

## Attributes Reference

The following attributes are exported.

- `item` - The config_preview's item. See [Item](#item) below for details.

### Item

- `json_array_response` - The item's json array response.
- `json_response` - The item's json response.
- `message` - The item's message.
- `status_code` - The item's status code.
