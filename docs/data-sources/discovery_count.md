---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_discovery_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Discovery.
  Returns the count of all available discovery jobs
---

# dnacenter_discovery_count (Data Source)

It performs read operation on Discovery.

- Returns the count of all available discovery jobs

## Example Usage

```terraform
data "dnacenter_discovery_count" "example" {
  provider = dnacenter
}

output "dnacenter_discovery_count_example" {
  value = data.dnacenter_discovery_count.example.item
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

- `response` (Number)
- `version` (String)
