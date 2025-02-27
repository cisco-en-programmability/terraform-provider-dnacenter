---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_flexible_report_schedule Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages read and update operations on Reports.
  Update schedule of flexible report
---

# dnacenter_flexible_report_schedule (Resource)

It manages read and update operations on Reports.

- Update schedule of flexible report

## Example Usage

```terraform
resource "dnacenter_flexible_report_schedule" "example" {
  provider = dnacenter
  parameters {

    report_id = "string"
    schedule  = "string"
  }
}

output "dnacenter_flexible_report_schedule_example" {
  value = dnacenter_flexible_report_schedule.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of String)
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `report_id` (String) reportId path parameter. Id of the report

Optional:

- `schedule` (String) Schedule information

## Import

Import is supported using the following syntax:

```shell
terraform import dnacenter_flexible_report_schedule.example "report_id:=string"
```
