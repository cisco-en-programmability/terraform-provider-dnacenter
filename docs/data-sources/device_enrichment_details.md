---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_device_enrichment_details Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Enriches a given network device context (device id or device Mac Address or device management IP address) with details
  about the device and neighbor topology
---

# dnacenter_device_enrichment_details (Data Source)

It performs read operation on Devices.

- Enriches a given network device context (device id or device Mac Address or device management IP address) with details
about the device and neighbor topology

## Example Usage

```terraform
data "dnacenter_device_enrichment_details" "example" {
  provider     = dnacenter
  entity_type  = "string"
  entity_value = "string"
}

output "dnacenter_device_enrichment_details_example" {
  value = data.dnacenter_device_enrichment_details.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `entity_type` (String) entity_type header parameter. Device enrichment details can be fetched based on either Device ID or Device MAC address or Device IP Address. This parameter value must either be device_id/mac_address/ip_address
- `entity_value` (String) entity_value header parameter. Contains the actual value for the entity type that has been defined

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `device_details` (List of Object) (see [below for nested schema](#nestedobjatt--items--device_details))

<a id="nestedobjatt--items--device_details"></a>
### Nested Schema for `items.device_details`

Read-Only:

- `ap_manager_interface_ip` (String)
- `associated_wlc_ip` (String)
- `boot_date_time` (String)
- `collection_interval` (String)
- `collection_status` (String)
- `error_code` (String)
- `error_description` (String)
- `family` (String)
- `hostname` (String)
- `id` (String)
- `instance_uuid` (String)
- `interface_count` (String)
- `inventory_status_detail` (String)
- `last_update_time` (Number)
- `last_updated` (String)
- `line_card_count` (String)
- `line_card_id` (String)
- `location` (String)
- `location_name` (String)
- `mac_address` (String)
- `management_ip_address` (String)
- `memory_size` (String)
- `neighbor_topology` (List of Object) (see [below for nested schema](#nestedobjatt--items--device_details--neighbor_topology))
- `platform_id` (String)
- `reachability_failure_reason` (String)
- `reachability_status` (String)
- `role` (String)
- `role_source` (String)
- `serial_number` (String)
- `series` (String)
- `snmp_contact` (String)
- `snmp_location` (String)
- `software_version` (String)
- `tag_count` (String)
- `tunnel_udp_port` (String)
- `type` (String)
- `up_time` (String)
- `waas_device_mode` (String)

<a id="nestedobjatt--items--device_details--neighbor_topology"></a>
### Nested Schema for `items.device_details.neighbor_topology`

Read-Only:

- `links` (List of Object) (see [below for nested schema](#nestedobjatt--items--device_details--neighbor_topology--links))
- `nodes` (List of Object) (see [below for nested schema](#nestedobjatt--items--device_details--neighbor_topology--nodes))

<a id="nestedobjatt--items--device_details--neighbor_topology--links"></a>
### Nested Schema for `items.device_details.neighbor_topology.nodes`

Read-Only:

- `id` (String)
- `label` (List of String)
- `link_status` (String)
- `port_utilization` (String)
- `source` (String)
- `target` (String)


<a id="nestedobjatt--items--device_details--neighbor_topology--nodes"></a>
### Nested Schema for `items.device_details.neighbor_topology.nodes`

Read-Only:

- `clients` (String)
- `connected_device` (String)
- `count` (String)
- `description` (String)
- `device_type` (String)
- `fabric_group` (String)
- `family` (String)
- `health_score` (Number)
- `id` (String)
- `ip` (String)
- `level` (Number)
- `name` (String)
- `node_type` (String)
- `platform_id` (String)
- `radio_frequency` (String)
- `role` (String)
- `software_version` (String)
- `user_id` (String)
