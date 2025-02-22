---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_sda_anycast_gateways_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Returns the count of anycast gateways that match the provided query parameters.
---

# dnacenter_sda_anycast_gateways_count (Data Source)

It performs read operation on SDA.

- Returns the count of anycast gateways that match the provided query parameters.

## Example Usage

```terraform
data "dnacenter_sda_anycast_gateways_count" "example" {
  provider             = dnacenter
  fabric_id            = "string"
  ip_pool_name         = "string"
  virtual_network_name = "string"
  vlan_id              = 1.0
  vlan_name            = "string"
}

output "dnacenter_sda_anycast_gateways_count_example" {
  value = data.dnacenter_sda_anycast_gateways_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `fabric_id` (String) fabricId query parameter. ID of the fabric the anycast gateway is assigned to.
- `ip_pool_name` (String) ipPoolName query parameter. Name of the IP pool associated with the anycast gateways.
- `virtual_network_name` (String) virtualNetworkName query parameter. Name of the virtual network associated with the anycast gateways.
- `vlan_id` (Number) vlanId query parameter. VLAN ID of the anycast gateways. The allowed range for vlanId is [2-4093] except for reserved VLANs [1002-1005], 2046, and 4094.
- `vlan_name` (String) vlanName query parameter. VLAN name of the anycast gateways.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)
