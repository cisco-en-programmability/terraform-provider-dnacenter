---
page_title: "dna_command_runner_run_command Data Source - terraform-provider-dnacenter"
subcategory: "Command Runner"
description: |-
  The dna_command_runner_run_command data source allows you to run a DNACenter command.
---

# Data Source dna_command_runner_run_command

The dna_command_runner_run_command data source allows you to run a DNACenter command.

## Example Usage

```hcl
data "dna_command_runner_run_command" "list" {
  provider = dnacenter
  commands     = ["pwd"]
  device_uuids = ["a9b86e42-6573-4f5d-a0bf-a743e290f46a"]
}
```

## Argument Reference

- `commands` - (Optional) DNACenter commands.
- `device_uuids` - (Optional) DNACenter device uuids.
- `description` - (Optional) DNACenter description.
- `timeout` - (Optional) DNACenter timeout.
- `name` - (Optional) DNACenter name.

## Attributes Reference

The following attributes are exported.

- `task_id` - DNACenter command runner task_id.
