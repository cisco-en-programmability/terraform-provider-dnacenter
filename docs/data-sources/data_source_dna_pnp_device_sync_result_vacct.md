---
page_title: "dna_pnp_device_sync_result_vacct Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_device_sync_result_vacct data source retrieves the sync result for DNACenter Virtual Account.
---

# Data Source dna_pnp_device_sync_result_vacct

The dna_pnp_device_sync_result_vacct data source retrieves the sync result for DNACenter Virtual Account.

## Example Usage

```hcl
data "dna_pnp_device_sync_result_vacct" "result" {
  provider = dnacenter
  # name     = var.name
  # domain   = var.domain
}
```

## Argument Reference

- `name` - (Optional) Virtual Account Name param.
- `domain` - (Optional) Smart Account Domain param.

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

- `auto_sync_period` - The item's auto sync period.
- `cco_user` - The item's cco user.
- `expiry` - The item's expiry.
- `last_sync` - The item's last sync.
- `profile` - The item's profile. See [profile](#profile) below for details.
- `smart_account_id` - The item's smart account id.
- `sync_result` - The item's sync result. See [sync_result](#sync_result) below for details.
- `sync_result_str` - The item's sync result str.
- `sync_start_time` - The item's sync start time.
- `sync_status` - The item's sync status.
- `tenant_id` - The item's tenant id.
- `token` - The item's token.
- `virtual_account_id` - The item's virtual account id.

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
