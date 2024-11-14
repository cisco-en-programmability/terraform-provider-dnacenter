---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_maps_supported_access_points Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Sites.
  Gets the list of supported access point types as well as valid antenna pattern names that can be used for each.
---

# dnacenter_maps_supported_access_points (Data Source)

It performs read operation on Sites.

- Gets the list of supported access point types as well as valid antenna pattern names that can be used for each.

## Example Usage

```terraform
data "dnacenter_maps_supported_access_points" "example" {
  provider = dnacenter
}

output "dnacenter_maps_supported_access_points_example" {
  value = data.dnacenter_maps_supported_access_points.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `antenna_patterns` (List of Object) (see [below for nested schema](#nestedobjatt--items--antenna_patterns))
- `ap_type` (String)

<a id="nestedobjatt--items--antenna_patterns"></a>
### Nested Schema for `items.antenna_patterns`

Read-Only:

- `band` (String)
- `names` (List of String)