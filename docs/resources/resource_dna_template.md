---
page_title: "dna_template Resource - terraform-provider-dnacenter"
subcategory: "Configuration Templates"
description: |-
  The dna_template resource allows you to configure a DNACenter template.
---

# Resource dna_template

The dna_template resource allows you to configure a DNACenter template.

## Example Usage

```hcl
resource "dna_template" "template_1" {
  provider   = dnacenter
  depends_on = [dna_template_project.project_1]
  item {
    project_id = dna_template_project.project_1.item.0.id
    name       = "DMVPN Spoke for Branch Router - System Default for Test Project"
    device_types {
      product_family = "Routers"
    }
    software_type = "IOS-XE"
  }
}
```

## Argument Reference

- `item` - (Required) Item in the DNACenter template. See [Template item](#template-item) below for details.

### Template item

- `project_id` - (Required) The template's project id.
- `author` - (Optional) The template's author.
- `composite` - (Optional) The template's composite flag.
- `create_time` - (Optional) The template's create time.
- `description` - (Optional) The template's description.
- `failure_policy` - (Optional) The template's failure policy.
- `id` - (Optional) The template's id.
- `parent_template_id` - (Optional) The template's parent template id.
- `last_update_time` - (Optional) The template's last update time.
- `name` - (Required) The template's name.
- `project_name` - (Optional) The template's project name.
- `rollback_template_content` - (Optional) The template's rollback template content.
- `software_type` - (Required) The template's software type.
- `software_variant` - (Optional) The template's software variant.
- `software_version` - (Optional) The template's software version.
- `template_content` - (Optional) The template's template content.
- `version` - (Optional) The template's version.
- `containing_templates` - (Optional) The template's containing templates. See [containing_templates](#containing_templates) below for details.
- `device_types` - (Required) The template's device types. See [device_types](#device_types) below for details.
- `rollback_template_params` - (Optional) The template's rollback template params. See [rollback_template_params](#template_params) below for details.
- `template_params` - (Optional) The template's template params. See [template_params](#template_params) below for details.

#### containing_templates

- `composite` - (Optional) The composite flag value of the contained template.
- `id` - (Optional) The id value of the contained template.
- `name` - (Optional) The name value of the contained template.
- `version` - (Optional) The version value of the contained template.

#### device_types

- `product_family` - (Optional) The product family.
- `product_series` - (Optional) The product series.
- `product_type` - (Optional) The product type.

#### template_params

- `binding` - (Optional) The param's binding.
- `data_type` - (Optional) The param's data_type.
- `default_value` - (Optional) The param's default_value.
- `description` - (Optional) The param's description.
- `display_name` - (Optional) The param's display_name.
- `group` - (Optional) The param's group.
- `id` - (Optional) The param's id.
- `instruction_text` - (Optional) The param's instruction_text.
- `key` - (Optional) The param's key.
- `not_param` - (Optional) The param's not_param.
- `order` - (Optional) The param's order.
- `param_array` - (Optional) The param's param_array.
- `parameter_name` - (Optional) The param's parameter_name.
- `provider` - (Optional) The param's provider.
- `required` - (Optional) The param's required.
- `range` - (Optional) The param's range. See below for details.
- `selection` - (Optional) The param's selection. See below for details.

##### range

- `id` - (Optional) The range's id.
- `max_value` - (Optional) The range's max value.
- `min_value` - (Optional) The range's min value.

##### selection

- `id` - (Optional) The selection's id.
- `selection_type` - (Optional) The selection's type.
- `selection_values` - (Optional) The selection's values. This is of TypeMap.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The template's updated time with format RFC850.
