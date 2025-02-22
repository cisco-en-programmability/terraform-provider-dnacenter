---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network_device_range Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Returns the list of network devices for the given pagination range. The maximum number of records that can be
  retrieved is 500
---

# dnacenter_network_device_range (Data Source)

It performs read operation on Devices.

- Returns the list of network devices for the given pagination range. The maximum number of records that can be
retrieved is 500

## Example Usage

```terraform
data "dnacenter_network_device_range" "example" {
  provider          = dnacenter
  records_to_return = 1
  start_index       = 1
}

output "dnacenter_network_device_range_example" {
  value = data.dnacenter_network_device_range.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `records_to_return` (Number) recordsToReturn path parameter. Number of records to return [1<= recordsToReturn <= 500]
- `start_index` (Number) startIndex path parameter. Start index [>=1]

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `ap_ethernet_mac_address` (String)
- `ap_manager_interface_ip` (String)
- `associated_wlc_ip` (String)
- `boot_date_time` (String)
- `collection_interval` (String)
- `collection_status` (String)
- `description` (String)
- `device_support_level` (String)
- `dns_resolved_management_address` (String)
- `error_code` (String)
- `error_description` (String)
- `family` (String)
- `hostname` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `interface_count` (String)
- `inventory_status_detail` (String)
- `last_device_resync_start_time` (String)
- `last_update_time` (Number)
- `last_updated` (String)
- `line_card_count` (String)
- `line_card_id` (String)
- `location` (String)
- `location_name` (String)
- `mac_address` (String)
- `managed_atleast_once` (String)
- `management_ip_address` (String)
- `management_state` (String)
- `memory_size` (String)
- `pending_sync_requests_count` (String)
- `platform_id` (String)
- `reachability_failure_reason` (String)
- `reachability_status` (String)
- `reasons_for_device_resync` (String)
- `reasons_for_pending_sync_requests` (String)
- `role` (String)
- `role_source` (String)
- `serial_number` (String)
- `series` (String)
- `snmp_contact` (String)
- `snmp_location` (String)
- `software_type` (String)
- `software_version` (String)
- `tag_count` (String)
- `tunnel_udp_port` (String)
- `type` (String)
- `up_time` (String)
- `uptime_seconds` (Number)
- `vendor` (String)
- `waas_device_mode` (String)
