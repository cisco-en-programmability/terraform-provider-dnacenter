---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_device_replacement_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Device Replacement.
  Get replacement devices count
---

# dnacenter_device_replacement_count (Data Source)

It performs read operation on Device Replacement.

- Get replacement devices count

## Example Usage

```terraform
data "dnacenter_device_replacement_count" "example" {
  provider           = dnacenter
  replacement_status = ["string"]
}

output "dnacenter_device_replacement_count_example" {
  value = data.dnacenter_device_replacement_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **replacement_status** (List of String) replacementStatus query parameter. Device Replacement status list[READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR]

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **response** (Number)
- **version** (String)

