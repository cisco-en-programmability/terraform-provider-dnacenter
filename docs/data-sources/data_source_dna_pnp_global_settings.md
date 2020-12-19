---
page_title: "dna_pnp_global_settings Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_pnp_global_settings data source allows you to retrieve information about a particular DNACenter PnP global settings.
---

# Data Source dna_pnp_global_settings

The dna_pnp_global_settings data source allows you to retrieve information about a particular DNACenter PnP global settings.

## Example Usage

```hcl
data "dna_pnp_global_settings" "result" {
  provider = dnacenter
}
```

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

- `id` - The settings' id.
- `aaa_credentials` - The settings' aaa credentials. See [aaa credentials](#aaa_credentials) below for details.
- `accept_eula` - The settings' accept eula flag.
- `default_profile` - The settings' default profile. See [default profile](#default_profile) below for details.
- `sava_mapping_list` - The settings' sava mapping list. See [sava mapping list](#sava_mapping_list) below for details.
- `task_time_outs` - The settings' task time outs. See [task time outs](#task_time_outs) below for details.
- `tenant_id` - The settings' tenant id.
- `version` - The settings' version.

#### aaa_credentials

- `password` - The aaa's password.
- `username` - The aaa's username.

#### default_profile

- `cert` - The profile's cert.
- `fqdn_addresses` - The profile's FQDN addresses.
- `ip_addresses` - The profile's IP addresses.
- `port` - The profile's port.
- `proxy` - The profile's proxy flag.

#### sava_mapping_list

- `auto_sync_period` - The mapping list's auto sync period.
- `cco_user` - The mapping list's cco user.
- `expiry` - The mapping list's expiry.
- `last_sync` - The mapping list's last sync.
- `profile` - The mapping list's profile. See [profile](#profile) below for details.
- `smart_account_id` - The mapping list's smart account id.
- `sync_result` - The mapping list's sync result. See [sync_result](#sync_result) below for details.
- `sync_result_str` - The mapping list's sync result str.
- `sync_start_time` - The mapping list's sync start time.
- `sync_status` - The mapping list's sync status.
- `tenant_id` - The mapping list's tenant id.
- `token` - The mapping list's token.
- `virtual_account_id` - The mapping list's virtual account id.

##### profile

- `address_fqdn` - The profile's address fqdn.
- `address_ip_v4` - The profile's address ip v4.
- `cert` - The profile's cert.
- `make_default` - The profile's make default flag.
- `name` - The profile's name.
- `port` - The profile's port.
- `profile_id` - The profile's profile id.
- `proxy` - The profile's proxy flag.

##### sync_result

- `sync_list` - The sync result's sync list. See below for details.

###### sync_list

- `device_sn_list` - The sync element's device sn list.
- `sync_type` - The sync element's sync type.

#### task_time_outs

- `config_time_out` - The task's config time out.
- `general_time_out` - The task's general time out.
- `image_download_time_out` - The task's image download time out.
