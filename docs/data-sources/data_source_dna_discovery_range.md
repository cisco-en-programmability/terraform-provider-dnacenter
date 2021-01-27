---
page_title: "dna_discovery_range Data Source - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_discovery_range data source allows you to retrieve information about Cisco DNA Center discoveries device by range.
---

# Data Source dna_discovery_range

The dna_discovery_range data source allows you to retrieve information about Cisco DNA Center discoveries device by range.

## Example Usage

```hcl
data "dna_discovery_range" "response" {
  provider = dnacenter
  start_index = 1
  records_to_return = 2
}
```

## Argument Reference

- `start_index` - (Required) The start_index param.
- `records_to_return` - (Required) The records_to_return param.

## Attributes Reference

In addition to all the attributes above, the following attributes are exported.

- `items` - Item in a Cisco DNA Center Discovery. See [Discovery item](#discovery-item) below for details.

### Discovery item

Each Discovery item contains the following attributes.

- `comments` - The Discovery's comments.
- `id` - The Discovery's id.
- `cdp_level` - The Discovery's cdp level.
- `discovery_type` - The Discovery's discovery type.
- `enable_password_list` - The Discovery's enable password list.
- `global_credential_id_list` - The Discovery's global credential id list.
- `http_read_credential` - The Discovery's HTTP Credential for read. See [Discovery Credential](#discovery-credential) below for details.
- `http_write_credential` - The Discovery's HTTP Credential for write. See [Discovery Credential](#discovery-credential) below for details.
- `ip_address_list` - The Discovery's IP address list.
- `ip_filter_list` - The Discovery's IP filter list.
- `lldp_level` - The Discovery's LLDP level.
- `name` - The Discovery's name.
- `netconf_port` - The Discovery's Netconf port.
- `no_add_new_device` - The Discovery's flag to no add new device.
- `parent_discovery_id` - The Discovery's parent discovery id.
- `password_list` - The Discovery's password list.
- `preferred_mgmt_ip_method` - The Discovery's preferred mgmt ip method.
- `protocol_order` - The Discovery's protocol_order.
- `re_discovery` - The Discovery's re-discovery.
- `retry` - The Discovery's retry.
- `snmp_auth_passphrase` - The Discovery's SNMP auth passphrase.
- `snmp_auth_protocol` - The Discovery's SNMP auth protocol.
- `snmp_mode` - The Discovery's SNMP mode.
- `snmp_priv_passphrase` - The Discovery's SNMP priv passphrase.
- `snmp_priv_protocol` - The Discovery's SNMP priv protocol.
- `snmp_ro_community` - The Discovery's SNMP ro community.
- `snmp_ro_community_desc` - The Discovery's SNMP ro community desc.
- `snmp_rw_community` - The Discovery's SNMP rw community.
- `snmp_rw_community_desc` - The Discovery's SNMP rw community desc.
- `snmp_user_name` - The Discovery's SNMP user name.
- `snmp_version` - The Discovery's SNMP version.
- `timeout` - The Discovery's timeout.
- `update_mgmt_ip` - The Discovery's update mgmt ip.
- `user_name_list` - The Discovery's user name list.
- `device_ids` - The Discovery's device ids.
- `num_devices` - The Discovery's num devices.
- `discovery_condition` - The Discovery's discovery condition.
- `discovery_status` - The Discovery's discovery status.
- `is_auto_cdp` - The Discovery's flag if is auto cdp.

### Discovery credential

Each HTTP credential contains the following attributes.

- `comments` - The HTTP credential's comments.
- `credential_type` - The HTTP credential's type.
- `description` - The HTTP credential's description.
- `id` - The HTTP credential's id.
- `instance_tenant_id` - The HTTP credential's instance tenant id.
- `instance_uuid` - The HTTP credential's instance uuid.
- `password` - The HTTP credential's password.
- `port` - The HTTP credential's port.
- `secure` - The HTTP credential's secure.
- `username` - The HTTP credential's username.

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
