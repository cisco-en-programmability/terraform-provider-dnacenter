---
page_title: "dna_template_details Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_template_details data source allows you to retrieve information about a particular DNACenter template details.
---

# Data Source dna_template_details

The dna_template_details data source allows you to retrieve information about a particular DNACenter template details.

## Example Usage

```hcl
data "dna_template_details" "response" {
  provider    = dnacenter
  template_id = "82def961-26b7-4828-a7a2-579a4e10e446"
  # latest_version = true
}
```

## Argument Reference

- `template_id` - (Required) The template_id param.
- `latest_version` - (Optional) The latest_version param.

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

- `project_id` - The template's project id.
- `author` - The template's author.
- `composite` - The template's composite flag.
- `create_time` - The template's create time.
- `description` - The template's description.
- `failure_policy` - The template's failure policy.
- `id` - The template's id.
- `parent_template_id` - The template's parent template id.
- `last_update_time` - The template's last update time.
- `name` - The template's name.
- `project_name` - The template's project name.
- `rollback_template_content` - The template's rollback template content.
- `software_type` - The template's software type.
- `software_variant` - The template's software variant.
- `software_version` - The template's software version.
- `template_content` - The template's template content.
- `version` - The template's version.
- `containing_templates` - The template's containing templates. See [containing_templates](#containing_templates) below for details.
- `device_types` - The template's device types. See [device_types](#device_types) below for details.
- `rollback_template_params` - The template's rollback template params. See [rollback_template_params](#template_params) below for details.
- `template_params` - The template's template params. See [template_params](#template_params) below for details.

#### containing_templates

- `composite` - The composite flag value of the contained template.
- `id` - The id value of the contained template.
- `name` - The name value of the contained template.
- `version` - The version value of the contained template.

#### device_types

- `product_family` - The product family.
- `product_series` - The product series.
- `product_type` - The product type.

#### template_params

- `binding` - The param's binding.
- `data_type` - The param's data_type.
- `default_value` - The param's default_value.
- `description` - The param's description.
- `display_name` - The param's display_name.
- `group` - The param's group.
- `id` - The param's id.
- `instruction_text` - The param's instruction_text.
- `key` - The param's key.
- `not_param` - The param's not_param.
- `order` - The param's order.
- `param_array` - The param's param_array.
- `parameter_name` - The param's parameter_name.
- `provider` - The param's provider.
- `required` - The param's required.
- `range` - The param's range. See below for details.
- `selection` - The param's selection. See below for details.

##### range

- `id` - The range's id.
- `max_value` - The range's max value.
- `min_value` - The range's min value.

##### selection

- `id` - The selection's id.
- `selection_type` - The selection's type.
- `selection_values` - The selection's values. This is of TypeMap.
