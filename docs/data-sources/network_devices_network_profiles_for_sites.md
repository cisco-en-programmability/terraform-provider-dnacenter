---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network_devices_network_profiles_for_sites Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Site Design.
  Retrieves the list of network profiles for sites.Retrieves a network profile for sites by id.
---

# dnacenter_network_devices_network_profiles_for_sites (Data Source)

It performs read operation on Site Design.

- Retrieves the list of network profiles for sites.

- Retrieves a network profile for sites by id.

## Example Usage

```terraform
data "dnacenter_network_devices_network_profiles_for_sites" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
  order    = "string"
  sort_by  = "string"
  type     = "string"
}

output "dnacenter_network_devices_network_profiles_for_sites_example" {
  value = data.dnacenter_network_devices_network_profiles_for_sites.example.items
}

data "dnacenter_network_devices_network_profiles_for_sites" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_devices_network_profiles_for_sites_example" {
  value = data.dnacenter_network_devices_network_profiles_for_sites.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter. The *id* of the network profile, retrievable from *GET /intent/api/v1/networkProfilesForSites*
- `limit` (Number) limit query parameter. The number of records to show for this page.
- `offset` (Number) offset query parameter. The first record to show for this page; the first record is numbered 1.
- `order` (String) order query parameter. Whether ascending or descending order should be used to sort the response.
- `sort_by` (String) sortBy query parameter. A property within the response to sort by.
- `type` (String) type query parameter. Filter responses to only include profiles of a given type

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `id` (String)
- `name` (String)
- `type` (String)


<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String)
- `name` (String)
- `type` (String)
