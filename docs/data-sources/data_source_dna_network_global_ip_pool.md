---
page_title: "dna_network_global_ip_pool Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_network_global_ip_pool data source allows you to retrieve information about a particular DNACenter global IP pools.
---

# Data Source dna_network_global_ip_pool

The dna_network_global_ip_pool data source allows you to retrieve information about a particular DNACenter global IP pools.

## Example Usage

```hcl
data "dna_network_global_ip_pool" "response" {
  provider = dnacenter
}
```

## Argument Reference

- `offset` - (Optional) The offset param.
- `limit` - (Optional) The limit param.

## Attributes Reference

The following attributes are exported.

- `items` - Items in a DNACenter global IP pool. See [Global IP Pool item](#global-ip-pool-item-1) below for details.

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
