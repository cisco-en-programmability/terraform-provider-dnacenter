---
page_title: "dna_pnp_workflow Resource - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_workflow resource allows you to configure a DNACenter PnP workflow.
---

# Resource dna_pnp_workflow

The dna_pnp_workflow resource allows you to configure a DNACenter PnP workflow.

## Example Usage

```hcl
resource "dna_pnp_workflow" "response" {
  provider = dnacenter
  item {
    name = "Workflow 1"
    tasks {
      name        = "Workflow 1 Task 1"
      task_seq_no = 0
      type        = "Reload"
    }
    tasks {
      name        = "Workflow 1 Task 2"
      task_seq_no = 1
      type        = "Reload"
    }
  }
}
```

## Argument Reference

- `item` - (Required) Item in the DNACenter PnP workflow. See [Workflow item](#workflow-item) below for details.

### Workflow item

- `id` - (Optional) The workflow's id.
- `add_to_inventory` - (Optional) The workflow's add to inventory flag.
- `added_on` - (Optional) The workflow's added on.
- `config_id` - (Optional) The workflow's config id.
- `curr_task_idx` - (Optional) The workflow's curr task idx.
- `description` - (Optional) The workflow's description.
- `end_time` - (Optional) The workflow's end time.
- `exec_time` - (Optional) The workflow's exec time.
- `image_id` - (Optional) The workflow's image id.
- `instance_type` - (Optional) The workflow's instance type.
- `lastupdate_on` - (Optional) The workflow's lastupdate on.
- `name` - (Required) The workflow's name.
- `start_time` - (Optional) The workflow's start time.
- `state` - (Optional) The workflow's state.
- `tasks` - (Required) The workflow's tasks. See [workflow tasks](#workflow_tasks) below for details.
- `tenant_id` - (Optional) The workflow's tenant id.
- `type` - (Optional) The workflow's type.
- `use_state` - (Optional) The workflow's use state.
- `version` - (Optional) The workflow's version.

### workflow_tasks

- `curr_work_item_idx` - (Optional) The task's curr work item index.
- `end_time` - (Optional) The task's end time.
- `name` - (Required) The task's name.
- `start_time` - (Optional) The task's start time.
- `state` - (Optional) The task's state.
- `task_seq_no` - (Required) The task's task sequential number.
- `time_taken` - (Optional) The task's time taken.
- `type` - (Required) The task's type.
- `work_item_list` - (Optional) The task's work item list. See [workflow item list](#workflow_item_list) below for details.

#### workflow_item_list

- `command` - (Optional) The item's command.
- `end_time` - (Optional) The item's end time.
- `output_str` - (Optional) The item's output string.
- `start_time` - (Optional) The item's start time.
- `state` - (Optional) The item's state.
- `time_taken` - (Optional) The item's time taken.

### workflow_parameters

- `config_list` - (Optional) The workflow parameters' config list. See below for details.
- `license_level` - (Optional) The workflow parameters' license level.
- `license_type` - (Optional) The workflow parameters' license type.
- `top_of_stack_serial_number` - (Optional) The workflow parameters' top of stack serial number.

### config_list

- `config_id` - (Optional) The config item's config id.
- `config_parameters` - (Optional) The config item's config parameters. See below for details.

#### config_parameters

- `key` - (Optional) The parameter's key.
- `value` - (Optional) The parameter's value.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The PnP workflow's updated time with format RFC850.
