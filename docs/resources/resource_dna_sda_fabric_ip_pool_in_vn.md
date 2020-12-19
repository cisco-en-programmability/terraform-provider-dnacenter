---
page_title: "dna_sda_fabric_ip_pool_in_vn Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_ip_pool_in_vn resource allows you to configure a DNACenter SDA IP pool in Virtual Network.
---

# Resource dna_sda_fabric_ip_pool_in_vn

The dna_sda_fabric_ip_pool_in_vn resource allows you to configure a DNACenter SDA IP pool in Virtual Network.

## Example Usage

```hcl
resource "dna_sda_fabric_ip_pool_in_vn" "response" {
  provider   = dnacenter
  virtual_network_name = pool.virtual_network_name
  ip_pool_name = pool.ip_pool_name
  traffic_type = pool.traffic_type
  authentication_policy_name = pool.authentication_policy_name
  scalable_group_name = pool.scalable_group_name
  is_l2_flooding_enabled = pool.is_l2_flooding_enabled
  is_this_critical_pool = pool.is_this_critical_pool
  pool_type = pool.pool_type
}
```

## Argument Reference

- `virtual_network_name` - (Required) The pool's virtual network name. If it's changed it forces the creation of a new resource.
- `ip_pool_name` - (Required) The pool's ip pool name. If it's changed it forces the creation of a new resource.
- `traffic_type` - (Required) The pool's traffic type.
- `authentication_policy_name` - (Required) The pool's authentication policy name.
- `scalable_group_name` - (Required) The pool's scalable group name.
- `is_l2_flooding_enabled` - (Required) The pool's is l2 flooding enabled.
- `is_this_critical_pool` - (Required) The pool's is this critical pool.
- `pool_type` - (Required) The pool's pool_type.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The pool's updated time with format RFC850.
