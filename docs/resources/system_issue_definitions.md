---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_system_issue_definitions Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages read and update operations on Issues.
  Update issue trigger threshold, priority for the given id.
  Also enable or disable issue trigger for the given id. For detailed information about the usage of the API, please refer
  to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
  specs/blob/main/Assurance/CECatCenter_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml
---

# dnacenter_system_issue_definitions (Resource)

It manages read and update operations on Issues.

- Update issue trigger threshold, priority for the given id.
Also enable or disable issue trigger for the given id. For detailed information about the usage of the API, please refer
to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml

## Example Usage

```terraform
resource "dnacenter_system_issue_definitions" "example" {
  provider = dnacenter

  parameters {

    id                              = "string"
    issue_enabled                   = "false"
    priority                        = "string"
    synchronize_to_health_threshold = "false"
    threshold_value                 = 1.0
  }
}

output "dnacenter_system_issue_definitions_example" {
  value = dnacenter_system_issue_definitions.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `id` (String) id path parameter. Issue trigger definition id.

Optional:

- `issue_enabled` (String) Issue Enabled
- `priority` (String) Priority
- `synchronize_to_health_threshold` (String) Synchronize To Health Threshold
- `threshold_value` (Number) Threshold Value


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `category_name` (String)
- `default_priority` (String)
- `definition_status` (String)
- `description` (String)
- `device_type` (String)
- `display_name` (String)
- `id` (String)
- `issue_enabled` (String)
- `last_modified` (String)
- `name` (String)
- `priority` (String)
- `profile_id` (String)
- `synchronize_to_health_threshold` (String)
- `threshold_value` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import dnacenter_system_issue_definitions.example "id:=string"
```