---
page_title: "dna_task Data Source - terraform-provider-dnacenter"
subcategory: "Task"
description: |-
  The dna_task data source allows you to get the result of a DNACenter task.
---

# Data Source dna_task

The dna_task data source allows you to get the result of a DNACenter task.

## Example Usage

```hcl
data "dna_task" "response" {
  provider   = dnacenter
  depends_on = [data.dna_command_runner_run_command.response]
  task_id    = data.dna_command_runner_run_command.response.task_id
}
```

## Argument Reference

- `task_id` - (Optional) DNACenter task_id.

## Attributes Reference

The following attributes are exported.

- `response` - DNACenter task's JSON response.
