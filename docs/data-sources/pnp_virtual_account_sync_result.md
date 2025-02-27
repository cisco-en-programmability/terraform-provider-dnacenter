---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_pnp_virtual_account_sync_result Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Device Onboarding (PnP).
  Returns the summary of devices synced from the given smart account & virtual account with PnP (Deprecated)
---

# dnacenter_pnp_virtual_account_sync_result (Data Source)

It performs read operation on Device Onboarding (PnP).

- Returns the summary of devices synced from the given smart account & virtual account with PnP (Deprecated)

## Example Usage

```terraform
data "dnacenter_pnp_virtual_account_sync_result" "example" {
  provider = dnacenter
  domain   = "string"
  name     = "string"
}

output "dnacenter_pnp_virtual_account_sync_result_example" {
  value = data.dnacenter_pnp_virtual_account_sync_result.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain` (String) domain path parameter. Smart Account Domain
- `name` (String) name path parameter. Virtual Account Name

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `auto_sync_period` (Number)
- `cco_user` (String)
- `expiry` (Number)
- `last_sync` (Number)
- `profile` (List of Object) (see [below for nested schema](#nestedobjatt--item--profile))
- `smart_account_id` (String)
- `sync_result` (List of Object) (see [below for nested schema](#nestedobjatt--item--sync_result))
- `sync_result_str` (String)
- `sync_start_time` (Number)
- `sync_status` (String)
- `tenant_id` (String)
- `token` (String)
- `virtual_account_id` (String)

<a id="nestedobjatt--item--profile"></a>
### Nested Schema for `item.profile`

Read-Only:

- `address_fqdn` (String)
- `address_ip_v4` (String)
- `cert` (String)
- `make_default` (String)
- `name` (String)
- `port` (Number)
- `profile_id` (String)
- `proxy` (String)


<a id="nestedobjatt--item--sync_result"></a>
### Nested Schema for `item.sync_result`

Read-Only:

- `sync_list` (List of Object) (see [below for nested schema](#nestedobjatt--item--sync_result--sync_list))
- `sync_msg` (String)

<a id="nestedobjatt--item--sync_result--sync_list"></a>
### Nested Schema for `item.sync_result.sync_list`

Read-Only:

- `device_sn_list` (List of String)
- `sync_type` (String)
