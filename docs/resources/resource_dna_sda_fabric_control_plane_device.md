---
page_title: "dna_sda_fabric_control_plane_device Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_control_plane_device resource allows you to configure a DNACenter SDA control plane device.
---

# Resource dna_sda_fabric_control_plane_device

The dna_sda_fabric_control_plane_device resource allows you to configure a DNACenter SDA control plane device.

## Example Usage

```hcl
resource "dna_sda_fabric_control_plane_device" "response" {
  provider   = dnacenter
  device_management_ip_address = device.device_ip_address
  site_name_hierarchy = device.site_name_hierarchy
}
```

## Argument Reference

- `device_management_ip_address` - (Required) The device's IP address. If it's changed it forces the creation of a new resource.
- `site_name_hierarchy` - (Required) The device's site name hierarchy. If it's changed it forces the creation of a new resource.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The device's updated time with format RFC850.
- `name` - The device's name.
- `roles` - The device's roles.
