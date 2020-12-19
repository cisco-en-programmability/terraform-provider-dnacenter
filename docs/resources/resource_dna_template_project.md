---
page_title: "dna_template_project Resource - terraform-provider-dnacenter"
subcategory: "Configuration Templates"
description: |-
  The dna_template_project resource allows you to configure a DNACenter Template project.
---

# Resource dna_template_project

The dna_template_project resource allows you to configure a DNACenter Template project.

## Example Usage

```hcl
resource "dna_template_project" "project_1" {
  provider = dnacenter
  item {
    is_deletable = true
    name         = "Cloud Test Template 2"
  }
}
```

## Argument Reference

- `item` - (Required) Item in the DNACenter template project. See [Project item](#project-item) below for details.

### Project item

- `is_deletable` - (Optional) The project's is deletable flag. Once set it cannot be changed. If its value is false the resource cannot be deleted.
- `name` - (Required) The project's name.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The Template project's updated time with format RFC850.
- `item` - Item in the DNACenter template project. See [Project item](#project-item-1) below for details.

### Project item

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
