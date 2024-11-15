---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_wireless_profiles Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Wireless.
  This data source allows the user to get all Wireless Network ProfilesThis data source allows the user to get a Wireless Network Profile by ID
---

# dnacenter_wireless_profiles (Data Source)

It performs read operation on Wireless.

- This data source allows the user to get all Wireless Network Profiles

- This data source allows the user to get a Wireless Network Profile by ID

## Example Usage

```terraform
data "dnacenter_wireless_profiles" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
}

output "dnacenter_wireless_profiles_example" {
  value = data.dnacenter_wireless_profiles.example.items
}

data "dnacenter_wireless_profiles" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_wireless_profiles_example" {
  value = data.dnacenter_wireless_profiles.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter. Wireless Profile Id
- `limit` (Number) limit query parameter.
- `offset` (Number) offset query parameter.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `id` (String)
- `ssid_details` (List of Object) (see [below for nested schema](#nestedobjatt--item--ssid_details))
- `wireless_profile_name` (String)

<a id="nestedobjatt--item--ssid_details"></a>
### Nested Schema for `item.ssid_details`

Read-Only:

- `dot11be_profile_id` (String)
- `enable_fabric` (String)
- `flex_connect` (List of Object) (see [below for nested schema](#nestedobjatt--item--ssid_details--flex_connect))
- `interface_name` (String)
- `policy_profile_name` (String)
- `ssid_name` (String)
- `wlan_profile_name` (String)

<a id="nestedobjatt--item--ssid_details--flex_connect"></a>
### Nested Schema for `item.ssid_details.flex_connect`

Read-Only:

- `enable_flex_connect` (String)
- `local_to_vlan` (Number)




<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String)
- `ssid_details` (List of Object) (see [below for nested schema](#nestedobjatt--items--ssid_details))
- `wireless_profile_name` (String)

<a id="nestedobjatt--items--ssid_details"></a>
### Nested Schema for `items.ssid_details`

Read-Only:

- `dot11be_profile_id` (String)
- `enable_fabric` (String)
- `flex_connect` (List of Object) (see [below for nested schema](#nestedobjatt--items--ssid_details--flex_connect))
- `interface_name` (String)
- `policy_profile_name` (String)
- `ssid_name` (String)
- `wlan_profile_name` (String)

<a id="nestedobjatt--items--ssid_details--flex_connect"></a>
### Nested Schema for `items.ssid_details.flex_connect`

Read-Only:

- `enable_flex_connect` (String)
- `local_to_vlan` (Number)
