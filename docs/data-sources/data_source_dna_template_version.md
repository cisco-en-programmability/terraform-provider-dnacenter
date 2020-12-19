---
page_title: "dna_template_version Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_template_version data source allows you to retrieve information about a particular DNACenter template version.
---

# Data Source dna_template_version

The dna_template_version data source allows you to retrieve information about a particular DNACenter template version.

## Example Usage

```hcl
data "dna_template_version" "response" {
  provider    = dnacenter
  template_id = "82def961-26b7-4828-a7a2-579a4e10e446"
}
```

## Argument Reference

- `template_id` - (Required) The template_id param.

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

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
