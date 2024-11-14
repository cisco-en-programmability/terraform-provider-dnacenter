---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_sda_transit_networks_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Returns the count of transit networks that match the provided query parameters.
---

# dnacenter_sda_transit_networks_count (Data Source)

It performs read operation on SDA.

- Returns the count of transit networks that match the provided query parameters.

## Example Usage

```terraform
data "dnacenter_sda_transit_networks_count" "example" {
  provider = dnacenter
  type     = "string"
}

output "dnacenter_sda_transit_networks_count_example" {
  value = data.dnacenter_sda_transit_networks_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `type` (String) type query parameter. Type of the transit network. Allowed values are [IP_BASED_TRANSIT, SDA_LISP_PUB_SUB_TRANSIT, SDA_LISP_BGP_TRANSIT].

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)