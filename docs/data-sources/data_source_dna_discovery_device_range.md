---
page_title: "dna_discovery_device_range Data Source - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_discovery_device_range data source allows you to retrieve information about a particular Cisco DNA Center discovery device by range.
---

# Data Source dna_discovery_device_range

The dna_discovery_device_range data source allows you to retrieve information about a particular Cisco DNA Center discovery device by range.

## Example Usage

```hcl
data "dna_discovery_device_range" "list" {
  provider = dnacenter
  id = var.discovery_id
  start_index = 1
  records_to_return = 4
}
```

## Argument Reference

- `task_id` - (Optional) The task_id param.
- `id` - (Required) The id param.
- `start_index` - (Required) The start_index param.
- `records_to_return` - (Required) The records_to_return param.

## Attributes Reference

The following attributes are exported.

- `items` - Items in a Cisco DNA Center discovery device. See [Items](#items) below for details.

### Items

Each item contains the following attributes.

- `anchor_wlc_for_ap` - The item anchor wlc for ap.
- `auth_model_id` - The item auth model id.
- `avg_update_frequency` - The item avg update frequency.
- `boot_date_time` - The item boot date time.
- `cli_status` - The item cli status.
- `duplicate_device_id` - The item duplicate device id.
- `error_code` - The item error code.
- `error_description` - The item error description.
- `family` - The item family.
- `hostname` - The item hostname.
- `http_status` - The item http status.
- `id` - The item id.
- `image_name` - The item image name.
- `ingress_queue_config` - The item ingress queue config.
- `interface_count` - The item interface count.
- `inventory_collection_status` - The item inventory collection status.
- `inventory_reachability_status` - The item inventory reachability status.
- `last_updated` - The item last updated.
- `line_card_count` - The item line card count.
- `line_card_id` - The item line card id.
- `location` - The item location.
- `location_name` - The item location name.
- `mac_address` - The item mac address.
- `management_ip_address` - The item management ip address.
- `memory_size` - The item memory size.
- `netconf_status` - The item netconf status.
- `num_updates` - The item num updates.
- `ping_status` - The item ping status.
- `platform_id` - The item platform id.
- `port_range` - The item port range.
- `qos_status` - The item qos status.
- `reachability_failure_reason` - The item reachability failure reason.
- `reachability_status` - The item reachability status.
- `role` - The item role.
- `role_source` - The item role source.
- `serial_number` - The item serial number.
- `snmp_contact` - The item snmp contact.
- `snmp_location` - The item snmp location.
- `snmp_status` - The item snmp status.
- `software_version` - The item software version.
- `tag` - The item tag.
- `tag_count` - The item tag count.
- `type` - The item type.
- `up_time` - The item up time.
- `vendor` - The item vendor.
- `wlc_ap_device_status` - The item wlc ap device status.
