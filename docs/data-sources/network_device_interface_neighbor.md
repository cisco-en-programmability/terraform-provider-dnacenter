---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network_device_interface_neighbor Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Get connected device detail for given deviceUuid and interfaceUuid
---

# dnacenter_network_device_interface_neighbor (Data Source)

It performs read operation on Devices.

- Get connected device detail for given deviceUuid and interfaceUuid

## Example Usage

```terraform
data "dnacenter_network_device_interface_neighbor" "example" {
  provider       = dnacenter
  device_uuid    = "string"
  interface_uuid = "string"
}

output "dnacenter_network_device_interface_neighbor_example" {
  value = data.dnacenter_network_device_interface_neighbor.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_uuid` (String) deviceUuid path parameter. instanceuuid of Device
- `interface_uuid` (String) interfaceUuid path parameter. instanceuuid of interface

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `capabilities` (List of String)
- `neighbor_device` (String)
- `neighbor_port` (String)
