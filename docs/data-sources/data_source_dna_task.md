---
page_title: "dna_task Data Source - terraform-provider-dnacenter"
subcategory: "Task"
description: |-
  The dna_task data source allows you to get the result of a Cisco DNA Center task.
---

# Data Source dna_task

The dna_task data source allows you to get the result of a Cisco DNA Center task.

## Example Usage

```hcl
data "dna_task" "response" {
  provider   = dnacenter
  depends_on = [data.dna_command_runner_run_command.response]
  task_id    = data.dna_command_runner_run_command.response.task_id
}
```

## Argument Reference

- `task_id` - (Optional) Cisco DNA Center task_id.

## Attributes Reference

The following attributes are exported.

- `response` - Cisco DNA Center task's JSON response.
