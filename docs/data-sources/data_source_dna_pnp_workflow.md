---
page_title: "dna_pnp_workflow Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_pnp_workflow data source allows you to retrieve information about a particular DNACenter PnP workflow.
---

# Data Source dna_pnp_workflow

The dna_pnp_workflow data source allows you to retrieve information about a particular DNACenter PnP workflow.

## Example Usage

```hcl
data "dna_pnp_workflow" "response" {
  provider = dnacenter
  name     = ["Workflow 1"]
}
```

## Argument Reference

- `offset` - (Optional) The offset param.
- `limit` - (Optional) The limit param.
- `sort` - (Optional) The sort param.
- `sort_order` - (Optional) The sort_order param.
- `type` - (Optional) The type param.
- `name` - (Optional) The name param.

## Attributes Reference

The following attributes are exported.

- `items` - The item response. See [Items](#items) below for details.

### Items

- `id` - The workflow's id.
- `add_to_inventory` - The workflow's add to inventory flag.
- `added_on` - The workflow's added on.
- `config_id` - The workflow's config id.
- `curr_task_idx` - The workflow's curr task idx.
- `description` - The workflow's description.
- `end_time` - The workflow's end time.
- `exec_time` - The workflow's exec time.
- `image_id` - The workflow's image id.
- `instance_type` - The workflow's instance type.
- `lastupdate_on` - The workflow's lastupdate on.
- `name` - (Required) The workflow's name.
- `start_time` - The workflow's start time.
- `state` - The workflow's state.
- `tasks` - (Required) The workflow's tasks. See [workflow tasks](#workflow_tasks) below for details.
- `tenant_id` - The workflow's tenant id.
- `type` - The workflow's type.
- `use_state` - The workflow's use state.
- `version` - The workflow's version.

### workflow_tasks

- `curr_work_item_idx` - The task's curr work item index.
- `end_time` - The task's end time.
- `name` - (Required) The task's name.
- `start_time` - The task's start time.
- `state` - The task's state.
- `task_seq_no` - (Required) The task's task sequential number.
- `time_taken` - The task's time taken.
- `type` - (Required) The task's type.
- `work_item_list` - The task's work item list. See [workflow item list](#workflow_item_list) below for details.

#### workflow_item_list

- `command` - The item's command.
- `end_time` - The item's end time.
- `output_str` - The item's output string.
- `start_time` - The item's start time.
- `state` - The item's state.
- `time_taken` - The item's time taken.

### workflow_parameters

- `config_list` - The workflow parameters' config list. See below for details.
- `license_level` - The workflow parameters' license level.
- `license_type` - The workflow parameters' license type.
- `top_of_stack_serial_number` - The workflow parameters' top of stack serial number.

### config_list

- `config_id` - The config item's config id.
- `config_parameters` - The config item's config parameters. See below for details.

#### config_parameters

- `key` - The parameter's key.
- `value` - The parameter's value.
