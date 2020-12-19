---
page_title: "dna_pnp_global_settings Resource - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_global_settings resource allows you to configure a DNACenter PnP's settings.
---

# Resource dna_pnp_global_settings

The dna_pnp_global_settings resource allows you to configure a DNACenter PnP's settings.

## Example Usage

```hcl
resource "dna_pnp_global_settings" "response" {
  provider = dnacenter
  item {
    version = 1
    aaa_credentials {
      username = ""
      password = ""
    }
    task_time_outs {
      config_time_out         = 10
      image_download_time_out = 120
      general_time_out        = 20
    }
    sava_mapping_list {
    }
    accept_eula = true
    default_profile {
      ip_addresses = [
        "",
        "172.168.196.2",
        "10.121.1.5"
      ]
      fqdn_addresses = []
      port           = 443
      # cert arg value was simplified
      cert           = "-----BEGIN CERTIFICATE-----\....\n-----END CERTIFICATE-----"
      proxy          = false
    }
  }
}
```

## Argument Reference

- `item` - (Required) Item in the DNACenter global settings. See [Settings item](#settings-item) below for details.

### Settings item

- `id` - (Optional) The settings' id.
- `aaa_credentials` - (Optional) The settings' aaa credentials. See [aaa credentials](#aaa_credentials) below for details.
- `accept_eula` - (Optional) The settings' accept eula flag.
- `default_profile` - (Optional) The settings' default profile. See [default profile](#default_profile) below for details.
- `sava_mapping_list` - (Optional) The settings' sava mapping list. See [sava mapping list](#sava_mapping_list) below for details.
- `task_time_outs` - (Optional) The settings' task time outs. See [task time outs](#task_time_outs) below for details.
- `tenant_id` - (Optional) The settings' tenant id.
- `version` - (Optional) The settings' version.

### aaa_credentials

- `password` - (Optional) The aaa's password.
- `username` - (Optional) The aaa's username.

### default_profile

- `cert` - (Optional) The profile's cert.
- `fqdn_addresses` - (Optional) The profile's FQDN addresses.
- `ip_addresses` - (Optional) The profile's IP addresses.
- `port` - (Optional) The profile's port.
- `proxy` - (Optional) The profile's proxy flag.

### sava_mapping_list

- `auto_sync_period` - (Optional) The mapping list's auto sync period.
- `cco_user` - (Optional) The mapping list's cco user.
- `expiry` - (Optional) The mapping list's expiry.
- `last_sync` - (Optional) The mapping list's last sync.
- `profile` - (Optional) The mapping list's profile. See [profile](#profile) below for details.
- `smart_account_id` - (Optional) The mapping list's smart account id.
- `sync_result` - (Optional) The mapping list's sync result. See [sync_result](#sync_result) below for details.
- `sync_result_str` - (Optional) The mapping list's sync result str.
- `sync_start_time` - (Optional) The mapping list's sync start time.
- `sync_status` - (Optional) The mapping list's sync status.
- `tenant_id` - (Optional) The mapping list's tenant id.
- `token` - (Optional) The mapping list's token.
- `virtual_account_id` - (Optional) The mapping list's virtual account id.

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

### task_time_outs

- `config_time_out` - (Optional) The task's config time out.
- `general_time_out` - (Optional) The task's general time out.
- `image_download_time_out` - (Optional) The task's image download time out.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The PnP workflow's updated time with format RFC850.
