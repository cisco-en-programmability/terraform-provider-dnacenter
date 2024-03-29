---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_device_interface Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Returns all available interfaces. This endpoint can return a maximum of 500 interfacesReturns the interface for the given interface ID
---

# dnacenter_device_interface (Data Source)

It performs read operation on Devices.

- Returns all available interfaces. This endpoint can return a maximum of 500 interfaces

- Returns the interface for the given interface ID

## Example Usage

```terraform
data "dnacenter_device_interface" "example" {
  provider         = dnacenter
  last_input_time  = "string"
  last_output_time = "string"
  limit            = 1
  offset           = 1
}

output "dnacenter_device_interface_example" {
  value = data.dnacenter_device_interface.example.items
}

data "dnacenter_device_interface" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_device_interface_example" {
  value = data.dnacenter_device_interface.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter. Interface ID
- `last_input_time` (String) lastInputTime query parameter. Last Input Time
- `last_output_time` (String) lastOutputTime query parameter. Last Output Time
- `limit` (Number) limit query parameter.
- `offset` (Number) offset query parameter.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

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
- `last_updated` (String)
- `mac_address` (String)
- `mapped_physical_interface_id` (String)
- `mapped_physical_interface_name` (String)
- `media_type` (String)
- `native_vlan_id` (String)
- `ospf_support` (String)
- `pid` (String)
- `port_mode` (String)
- `port_name` (String)
- `port_type` (String)
- `serial_no` (String)
- `series` (String)
- `speed` (String)
- `status` (String)
- `vlan_id` (String)
- `voice_vlan` (String)


<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

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
- `last_updated` (String)
- `mac_address` (String)
- `mapped_physical_interface_id` (String)
- `mapped_physical_interface_name` (String)
- `media_type` (String)
- `native_vlan_id` (String)
- `ospf_support` (String)
- `pid` (String)
- `port_mode` (String)
- `port_name` (String)
- `port_type` (String)
- `serial_no` (String)
- `series` (String)
- `speed` (String)
- `status` (String)
- `vlan_id` (String)
- `voice_vlan` (String)


