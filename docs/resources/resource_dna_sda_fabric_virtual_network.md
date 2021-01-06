---
page_title: "dna_sda_fabric_virtual_network Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_virtual_network resource allows you to configure a DNACenter SDA virtual network.
---

# Resource dna_sda_fabric_virtual_network

The dna_sda_fabric_virtual_network resource allows you to configure a DNACenter SDA virtual network.

## Example Usage

```hcl
resource "dna_sda_fabric_virtual_network" "response" {
  provider   = dnacenter
  site_name_hierarchy = var.site_name_hierarchy
  virtual_network_name = var.virtual_network_name
}
```

## Argument Reference

- `virtual_network_name` - (Required) The virtual network's virtual network name. If it's changed it forces the creation of a new resource.
- `site_name_hierarchy` - (Required) The virtual network's site name hierarchy. If it's changed it forces the creation of a new resource.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The virtual network's updated time with format RFC850.
- `device_management_ip_address` - The virtual network's device management ip address.
- `roles` - The virtual network's roles.
