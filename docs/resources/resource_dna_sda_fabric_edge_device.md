---
page_title: "dna_sda_fabric_edge_device Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_edge_device resource allows you to configure a DNACenter SDA edge device.
---

# Resource dna_sda_fabric_edge_device

The dna_sda_fabric_edge_device resource allows you to configure a DNACenter SDA edge device.

## Example Usage

```hcl
resource "dna_sda_fabric_edge_device" "response" {
  provider   = dnacenter
  device_management_ip_address = var.device_ip_address
  site_name_hierarchy = var.site_name_hierarchy
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
