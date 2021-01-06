---
page_title: "dna_sda_fabric_border_device Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_border_device resource allows you to configure a DNACenter SDA border device.
---

# Resource dna_sda_fabric_border_device

The dna_sda_fabric_border_device resource allows you to configure a DNACenter SDA border device.

## Example Usage

```hcl
resource "dna_sda_fabric_border_device" "response" {
  provider   = dnacenter
  item {
  }
  device_ip_address = var.device_ip_address
  site_name_hierarchy = var.site_name_hierarchy
  external_domain_routing_protocol_name = var.external_domain_routing_protocol_name
  external_connectivity_ip_pool_name = var.external_connectivity_ip_pool_name
  internal_autonomous_system_number = var.internal_autonomous_system_number
  border_session_type = var.border_session_type
  connected_to_internet = var.connected_to_internet
  external_connectivity_settings {
    interface_name = "InterfaceName"
    external_autonomous_system_number = 1
    l3_handoff {
      virtual_network {
        virtual_network_name = "VN_name"
      }
    }
  }
}
```

## Argument Reference

- `item` - (Required) The device's item. See [Device item](#device-item) below for details.
- `device_ip_address` - (Required) The device's IP address. If it's changed it forces the creation of a new resource.
- `site_name_hierarchy` - (Required) The device's site name hierarchy. If it's changed it forces the creation of a new resource.
- `external_domain_routing_protocol_name` - (Required) The device's external domain routing protocol name.
- `external_connectivity_ip_pool_name` - (Required) The device's external connectivity IP pool name.
- `internal_autonomous_system_number` - (Required) The device's internal autonomous system number.
- `border_session_type` - (Required) The device's session type.
- `connected_to_internet` - (Required) The device's connected to internet.
- `external_connectivity_settings` - (Required) The device's external connectivity settings. See [external_connectivity_settings](#external_connectivity_settings) below for details.

### Device item

All of this elements are computed. Check the GET method for fabric border device from the DNACenter API for more details.

### external_connectivity_settings

- `interface_name` - (Required) The border device's interface name.
- `external_autonomous_system_number` - (Required) The border device's external autonomous system number.
- `l3_handoff` - (Required) The border device's l3 handoff. See below for details.

#### l3_handoff

- `virtual_network` - (Required) l3 handoff's virtual network.

#### virtual_network

- `virtual_network_name` - (Required) l3 handoff's virtual network.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The device's updated time with format RFC850.
