---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network_device_user_defined_field Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Gets existing global User Defined Fields. If no input is given, it fetches ALL the Global UDFs. Filter/search is
  supported by UDF Id(s) or UDF name(s) or both.
---

# dnacenter_network_device_user_defined_field (Data Source)

It performs read operation on Devices.

- Gets existing global User Defined Fields. If no input is given, it fetches ALL the Global UDFs. Filter/search is
supported by UDF Id(s) or UDF name(s) or both.

## Example Usage

```terraform
data "dnacenter_network_device_user_defined_field" "example" {
  provider = dnacenter
  id       = "string"
  name     = "string"
}

output "dnacenter_network_device_user_defined_field_example" {
  value = data.dnacenter_network_device_user_defined_field.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id query parameter. Comma-seperated id(s) used for search/filtering
- `name` (String) name query parameter. Comma-seperated name(s) used for search/filtering

### Read-Only

- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `description` (String)
- `id` (String)
- `name` (String)
