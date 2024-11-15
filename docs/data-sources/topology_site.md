---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_topology_site Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Topology.
  Returns site topology
---

# dnacenter_topology_site (Data Source)

It performs read operation on Topology.

- Returns site topology

## Example Usage

```terraform
data "dnacenter_topology_site" "example" {
  provider = dnacenter
}

output "dnacenter_topology_site_example" {
  value = data.dnacenter_topology_site.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `sites` (List of Object) (see [below for nested schema](#nestedobjatt--item--sites))

<a id="nestedobjatt--item--sites"></a>
### Nested Schema for `item.sites`

Read-Only:

- `display_name` (String)
- `group_name_hierarchy` (String)
- `id` (String)
- `latitude` (String)
- `location_address` (String)
- `location_country` (String)
- `location_type` (String)
- `longitude` (String)
- `name` (String)
- `parent_id` (String)
