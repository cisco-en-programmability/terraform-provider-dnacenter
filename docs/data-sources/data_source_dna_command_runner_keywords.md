---
page_title: "dna_command_runner_keywords Data Source - terraform-provider-dnacenter"
subcategory: "Command Runner"
description: |-
  The command runner keywords data source allows you to list Cisco DNA Center command runner keywords.
---

# Data Source dna_command_runner_keywords

The command runner keywords data source allows you to list Cisco DNA Center command runner keywords.

## Example Usage

```hcl
data "dna_command_runner_keywords" "list" {
  provider = dnacenter
}
```

## Attributes Reference

The following attributes are exported.

- `response` - Cisco DNA Center command runner keywords.
