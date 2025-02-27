---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_task_tree Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Task.
  Returns a task with its children tasks by based on their id
---

# dnacenter_task_tree (Data Source)

It performs read operation on Task.

- Returns a task with its children tasks by based on their id

## Example Usage

```terraform
data "dnacenter_task_tree" "example" {
  provider = dnacenter
  task_id  = "string"
}

output "dnacenter_task_tree_example" {
  value = data.dnacenter_task_tree.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `task_id` (String) taskId path parameter. UUID of the Task

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `additional_status_url` (String)
- `data` (String)
- `end_time` (Number)
- `error_code` (String)
- `error_key` (String)
- `failure_reason` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `is_error` (String)
- `last_update` (Number)
- `operation_id_list` (List of String)
- `parent_id` (String)
- `progress` (String)
- `root_id` (String)
- `service_type` (String)
- `start_time` (Number)
- `username` (String)
- `version` (Number)
