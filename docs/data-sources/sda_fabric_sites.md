---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_sda_fabric_sites Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Returns a list of fabric sites that match the provided query parameters.
---

# dnacenter_sda_fabric_sites (Data Source)

It performs read operation on SDA.

- Returns a list of fabric sites that match the provided query parameters.

## Example Usage

```terraform
data "dnacenter_sda_fabric_sites" "example" {
  provider = dnacenter
  id       = "string"
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "dnacenter_sda_fabric_sites_example" {
  value = data.dnacenter_sda_fabric_sites.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id query parameter. ID of the fabric site.
- `limit` (Number) limit query parameter. Maximum number of records to return.
- `offset` (Number) offset query parameter. Starting record for pagination.
- `site_id` (String) siteId query parameter. ID of the network hierarchy associated with the fabric site.

### Read-Only

- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `authentication_profile_name` (String)
- `id` (String)
- `is_pub_sub_enabled` (String)
- `site_id` (String)
