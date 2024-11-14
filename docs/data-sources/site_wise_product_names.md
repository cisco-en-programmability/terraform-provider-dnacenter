---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_site_wise_product_names Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Software Image Management (SWIM).
  Provides network device product names for a site. The default value of siteId is global. The response will include
  the network device count and image summary.
---

# dnacenter_site_wise_product_names (Data Source)

It performs read operation on Software Image Management (SWIM).

- Provides network device product names for a site. The default value of *siteId* is global. The response will include
the network device count and image summary.

## Example Usage

```terraform
data "dnacenter_site_wise_product_names" "example" {
  provider     = dnacenter
  limit        = 1
  offset       = 1
  product_name = "string"
  site_id      = "string"
}

output "dnacenter_site_wise_product_names_example" {
  value = data.dnacenter_site_wise_product_names.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `limit` (Number) limit query parameter. The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively
- `offset` (Number) offset query parameter. The first record to show for this page; the first record is numbered 1. The minimum value is 1
- `product_name` (String) productName query parameter. Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
- `site_id` (String) siteId query parameter. Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for siteId

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `id` (String)
- `image_summary` (List of Object) (see [below for nested schema](#nestedobjatt--item--image_summary))
- `network_device_count` (Number)
- `product_name` (String)
- `product_name_ordinal` (Number)
- `supervisor_product_name` (String)
- `supervisor_product_name_ordinal` (Number)

<a id="nestedobjatt--item--image_summary"></a>
### Nested Schema for `item.image_summary`

Read-Only:

- `golden_image_count` (Number)
- `installed_image_advisor_count` (Number)
- `installed_image_count` (Number)