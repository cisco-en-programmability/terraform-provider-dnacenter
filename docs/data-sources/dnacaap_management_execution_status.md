---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_dnacaap_management_execution_status Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Task.
  Retrieves the execution details of a Business API
---

# dnacenter_dnacaap_management_execution_status (Data Source)

It performs read operation on Task.

- Retrieves the execution details of a Business API

## Example Usage

```terraform
data "dnacenter_dnacaap_management_execution_status" "example" {
  provider     = dnacenter
  execution_id = "string"
}

output "dnacenter_dnacaap_management_execution_status_example" {
  value = data.dnacenter_dnacaap_management_execution_status.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `execution_id` (String) executionId path parameter. Execution Id of API

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `bapi_error` (String)
- `bapi_execution_id` (String)
- `bapi_key` (String)
- `bapi_name` (String)
- `bapi_sync_response` (String)
- `bapi_sync_response_json` (String)
- `end_time` (String)
- `end_time_epoch` (Number)
- `runtime_instance_id` (String)
- `start_time` (String)
- `start_time_epoch` (Number)
- `status` (String)
- `time_duration` (Number)
