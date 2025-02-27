---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_discovery Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Discovery.
  Returns discovery by Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.
---

# dnacenter_discovery (Data Source)

It performs read operation on Discovery.

- Returns discovery by Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.

## Example Usage

```terraform
data "dnacenter_discovery" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_discovery_example" {
  value = data.dnacenter_discovery.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. Discovery ID

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `attribute_info` (String)
- `cdp_level` (Number)
- `device_ids` (String)
- `discovery_condition` (String)
- `discovery_status` (String)
- `discovery_type` (String)
- `enable_password_list` (String)
- `global_credential_id_list` (List of String)
- `http_read_credential` (List of Object) (see [below for nested schema](#nestedobjatt--item--http_read_credential))
- `http_write_credential` (List of Object) (see [below for nested schema](#nestedobjatt--item--http_write_credential))
- `id` (String)
- `ip_address_list` (String)
- `ip_filter_list` (String)
- `is_auto_cdp` (String)
- `lldp_level` (Number)
- `name` (String)
- `netconf_port` (String)
- `num_devices` (Number)
- `parent_discovery_id` (String)
- `password_list` (String)
- `preferred_mgmt_ipmethod` (String)
- `protocol_order` (String)
- `retry_count` (Number)
- `snmp_auth_passphrase` (String)
- `snmp_auth_protocol` (String)
- `snmp_mode` (String)
- `snmp_priv_passphrase` (String)
- `snmp_priv_protocol` (String)
- `snmp_ro_community` (String)
- `snmp_ro_community_desc` (String)
- `snmp_rw_community` (String)
- `snmp_rw_community_desc` (String)
- `snmp_user_name` (String)
- `time_out` (Number)
- `update_mgmt_ip` (String)
- `user_name_list` (String)

<a id="nestedobjatt--item--http_read_credential"></a>
### Nested Schema for `item.http_read_credential`

Read-Only:

- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `password` (String)
- `port` (Number)
- `secure` (String)
- `username` (String)


<a id="nestedobjatt--item--http_write_credential"></a>
### Nested Schema for `item.http_write_credential`

Read-Only:

- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `password` (String)
- `port` (Number)
- `secure` (String)
- `username` (String)
