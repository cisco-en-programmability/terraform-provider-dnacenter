---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_app_policy Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Application Policy.
  Get all existing application policies
---

# dnacenter_app_policy (Data Source)

It performs read operation on Application Policy.

- Get all existing application policies

## Example Usage

```terraform
data "dnacenter_app_policy" "example" {
  provider     = dnacenter
  policy_scope = "string"
}

output "dnacenter_app_policy_example" {
  value = data.dnacenter_app_policy.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `policy_scope` (String) policyScope query parameter. policy scope name

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `advanced_policy_scope` (List of Object) (see [below for nested schema](#nestedobjatt--items--advanced_policy_scope))
- `cfs_change_info` (List of String)
- `consumer` (List of Object) (see [below for nested schema](#nestedobjatt--items--consumer))
- `contract_list` (List of String)
- `create_time` (Number)
- `custom_provisions` (List of String)
- `delete_policy_status` (String)
- `deployed` (String)
- `display_name` (String)
- `exclusive_contract` (List of Object) (see [below for nested schema](#nestedobjatt--items--exclusive_contract))
- `id` (String)
- `identity_source` (List of Object) (see [below for nested schema](#nestedobjatt--items--identity_source))
- `instance_created_on` (Number)
- `instance_id` (Number)
- `instance_updated_on` (Number)
- `instance_version` (Number)
- `internal` (String)
- `is_deleted` (String)
- `is_enabled` (String)
- `is_scope_stale` (String)
- `is_seeded` (String)
- `is_stale` (String)
- `ise_reserved` (String)
- `last_update_time` (Number)
- `name` (String)
- `namespace` (String)
- `policy_scope` (String)
- `policy_status` (String)
- `priority` (Number)
- `producer` (List of Object) (see [below for nested schema](#nestedobjatt--items--producer))
- `provisioning_state` (String)
- `pushed` (String)
- `qualifier` (String)
- `resource_version` (Number)
- `target_id_list` (List of String)
- `type` (String)

<a id="nestedobjatt--items--advanced_policy_scope"></a>
### Nested Schema for `items.advanced_policy_scope`

Read-Only:

- `advanced_policy_scope_element` (List of Object) (see [below for nested schema](#nestedobjatt--items--advanced_policy_scope--advanced_policy_scope_element))
- `display_name` (String)
- `id` (String)
- `instance_created_on` (Number)
- `instance_id` (Number)
- `instance_updated_on` (Number)
- `instance_version` (Number)
- `name` (String)

<a id="nestedobjatt--items--advanced_policy_scope--advanced_policy_scope_element"></a>
### Nested Schema for `items.advanced_policy_scope.advanced_policy_scope_element`

Read-Only:

- `display_name` (String)
- `group_id` (List of String)
- `id` (String)
- `instance_created_on` (Number)
- `instance_id` (Number)
- `instance_updated_on` (Number)
- `instance_version` (Number)
- `ssid` (List of String)



<a id="nestedobjatt--items--consumer"></a>
### Nested Schema for `items.consumer`

Read-Only:

- `display_name` (String)
- `id` (String)
- `instance_created_on` (Number)
- `instance_id` (Number)
- `instance_updated_on` (Number)
- `instance_version` (Number)
- `scalable_group` (List of Object) (see [below for nested schema](#nestedobjatt--items--consumer--scalable_group))

<a id="nestedobjatt--items--consumer--scalable_group"></a>
### Nested Schema for `items.consumer.scalable_group`

Read-Only:

- `id_ref` (String)



<a id="nestedobjatt--items--exclusive_contract"></a>
### Nested Schema for `items.exclusive_contract`

Read-Only:

- `clause` (List of Object) (see [below for nested schema](#nestedobjatt--items--exclusive_contract--clause))
- `display_name` (String)
- `id` (String)
- `instance_created_on` (Number)
- `instance_id` (Number)
- `instance_updated_on` (Number)
- `instance_version` (Number)

<a id="nestedobjatt--items--exclusive_contract--clause"></a>
### Nested Schema for `items.exclusive_contract.clause`

Read-Only:

- `device_removal_behavior` (String)
- `display_name` (String)
- `host_tracking_enabled` (String)
- `id` (String)
- `instance_created_on` (Number)
- `instance_id` (Number)
- `instance_updated_on` (Number)
- `instance_version` (Number)
- `priority` (Number)
- `relevance_level` (String)
- `type` (String)



<a id="nestedobjatt--items--identity_source"></a>
### Nested Schema for `items.identity_source`

Read-Only:

- `display_name` (String)
- `id` (String)
- `instance_created_on` (Number)
- `instance_id` (Number)
- `instance_updated_on` (Number)
- `instance_version` (Number)
- `state` (String)
- `type` (String)


<a id="nestedobjatt--items--producer"></a>
### Nested Schema for `items.producer`

Read-Only:

- `display_name` (String)
- `id` (String)
- `instance_created_on` (Number)
- `instance_id` (Number)
- `instance_updated_on` (Number)
- `instance_version` (Number)
- `scalable_group` (List of Object) (see [below for nested schema](#nestedobjatt--items--producer--scalable_group))

<a id="nestedobjatt--items--producer--scalable_group"></a>
### Nested Schema for `items.producer.scalable_group`

Read-Only:

- `id_ref` (String)
