---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_client_detail Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Clients.
  Returns detailed Client information retrieved by Mac Address for any given point of time.
---

# dnacenter_client_detail (Data Source)

It performs read operation on Clients.

- Returns detailed Client information retrieved by Mac Address for any given point of time.

## Example Usage

```terraform
data "dnacenter_client_detail" "example" {
  provider    = dnacenter
  mac_address = "string"
  timestamp   = 1.0
}

output "dnacenter_client_detail_example" {
  value = data.dnacenter_client_detail.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `mac_address` (String) macAddress query parameter. MAC Address of the client

### Optional

- `timestamp` (Number) timestamp query parameter. Epoch time(in milliseconds) when the Client health data is required

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `connection_info` (List of Object) (see [below for nested schema](#nestedobjatt--item--connection_info))
- `detail` (List of Object) (see [below for nested schema](#nestedobjatt--item--detail))
- `topology` (List of Object) (see [below for nested schema](#nestedobjatt--item--topology))

<a id="nestedobjatt--item--connection_info"></a>
### Nested Schema for `item.connection_info`

Read-Only:

- `band` (String)
- `channel` (String)
- `channel_width` (String)
- `host_type` (String)
- `nw_device_mac` (String)
- `nw_device_name` (String)
- `protocol` (String)
- `spatial_stream` (String)
- `timestamp` (Number)
- `uapsd` (String)
- `wmm` (String)


<a id="nestedobjatt--item--detail"></a>
### Nested Schema for `item.detail`

Read-Only:

- `aaa_server_eap_latency` (Number)
- `aaa_server_failed_transaction` (Number)
- `aaa_server_ip` (String)
- `aaa_server_latency` (Number)
- `aaa_server_mab_latency` (Number)
- `aaa_server_success_transaction` (Number)
- `aaa_server_transaction` (Number)
- `ap_group` (String)
- `auth_type` (String)
- `avg_rssi` (String)
- `avg_snr` (String)
- `bridge_vmmode` (String)
- `channel` (String)
- `client_connection` (String)
- `client_type` (String)
- `connected_device` (List of Object) (see [below for nested schema](#nestedobjatt--item--detail--connected_device))
- `connected_upn` (String)
- `connected_upn_id` (String)
- `connected_upn_owner` (String)
- `connection_status` (String)
- `country_code` (String)
- `data_rate` (String)
- `device_form` (String)
- `device_vendor` (String)
- `dhcp_decline_ip` (String)
- `dhcp_nak_ip` (String)
- `dhcp_server_dolatency` (Number)
- `dhcp_server_failed_transaction` (Number)
- `dhcp_server_ip` (String)
- `dhcp_server_latency` (Number)
- `dhcp_server_ralatency` (Number)
- `dhcp_server_success_transaction` (Number)
- `dhcp_server_transaction` (Number)
- `dns_request` (String)
- `dns_response` (String)
- `dot11_protocol` (String)
- `dot11_protocol_capability` (String)
- `duid` (String)
- `firmware_version` (String)
- `frequency` (String)
- `health_score` (List of Object) (see [below for nested schema](#nestedobjatt--item--detail--health_score))
- `host_ip_v4` (String)
- `host_ip_v6` (List of String)
- `host_mac` (String)
- `host_name` (String)
- `host_os` (String)
- `host_type` (String)
- `host_version` (String)
- `hw_model` (String)
- `id` (String)
- `identifier` (String)
- `intel_capable` (String)
- `ios_capable` (String)
- `is_guest_upn_endpoint` (String)
- `issue_count` (Number)
- `l2_virtual_network` (String)
- `l3_virtual_network` (String)
- `last_updated` (Number)
- `latency_be` (Number)
- `latency_bg` (Number)
- `latency_video` (Number)
- `latency_voice` (Number)
- `link_speed` (Number)
- `link_threshold` (String)
- `location` (String)
- `max_roaming_duration` (String)
- `model_name` (String)
- `onboarding` (List of Object) (see [below for nested schema](#nestedobjatt--item--detail--onboarding))
- `onboarding_time` (Number)
- `port` (String)
- `port_description` (String)
- `power_type` (String)
- `private_mac` (String)
- `remote_end_duplex_mode` (String)
- `rssi` (String)
- `rssi_is_include` (String)
- `rssi_threshold` (String)
- `rx_bytes` (String)
- `rx_link_error` (Number)
- `rx_rate` (Number)
- `rx_retry_pct` (String)
- `sales_code` (String)
- `session_duration` (String)
- `sgt` (String)
- `slot_id` (Number)
- `snr` (String)
- `snr_is_include` (String)
- `snr_threshold` (String)
- `ssid` (String)
- `sub_type` (String)
- `tracked` (String)
- `trust_details` (String)
- `trust_score` (String)
- `tx_bytes` (String)
- `tx_link_error` (Number)
- `tx_rate` (Number)
- `upn_id` (String)
- `upn_name` (String)
- `upn_owner` (String)
- `usage` (Number)
- `user_id` (String)
- `version_time` (Number)
- `vlan_id` (Number)
- `vnid` (Number)
- `wlc_name` (String)
- `wlc_uuid` (String)

<a id="nestedobjatt--item--detail--connected_device"></a>
### Nested Schema for `item.detail.connected_device`

Read-Only:

- `band` (String)
- `id` (String)
- `ip_address` (String)
- `mac` (String)
- `mgmt_ip` (String)
- `mode` (String)
- `name` (String)
- `type` (String)


<a id="nestedobjatt--item--detail--health_score"></a>
### Nested Schema for `item.detail.health_score`

Read-Only:

- `health_type` (String)
- `reason` (String)
- `score` (Number)


<a id="nestedobjatt--item--detail--onboarding"></a>
### Nested Schema for `item.detail.onboarding`

Read-Only:

- `aaa_rootcause_list` (List of String)
- `aaa_server_ip` (String)
- `assoc_done_time` (Number)
- `assoc_rootcause_list` (List of String)
- `auth_done_time` (Number)
- `average_assoc_duration` (String)
- `average_auth_duration` (String)
- `average_dhcp_duration` (String)
- `average_run_duration` (String)
- `dhcp_done_time` (Number)
- `dhcp_rootcause_list` (List of String)
- `dhcp_server_ip` (String)
- `latest_root_cause_list` (List of String)
- `max_assoc_duration` (String)
- `max_auth_duration` (String)
- `max_dhcp_duration` (String)
- `max_run_duration` (String)
- `other_rootcause_list` (List of String)



<a id="nestedobjatt--item--topology"></a>
### Nested Schema for `item.topology`

Read-Only:

- `links` (List of Object) (see [below for nested schema](#nestedobjatt--item--topology--links))
- `nodes` (List of Object) (see [below for nested schema](#nestedobjatt--item--topology--nodes))

<a id="nestedobjatt--item--topology--links"></a>
### Nested Schema for `item.topology.links`

Read-Only:

- `ap_radio_admin_status` (String)
- `ap_radio_oper_status` (String)
- `id` (String)
- `interface_details` (List of Object) (see [below for nested schema](#nestedobjatt--item--topology--links--interface_details))
- `label` (List of String)
- `link_status` (String)
- `port_utilization` (Number)
- `source` (String)
- `source_admin_status` (String)
- `source_duplex_info` (String)
- `source_interface_name` (String)
- `source_link_status` (String)
- `source_port_mode` (String)
- `source_port_vla_n_info` (String)
- `target` (String)
- `target_admin_status` (String)
- `target_duplex_info` (String)
- `target_interface_name` (String)
- `target_link_status` (String)
- `target_port_mode` (String)
- `target_port_vla_n_info` (String)

<a id="nestedobjatt--item--topology--links--interface_details"></a>
### Nested Schema for `item.topology.links.target_port_vla_n_info`

Read-Only:

- `admin_status` (String)
- `client_mac_address` (String)
- `connected_device_int_name` (String)
- `duplex` (String)
- `port_mode` (String)



<a id="nestedobjatt--item--topology--nodes"></a>
### Nested Schema for `item.topology.nodes`

Read-Only:

- `clients` (Number)
- `connected_device` (String)
- `count` (Number)
- `description` (String)
- `device_type` (String)
- `fabric_group` (String)
- `fabric_role` (List of String)
- `family` (String)
- `health_score` (Number)
- `id` (String)
- `ip` (String)
- `ipv6` (List of String)
- `level` (Number)
- `name` (String)
- `node_type` (String)
- `platform_id` (String)
- `radio_frequency` (String)
- `role` (String)
- `software_version` (String)
- `stack_type` (String)
- `user_id` (String)
