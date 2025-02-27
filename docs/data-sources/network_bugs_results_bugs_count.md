---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network_bugs_results_bugs_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Compliance.
  Get count of network bugs
---

# dnacenter_network_bugs_results_bugs_count (Data Source)

It performs read operation on Compliance.

- Get count of network bugs

## Example Usage

```terraform
data "dnacenter_network_bugs_results_bugs_count" "example" {
  provider     = dnacenter
  device_count = 1.0
  id           = "string"
  severity     = "string"
}

output "dnacenter_network_bugs_results_bugs_count_example" {
  value = data.dnacenter_network_bugs_results_bugs_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device_count` (Number) deviceCount query parameter. Return network bugs with deviceCount greater than this deviceCount
- `id` (String) id query parameter. Id of the network bug
- `severity` (String) severity query parameter. Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)
