---
page_title: "dna_template Data Source - terraform-provider-dnacenter"
subcategory: "Configuration Templates"
description: |-
  The dna_template data source allows you to retrieve information about a particular DNACenter template.
---

# Data Source dna_template

The dna_template data source allows you to retrieve information about a particular DNACenter template.

## Example Usage

```hcl
data "dna_template" "response" {
  provider   = dnacenter
  project_id = data.dna_template_project.response.items.0.id
  # product_series = "Cisco Cloud Services Router 1000V Series"
  # product_family = "Routers"
  # software_type  = "IOS"
}
```

## Argument Reference

- `project_id` - (Optional) The project_id param.
- `software_type` - (Optional) The software_type param.
- `software_version` - (Optional) The software_version param.
- `product_family` - (Optional) The product_family param.
- `product_series` - (Optional) The product_series param.
- `product_type` - (Optional) The product_type param.
- `filter_conflicting_templates` - (Optional) The filter_conflicting_templates param.

## Attributes Reference

The following attributes are exported.

- `items` - The items response. See [Items](#items) below for details.

### Items

- `composite` - The template's composite.
- `name` - The template's name.
- `project_id` - The template's project id.
- `project_name` - The template's project name.
- `template_id` - The template's id.
- `versions_info` - The template's versions info. See [versions_info](#versions_info) below for details.

#### versions_info

- `description` - The version's description.
- `id` - The version's id.
- `version_time` - The version's time.
- `author` - The version's author.
- `version` - The version's version.
- `version_comment` - The version's comment.
