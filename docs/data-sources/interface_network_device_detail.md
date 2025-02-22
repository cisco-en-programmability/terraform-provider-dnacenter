---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_interface_network_device_detail Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Returns interface by specified device Id and interface name
---

# dnacenter_interface_network_device_detail (Data Source)

It performs read operation on Devices.

- Returns interface by specified device Id and interface name

## Example Usage

```terraform
data "dnacenter_interface_network_device_detail" "example" {
  provider  = dnacenter
  device_id = "string"
  name      = "string"
}

output "dnacenter_interface_network_device_detail_example" {
  value = data.dnacenter_interface_network_device_detail.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (String) deviceId path parameter. Device ID
- `name` (String) name query parameter. Interface name

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `addresses` (List of Object) (see [below for nested schema](#nestedobjatt--item--addresses))
- `admin_status` (String)
- `class_name` (String)
- `description` (String)
- `device_id` (String)
- `duplex` (String)
- `id` (String)
- `if_index` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `interface_type` (String)
- `ipv4_address` (String)
- `ipv4_mask` (String)
- `isis_support` (String)
- `last_incoming_packet_time` (Number)
- `last_outgoing_packet_time` (Number)
- `last_updated` (String)
- `mac_address` (String)
- `managed_compute_element` (String)
- `managed_compute_element_url` (String)
- `managed_network_element` (String)
- `managed_network_element_url` (String)
- `mapped_physical_interface_id` (String)
- `mapped_physical_interface_name` (String)
- `media_type` (String)
- `mtu` (String)
- `name` (String)
- `native_vlan_id` (String)
- `networkdevice_id` (String)
- `ospf_support` (String)
- `pid` (String)
- `port_mode` (String)
- `port_name` (String)
- `port_type` (String)
- `poweroverethernet` (String)
- `serial_no` (String)
- `series` (String)
- `speed` (String)
- `status` (String)
- `vlan_id` (String)
- `voice_vlan` (String)

<a id="nestedobjatt--item--addresses"></a>
### Nested Schema for `item.addresses`

Read-Only:

- `address` (List of Object) (see [below for nested schema](#nestedobjatt--item--addresses--address))
- `type` (String)

<a id="nestedobjatt--item--addresses--address"></a>
### Nested Schema for `item.addresses.address`

Read-Only:

- `ip_address` (List of Object) (see [below for nested schema](#nestedobjatt--item--addresses--address--ip_address))
- `ip_mask` (List of Object) (see [below for nested schema](#nestedobjatt--item--addresses--address--ip_mask))
- `is_inverse_mask` (String)

<a id="nestedobjatt--item--addresses--address--ip_address"></a>
### Nested Schema for `item.addresses.address.is_inverse_mask`

Read-Only:

- `address` (String)


<a id="nestedobjatt--item--addresses--address--ip_mask"></a>
### Nested Schema for `item.addresses.address.is_inverse_mask`

Read-Only:

- `address` (String)
