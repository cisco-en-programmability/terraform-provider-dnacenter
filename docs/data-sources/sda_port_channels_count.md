---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_sda_port_channels_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Returns the count of port channels that match the provided query parameters.
---

# dnacenter_sda_port_channels_count (Data Source)

It performs read operation on SDA.

- Returns the count of port channels that match the provided query parameters.

## Example Usage

```terraform
data "dnacenter_sda_port_channels_count" "example" {
  provider              = dnacenter
  connected_device_type = "string"
  fabric_id             = "string"
  network_device_id     = "string"
  port_channel_name     = "string"
}

output "dnacenter_sda_port_channels_count_example" {
  value = data.dnacenter_sda_port_channels_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `connected_device_type` (String) connectedDeviceType query parameter. Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
- `fabric_id` (String) fabricId query parameter. ID of the fabric the device is assigned to.
- `network_device_id` (String) networkDeviceId query parameter. ID of the network device.
- `port_channel_name` (String) portChannelName query parameter. Name of the port channel.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)