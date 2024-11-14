---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_sda_port_assignments_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Returns the count of port assignments that match the provided query parameters.
---

# dnacenter_sda_port_assignments_count (Data Source)

It performs read operation on SDA.

- Returns the count of port assignments that match the provided query parameters.

## Example Usage

```terraform
data "dnacenter_sda_port_assignments_count" "example" {
  provider          = dnacenter
  data_vlan_name    = "string"
  fabric_id         = "string"
  interface_name    = "string"
  network_device_id = "string"
  voice_vlan_name   = "string"
}

output "dnacenter_sda_port_assignments_count_example" {
  value = data.dnacenter_sda_port_assignments_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `data_vlan_name` (String) dataVlanName query parameter. Data VLAN name of the port assignment.
- `fabric_id` (String) fabricId query parameter. ID of the fabric the device is assigned to.
- `interface_name` (String) interfaceName query parameter. Interface name of the port assignment.
- `network_device_id` (String) networkDeviceId query parameter. Network device ID of the port assignment.
- `voice_vlan_name` (String) voiceVlanName query parameter. Voice VLAN name of the port assignment.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)