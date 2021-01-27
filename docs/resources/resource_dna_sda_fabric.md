---
page_title: "dna_sda_fabric_virtual_network Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_virtual_network resource allows you to configure a Cisco DNA Center SDA virtual network.
---

# Resource dna_sda_fabric_virtual_network

The dna_sda_fabric_virtual_network resource allows you to configure a Cisco DNA Center SDA virtual network.

## Example Usage

```hcl
resource "dna_sda_fabric_virtual_network" "response" {
  provider   = dnacenter
  fabric_name = var.fabric_name
}
```

## Argument Reference

- `fabric_name` - (Required) The fabric's name. If it's changed it forces the creation of a new resource.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The fabric's updated time with format RFC850.
- `fabric_type` - The fabric's type.
- `fabric_domain_type` - The fabric's domain type.
