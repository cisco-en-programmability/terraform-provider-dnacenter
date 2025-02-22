---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_reports_executions Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Reports.
  Get details of all executions for a given report
---

# dnacenter_reports_executions (Data Source)

It performs read operation on Reports.

- Get details of all executions for a given report

## Example Usage

```terraform
data "dnacenter_reports_executions" "example" {
  provider  = dnacenter
  report_id = "string"
}

output "dnacenter_reports_executions_example" {
  value = data.dnacenter_reports_executions.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `report_id` (String) reportId path parameter. reportId of report

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `data_category` (String)
- `deliveries` (List of String)
- `execution_count` (Number)
- `executions` (List of Object) (see [below for nested schema](#nestedobjatt--item--executions))
- `name` (String)
- `report_id` (String)
- `report_was_executed` (String)
- `schedule` (String)
- `tags` (List of String)
- `view` (List of Object) (see [below for nested schema](#nestedobjatt--item--view))
- `view_group_id` (String)
- `view_group_version` (String)

<a id="nestedobjatt--item--executions"></a>
### Nested Schema for `item.executions`

Read-Only:

- `end_time` (Number)
- `errors` (List of String)
- `execution_id` (String)
- `process_status` (String)
- `request_status` (String)
- `start_time` (Number)
- `warnings` (List of String)


<a id="nestedobjatt--item--view"></a>
### Nested Schema for `item.view`

Read-Only:

- `description` (String)
- `field_groups` (List of String)
- `filters` (List of String)
- `format` (String)
- `name` (String)
- `view_id` (String)
- `view_info` (String)
