---
page_title: "dna_pnp_device_sync_vaact Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_device_sync_vaact data source allows you to sync a Cisco DNA Center PnP Virtual Account devices.
---

# Data Source dna_pnp_device_sync_vaact

The dna_pnp_device_sync_vaact data source allows you to sync a Cisco DNA Center PnP Virtual Account devices.

## Example Usage

```hcl
data "dna_pnp_device_sync_vaact" "result" {
  provider = dnacenter
  request {
    auto_sync_period = 0
    cco_user = ""
    expiry = 0
    last_sync = 0
    profile {
      address_fqdn = ""
      address_ip_v4 = ""
      cert = ""
      make_default = false
      name = ""
      port = 23
      profile_id = ""
      proxy = false
    }
    smart_account_id = ""
    sync_result {
      sync_list {
        device_sn_list = []
      }
    }
    sync_result_str = ""
    sync_start_time = 0
    sync_status = ""
    tenant_id = ""
    token = ""
    virtual_account_id = ""
  }
}
```

## Argument Reference

- `request` - (Optional) The request argument. See [request](#request) below for details.

### request

- `auto_sync_period` - (Optional) The request's auto sync period.
- `cco_user` - (Optional) The request's cco user.
- `expiry` - (Optional) The request's expiry.
- `last_sync` - (Optional) The request's last sync.
- `profile` - (Optional) The request's profile. See [profile](#profile) below for details.
- `smart_account_id` - (Optional) The request's smart account id.
- `sync_result` - (Optional) The request's sync result. See [sync_result](#sync_result) below for details.
- `sync_result_str` - (Optional) The request's sync result str.
- `sync_start_time` - (Optional) The request's sync start time.
- `sync_status` - (Optional) The request's sync status.
- `tenant_id` - (Optional) The request's tenant id.
- `token` - (Optional) The request's token.
- `virtual_account_id` - (Optional) The request's virtual account id.

#### profile

- `address_fqdn` - (Optional) The profile's address fqdn.
- `address_ip_v4` - (Optional) The profile's address ip v4.
- `cert` - (Optional) The profile's cert.
- `make_default` - (Optional) The profile's make default flag.
- `name` - (Optional) The profile's name.
- `port` - (Optional) The profile's port.
- `profile_id` - (Optional) The profile's profile id.
- `proxy` - (Optional) The profile's proxy flag.

#### sync_result

- `sync_list` - (Optional) The sync result's sync list. See below for details.

##### sync_list

- `device_sn_list` - (Optional) The sync element's device sn list.
- `sync_type` - (Optional) The sync element's sync type.

## Attributes Reference

The following attributes are exported.

- `item` - The config_preview's item. See [Item](#item) below for details.

### Item

- `auto_sync_period` - The response's auto sync period.
- `cco_user` - The response's cco user.
- `expiry` - The response's expiry.
- `last_sync` - The response's last sync.
- `profile` - The response's profile. See [profile](#profile-1) below for details.
- `smart_account_id` - The response's smart account id.
- `sync_result` - The response's sync result. See [sync_result](#sync_result-1) below for details.
- `sync_result_str` - The response's sync result str.
- `sync_start_time` - The response's sync start time.
- `sync_status` - The response's sync status.
- `tenant_id` - The response's tenant id.
- `token` - The response's token.
- `virtual_account_id` - The response's virtual account id.

#### profile

- `address_fqdn` - The profile's address fqdn.
- `address_ip_v4` - The profile's address ip v4.
- `cert` - The profile's cert.
- `make_default` - The profile's make default flag.
- `name` - The profile's name.
- `port` - The profile's port.
- `profile_id` - The profile's profile id.
- `proxy` - The profile's proxy flag.

#### sync_result

- `sync_list` - The sync result's sync list. See below for details.

##### sync_list

- `device_sn_list` - The sync element's device sn list.
- `sync_type` - The sync element's sync type.
