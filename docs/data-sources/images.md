---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_images Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Software Image Management (SWIM).
  A list of available images for the specified site is provided. The default value of the site is set to global. The
  list includes images that have been imported onto the disk, as well as the latest and suggested images from Cisco.com.
---

# dnacenter_images (Data Source)

It performs read operation on Software Image Management (SWIM).

- A list of available images for the specified site is provided. The default value of the site is set to global. The
list includes images that have been imported onto the disk, as well as the latest and suggested images from Cisco.com.

## Example Usage

```terraform
data "dnacenter_images" "example" {
  provider                        = dnacenter
  golden                          = "false"
  has_addon_images                = "false"
  imported                        = "false"
  integrity                       = "string"
  is_addon_images                 = "false"
  limit                           = 1
  name                            = "string"
  offset                          = 1
  product_name_ordinal            = 1.0
  site_id                         = "string"
  supervisor_product_name_ordinal = 1.0
  version                         = "string"
}

output "dnacenter_images_example" {
  value = data.dnacenter_images.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `golden` (Boolean) golden query parameter. When set to *true*, it will retrieve the images marked as tagged golden. When set to *false*, it will retrieve the images marked as not tagged golden.
- `has_addon_images` (Boolean) hasAddonImages query parameter. When set to *true*, it will retrieve the images which have add-on images. When set to *false*, it will retrieve the images which do not have add-on images.
- `imported` (Boolean) imported query parameter. When the value is set to *true*, it will include physically imported images. Conversely, when the value is set to *false*, it will include image records from the cloud. The identifier for cloud images can be utilized to download images from Cisco.com to the disk.
- `integrity` (String) integrity query parameter. Filter with verified images using Integrity Verification Available values: UNKNOWN, VERIFIED
- `is_addon_images` (Boolean) isAddonImages query parameter. When set to *true*, it will retrieve the images that an add-on image.  When set to *false*, it will retrieve the images that are not add-on images
- `limit` (Number) limit query parameter. The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively.
- `name` (String) name query parameter. Filter with software image or add-on name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
- `offset` (Number) offset query parameter. The first record to show for this page; the first record is numbered 1. The minimum value is 1.
- `product_name_ordinal` (Number) productNameOrdinal query parameter. The product name ordinal is a unique value for each network device product. The productNameOrdinal can be obtained from the response of API */dna/intent/api/v1/siteWiseProductNames*
- `site_id` (String) siteId query parameter. Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for *siteId*
- `supervisor_product_name_ordinal` (Number) supervisorProductNameOrdinal query parameter. The supervisor engine module ordinal is a unique value for each supervisor module. The *supervisorProductNameOrdinal* can be obtained from the response of API */dna/intent/api/v1/siteWiseProductNames*
- `version` (String) version query parameter. Filter with image version. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `cisco_latest` (String)
- `golden_tagging_details` (List of Object) (see [below for nested schema](#nestedobjatt--items--golden_tagging_details))
- `has_addon_images` (String)
- `id` (String)
- `image_type` (String)
- `imported` (String)
- `integrity_status` (String)
- `is_addon_image` (String)
- `is_golden_tagged` (String)
- `name` (String)
- `product_names` (List of Object) (see [below for nested schema](#nestedobjatt--items--product_names))
- `recommended` (String)
- `version` (String)

<a id="nestedobjatt--items--golden_tagging_details"></a>
### Nested Schema for `items.golden_tagging_details`

Read-Only:

- `device_roles` (List of String)
- `device_tags` (List of String)
- `inherited_site_id` (String)
- `inherited_site_name` (String)


<a id="nestedobjatt--items--product_names"></a>
### Nested Schema for `items.product_names`

Read-Only:

- `id` (String)
- `product_name` (String)
- `product_name_ordinal` (Number)
- `supervisor_product_name` (String)
- `supervisor_product_name_ordinal` (Number)