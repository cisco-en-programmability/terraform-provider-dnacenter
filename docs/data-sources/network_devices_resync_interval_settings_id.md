---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network_devices_resync_interval_settings_id Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Fetch the reysnc interval for the given network device id.
---

# dnacenter_network_devices_resync_interval_settings_id (Data Source)

It performs read operation on Devices.

- Fetch the reysnc interval for the given network device id.

## Example Usage

```terraform
data "dnacenter_network_devices_resync_interval_settings_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_devices_resync_interval_settings_id_example" {
  value = data.dnacenter_network_devices_resync_interval_settings_id.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. The id of the network device.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `interval` (Number)