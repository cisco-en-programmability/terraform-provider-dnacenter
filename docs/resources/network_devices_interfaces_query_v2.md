---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network_devices_interfaces_query_v2 Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Devices.
  This data source action returns the Interface Stats for the given Device Id. Please refer to the Feature tab for the
  Request Body usage and the API filtering support.
---

# dnacenter_network_devices_interfaces_query_v2 (Resource)

It performs create operation on Devices.

- This data source action returns the Interface Stats for the given Device Id. Please refer to the Feature tab for the
Request Body usage and the API filtering support.
~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "dnacenter_network_devices_interfaces_query_v2" "example" {
  provider  = dnacenter
  device_id = "string"
  parameters {

    end_time = 1
    query {

      fields = ["string"]
      filters {

        key      = "string"
        operator = "string"
        value    = "string"
      }
      page {

        limit  = 1
        offset = 1
        order_by {

          name  = "string"
          order = "string"
        }
      }
    }
    start_time = 1
  }
}

output "dnacenter_network_devices_interfaces_query_v2_example" {
  value = dnacenter_network_devices_interfaces_query_v2.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `device_id` (String) deviceId path parameter. Network Device Id

Optional:

- `end_time` (Number) UTC epoch timestamp in milliseconds
- `query` (Block List) (see [below for nested schema](#nestedblock--parameters--query))
- `start_time` (Number) UTC epoch timestamp in milliseconds

Read-Only:

- `items` (List of Object) (see [below for nested schema](#nestedatt--parameters--items))

<a id="nestedblock--parameters--query"></a>
### Nested Schema for `parameters.query`

Optional:

- `fields` (List of String) Required field names, default ALL
- `filters` (Block List) (see [below for nested schema](#nestedblock--parameters--query--filters))
- `page` (Block List) (see [below for nested schema](#nestedblock--parameters--query--page))

<a id="nestedblock--parameters--query--filters"></a>
### Nested Schema for `parameters.query.filters`

Optional:

- `key` (String) Name of the field that the filter should be applied to
- `operator` (String) Supported operators are eq,in,like
- `value` (String) Value of the field


<a id="nestedblock--parameters--query--page"></a>
### Nested Schema for `parameters.query.page`

Optional:

- `limit` (Number) Number of records, Max is 1000
- `offset` (Number) Record offset value, default 0
- `order_by` (Block List) (see [below for nested schema](#nestedblock--parameters--query--page--order_by))

<a id="nestedblock--parameters--query--page--order_by"></a>
### Nested Schema for `parameters.query.page.order_by`

Optional:

- `name` (String) Name of the field used to sort
- `order` (String) Possible values asc, des




<a id="nestedatt--parameters--items"></a>
### Nested Schema for `parameters.items`

Read-Only:

- `id` (String)
- `values` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--values))

<a id="nestedobjatt--parameters--items--values"></a>
### Nested Schema for `parameters.items.values`

Read-Only:

- `admin_status` (String)
- `description` (String)
- `device_id` (String)
- `duplex_config` (String)
- `duplex_oper` (String)
- `instance_id` (String)
- `interface_id` (String)
- `interface_type` (String)
- `ipv4_address` (String)
- `ipv6_address_list` (List of String)
- `is_l3_interface` (String)
- `is_wan` (String)
- `mac_addr` (String)
- `media_type` (String)
- `name` (String)
- `oper_status` (String)
- `peer_stack_member` (String)
- `peer_stack_port` (String)
- `port_channel_id` (String)
- `port_mode` (String)
- `port_type` (String)
- `rx_discards` (String)
- `rx_error` (String)
- `rx_rate` (String)
- `rx_utilization` (String)
- `speed` (String)
- `stack_port_type` (String)
- `timestamp` (String)
- `tx_discards` (String)
- `tx_error` (String)
- `tx_rate` (String)
- `tx_utilization` (String)
- `vlan_id` (String)
