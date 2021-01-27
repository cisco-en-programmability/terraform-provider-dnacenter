---
page_title: "dna_discovery Resource - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_discovery resource allows you to configure a Cisco DNA Center Discovery.
---

# Resource dna_discovery

The dna_discovery resource allows you to configure a Cisco DNA Center Discovery.

## Example Usage

```hcl
resource "dna_discovery" "response" {
  provider = dnacenter
  item {
    cdp_level                 = 16
    discovery_type            = "CDP"
    global_credential_id_list = ["90acbab8-03d5-4726-9c19-e1e51a40b3cd", "f979d842-f6fd-456a-8137-2cb5113cd2e8"]
    ip_address_list           = "10.10.22.22"
    name                      = "start_discovery_test"
    netconf_port              = "65535"
    protocol_order            = "ssh"
  }
}
```

## Argument Reference

- `item` - (Required) Item in a Cisco DNA Center Discovery. See [Discovery item](#discovery-item) below for details.

### Discovery item

Each Discovery item contains the following arguments.

- `comments` - (Required) The Discovery's comments.
- `id` - (Optional) The Discovery's id.
- `cdp_level` - (Optional) The Discovery's cdp level.
- `discovery_type` - (Required) The Discovery's discovery type.
- `enable_password_list` - (Optional) The Discovery's enable password list.
- `global_credential_id_list` - (Optional) The Discovery's global credential id list.
- `http_read_credential` - (Optional) The Discovery's HTTP Credential for read. See [Discovery Credential](#discovery-credential) below for details.
- `http_write_credential` - (Optional) The Discovery's HTTP Credential for write. See [Discovery Credential](#discovery-credential) below for details.
- `ip_address_list` - (Required) The Discovery's IP address list.
- `ip_filter_list` - (Optional) The Discovery's IP filter list.
- `lldp_level` - (Optional) The Discovery's LLDP level.
- `name` - (Required) The Discovery's name.
- `netconf_port` - (Optional) The Discovery's Netconf port.
- `no_add_new_device` - (Optional) The Discovery's flag to no add new device.
- `parent_discovery_id` - (Optional) The Discovery's parent discovery id.
- `password_list` - (Optional) The Discovery's password list.
- `preferred_mgmt_ip_method` - (Optional) The Discovery's preferred mgmt ip method.
- `protocol_order` - (Optional) The Discovery's protocol_order.
- `re_discovery` - (Optional) The Discovery's re-discovery.
- `retry` - (Optional) The Discovery's retry.
- `snmp_auth_passphrase` - (Optional) The Discovery's SNMP auth passphrase.
- `snmp_auth_protocol` - (Optional) The Discovery's SNMP auth protocol.
- `snmp_mode` - (Optional) The Discovery's SNMP mode.
- `snmp_priv_passphrase` - (Optional) The Discovery's SNMP priv passphrase.
- `snmp_priv_protocol` - (Optional) The Discovery's SNMP priv protocol.
- `snmp_ro_community` - (Optional) The Discovery's SNMP ro community.
- `snmp_ro_community_desc` - (Optional) The Discovery's SNMP ro community desc.
- `snmp_rw_community` - (Optional) The Discovery's SNMP rw community.
- `snmp_rw_community_desc` - (Optional) The Discovery's SNMP rw community desc.
- `snmp_user_name` - (Optional) The Discovery's SNMP user name.
- `snmp_version` - (Optional) The Discovery's SNMP version.
- `timeout` - (Optional) The Discovery's timeout.
- `update_mgmt_ip` - (Optional) The Discovery's update mgmt ip.
- `user_name_list` - (Optional) The Discovery's user name list.
- `device_ids` - (Optional) The Discovery's device ids.
- `num_devices` - (Optional) The Discovery's num devices.
- `discovery_condition` - (Optional) The Discovery's discovery condition.
- `discovery_status` - (Optional) The Discovery's discovery status.
- `is_auto_cdp` - (Optional) The Discovery's flag if is auto cdp.

### Discovery credential

Each HTTP credential contains the following arguments.

- `comments` - (Optional) The HTTP credential's comments.
- `credential_type` - (Optional) The HTTP credential's type.
- `description` - (Optional) The HTTP credential's description.
- `id` - (Optional) The HTTP credential's id.
- `instance_tenant_id` - (Optional) The HTTP credential's instance tenant id.
- `instance_uuid` - (Optional) The HTTP credential's instance uuid.
- `password` - (Optional) The HTTP credential's password.
- `port` - (Optional) The HTTP credential's port.
- `secure` - (Optional) The HTTP credential's secure.
- `username` - (Optional) The HTTP credential's username.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The Discovery's updated time with format RFC850.
