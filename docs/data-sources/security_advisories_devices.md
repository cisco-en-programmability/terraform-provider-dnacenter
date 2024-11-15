---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_security_advisories_devices Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Security Advisories.
  Retrieves list of devices for an advisory
---

# dnacenter_security_advisories_devices (Data Source)

It performs read operation on Security Advisories.

- Retrieves list of devices for an advisory

## Example Usage

```terraform
data "dnacenter_security_advisories_devices" "example" {
  provider    = dnacenter
  advisory_id = "string"
}

output "dnacenter_security_advisories_devices_example" {
  value = data.dnacenter_security_advisories_devices.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `advisory_id` (String) advisoryId path parameter. Advisory ID

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `response` (List of String)
- `version` (String)
