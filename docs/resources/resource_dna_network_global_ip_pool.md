---
page_title: "dna_network_global_ip_pool Resource - terraform-provider-dnacenter"
subcategory: "Network Settings"
description: |-
  The dna_network_global_ip_pool resource allows you to configure a Cisco DNA Center global IP pool.
---

# Resource dna_network_global_ip_pool

The dna_network_global_ip_pool resource allows you to configure a Cisco DNA Center global IP pool.

## Example Usage

```hcl
resource "dna_network_global_ip_pool" "response" {
  provider         = dnacenter
  type             = "Generic"
  gateway          = ""
  ip_address_space = "IPv4"
  item {
    id             = "22f70f75-5dae-4494-9965-d4b85e101898"
    ip_pool_name   = "dna-usa"
    dns_server_ips = ["34.245.38.218"]
    ip_pool_cidr   = "10.64.0.0/12"
  }
}
```

## Argument Reference

- `item` - (Required) Item in a Cisco DNA Center Global IP Pool. See [Global IP Pool item](#global-ip-pool-item) below for details.
- `type` - (Required) The Global IP pool's type.
- `gateway` - (Required) The Global IP pool's gateway.
- `ip_address_space` - (Required) The Global IP pool's IP address space.

### global ip pool item

- `dhcp_server_ips` - (Optional) The IP pool's DHCP server IPs.
- `dns_server_ips` - (Optional) The IP pool's DNS server IPs.
- `id` - (Optional) The IP pool's id.
- `ip_pool_cidr` - (Required) The IP pool's CIDR.
- `ip_pool_name` - (Required) The IP pool's name.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The Global IP pool's updated time with format RFC850.
- `item` - Item in a Cisco DNA Center Global IP Pool. See [Global IP Pool item](#global-ip-pool-item-1) below for details.

### global ip pool item

- `client_options` - The Global IP pool item's client options.
- `configure_external_dhcp` - The Global IP pool item's configure external DHCP flag.
- `context` - The Global IP pool item's context. See [Context item](#context-item) below for details.
- `owner` - The Global IP pool item's owner.
- `create_time` - The Global IP pool item's create time.
- `gateways` - The Global IP pool item's gateways.
- `ipv6` - The Global IP pool item's IPv6.
- `last_update_time` - The Global IP pool item's last update time.
- `overlapping` - The Global IP pool item's overlapping.
- `owner` - The Global IP pool item's owner.
- `parent_uuid` - The Global IP pool item's parent uuid.
- `shared` - The Global IP pool item's shared.
- `total_ip_address_count` - The Global IP pool item's total IP address count.
- `used_ip_address_count` - The Global IP pool item's used IP addresscount.
- `used_percentage` - The Global IP pool item's used percentage.

#### context item

- `context_key` - The context's key.
- `context_value` - The context's value.
- `owner` - The context's owner.
