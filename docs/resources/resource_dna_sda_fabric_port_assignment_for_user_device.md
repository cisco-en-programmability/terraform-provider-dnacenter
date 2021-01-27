---
page_title: "dna_sda_fabric_port_assignment_for_user_device Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_port_assignment_for_user_device resource allows you to assign a Cisco DNA Center SDA port for a user device.
---

# Resource dna_sda_fabric_port_assignment_for_user_device

The dna_sda_fabric_port_assignment_for_user_device resource allows you to assign a Cisco DNA Center SDA port for a user device.

## Example Usage

```hcl
resource "dna_sda_fabric_port_assignment_for_user_device" "response" {
  provider   = dnacenter
  site_name_hierarchy = var.site_name_hierarchy
  device_management_ip_address = var.device_management_ip_address
  interface_name = var.interface_name
  data_ip_address_pool_name = var.data_ip_address_pool_name
  voice_ip_address_pool_name = var.voice_ip_address_pool_name
  authenticate_template_name = var.authenticate_template_name
  scalable_group_name = var.scalable_group_name
}
```

## Argument Reference

- `site_name_hierarchy` - (Required) The assignment's site name hierarchy.
- `device_management_ip_address` - (Required) The assignment's device management ip address. If it's changed it forces the creation of a new resource.
- `interface_name` - (Required) The assignment's interface name. If it's changed it forces the creation of a new resource.
- `data_ip_address_pool_name` - (Required) The assignment's data ip address pool name.
- `voice_ip_address_pool_name` - (Required) The assignment's voice ip address pool name.
- `authenticate_template_name` - (Required) The assignment's authenticate template name.
- `scalable_group_name` - (Required) The assignment's scalable group name.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The assignment's updated time with format RFC850.
