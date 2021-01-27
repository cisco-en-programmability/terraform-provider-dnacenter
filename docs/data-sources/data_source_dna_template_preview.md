---
page_title: "dna_template_preview Data Source - terraform-provider-dnacenter"
subcategory: "Configuration Templates"
description: |-
  The dna_template_preview data source allows you to preview a Cisco DNA Center template.
---

# Data Source dna_template_preview

The dna_template_preview data source allows you to preview a Cisco DNA Center template.

## Example Usage

```hcl
data "dna_template_preview" "response" {
  provider   = dnacenter
  template_id = "82def961-26b7-4828-a7a2-579a4e10e446"
  params = {
    NetworkId = 1
  }
}
```

## Argument Reference

- `template_id` - (Required) The template_id params.
- `params` - (Optional) The params for the template preview. This is a TypeMap.

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

- `cli_preview` - The template's cli preview.
- `template_id` - The template preview's template id.
- `validation_errors` - The template preview's validation errors. See [validation_errors](#validation_errors) below for details.

#### validation_errors

- `type` - The validation error's type.
- `message` - The validation error's message.
