---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_analytics_anc_policies Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on AI Endpoint Analytics.
  Fetches the list of ANC policies available in ISE.
---

# dnacenter_analytics_anc_policies (Data Source)

It performs read operation on AI Endpoint Analytics.

- Fetches the list of ANC policies available in ISE.

## Example Usage

```terraform
data "dnacenter_analytics_anc_policies" "example" {
  provider = dnacenter
}

output "dnacenter_analytics_anc_policies_example" {
  value = data.dnacenter_analytics_anc_policies.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `name` (String)