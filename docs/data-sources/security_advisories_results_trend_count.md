---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_security_advisories_results_trend_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Compliance.
  Get count of security advisories results trend over time.
---

# dnacenter_security_advisories_results_trend_count (Data Source)

It performs read operation on Compliance.

- Get count of security advisories results trend over time.

## Example Usage

```terraform
data "dnacenter_security_advisories_results_trend_count" "example" {
  provider  = dnacenter
  scan_time = 1.0
}

output "dnacenter_security_advisories_results_trend_count_example" {
  value = data.dnacenter_security_advisories_results_trend_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `scan_time` (Number) scanTime query parameter. Return advisories trend with scanTime greater than this scanTime

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)
