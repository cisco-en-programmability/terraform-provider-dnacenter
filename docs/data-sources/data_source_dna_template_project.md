---
page_title: "dna_template_project Data Source - terraform-provider-dnacenter"
subcategory: "Configuration Templates"
description: |-
  The dna_template_project data source allows you to retrieve information about a particular Cisco DNA Center template project.
---

# Data Source dna_template_project

The dna_template_project data source allows you to retrieve information about a particular Cisco DNA Center template project.

## Example Usage

```hcl
data "dna_template_project" "response" {
  provider = dnacenter
  name     = "Cloud DayN Templates"
}
```

## Argument Reference

- `name` - (Optional) The project name param.

## Attributes Reference

The following attributes are exported.

- `items` - The items response. See [Items](#items) below for details.

### Items

- `is_deletable` - The project's is deletable flag. If its value is false the resource cannot be deleted.
- `name` - The project's name.
- `id` - The project's id.
- `templates` - The project's templates. See [templates](#templates) below for details.

### templates

- `name` - The template's name.
- `composite` - The template's composite.
- `language` - The template's language.
- `id` - The template's id.
- `custom_params_order` - The template's custom params order.
- `last_update_time` - The template's last update time.
- `latest_version_time` - The template's latest version time.
- `project_associated` - The template's project associated.
- `document_database` - The template's document database.
