---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_sda_fabric_site Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Get Site info from SDA Fabric
---

# dnacenter_sda_fabric_site (Data Source)

It performs read operation on SDA.

- Get Site info from SDA Fabric

## Example Usage

```terraform
data "dnacenter_sda_fabric_site" "example" {
  provider            = dnacenter
  site_name_hierarchy = "string"
}

output "dnacenter_sda_fabric_site_example" {
  value = data.dnacenter_sda_fabric_site.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **site_name_hierarchy** (String) siteNameHierarchy query parameter. Site Name Hierarchy

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **description** (String)
- **execution_status_url** (String)
- **status** (String)

