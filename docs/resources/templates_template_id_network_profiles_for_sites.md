---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_templates_template_id_network_profiles_for_sites Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages create and read operations on Configuration Templates.
  Attaches a network profile to a Day-N CLI template by passing the profile ID and template ID.
---

# dnacenter_templates_template_id_network_profiles_for_sites (Resource)

It manages create and read operations on Configuration Templates.

- Attaches a network profile to a Day-N CLI template by passing the profile ID and template ID.

## Example Usage

```terraform
resource "dnacenter_templates_template_id_network_profiles_for_sites" "example" {
  provider = dnacenter

  parameters {

    profile_id  = "string"
    template_id = "string"
  }
}

output "dnacenter_templates_template_id_network_profiles_for_sites_example" {
  value = dnacenter_templates_template_id_network_profiles_for_sites.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (String)
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `template_id` (String) templateId path parameter. The id of the template, retrievable from GET /intent/api/v1/templates

Optional:

- `profile_id` (String) The id of the network profile, retrievable from /intent/api/v1/networkProfilesForSites

## Import

Import is supported using the following syntax:

```shell
terraform import dnacenter_templates_template_id_network_profiles_for_sites.example "template_id:=string"
```
