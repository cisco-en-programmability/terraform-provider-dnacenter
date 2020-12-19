---
page_title: "dna_pnp_device_claim Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_pnp_device_claim data source allows you to claim a DNACenter PnP device.
---

# Data Source dna_pnp_device_claim

The dna_pnp_device_claim data source allows you to claim a DNACenter PnP device.

## Example Usage

```hcl
data "dna_pnp_device_claim" "response" {
  provider   = dnacenter
  config_id = pnp_device.config_id
  image_id = pnp_device.image_id
  project_id = pnp_device.project_id
  workflow_id = pnp_device.workflow_id
}
```

## Argument Reference

- `config_file_url` - (Optional) The config file url.
- `config_id` - (Optional) The config id.
- `device_claim_list` - (Optional) The device claim list. See [device_claim_list](#device_claim_list) below for details.
- `file_service_id` - (Optional) The file service id.
- `image_id` - (Optional) The image id.
- `image_url` - (Optional) The image url.
- `populate_inventory` - (Optional) The populate inventory.
- `project_id` - (Optional) The project id.
- `workflow_id` - (Optional) The workflow id.

### device_claim_list

- `config_list` - (Optional) The device claim's config list. See [config_list](#config_list) below for details.
- `device_id` - (Optional) The device claim's device id.
- `license_level` - (Optional) The device claim's license level.
- `license_type` - (Optional) The device claim's license type.
- `top_of_stack_serial_number` - (Optional) The device claim's top of stack serial number.

#### config_list

- `config_id` - (Optional) The config id.
- `config_parameters` - (Optional) The config's parameters. See [config_parameters](#config_parameters) below for details.

##### config_parameters

- `key` - (Optional) The config param's key.
- `value` - (Optional) The config param's value.

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

- `cli_preview` - The template's cli preview.
- `template_id` - The template preview's template id.
- `validation_errors` - The template preview's validation errors. See [validation_errors](#validation_errors) below for details.

#### validation_errors

- `type` - The validation error's type.
- `message` - The validation error's message.
