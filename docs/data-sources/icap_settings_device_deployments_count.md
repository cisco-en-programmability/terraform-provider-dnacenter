---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_icap_settings_device_deployments_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Sensors.
  Returns the count of device deployment status(s) based on filter criteria. For detailed information about the usage of
  the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
  center-api-specs/blob/main/Assurance/CECatCenterOrg-ICAPAPIs-1.0.0-resolved.yaml
---

# dnacenter_icap_settings_device_deployments_count (Data Source)

It performs read operation on Sensors.

- Returns the count of device deployment status(s) based on filter criteria. For detailed information about the usage of
the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml

## Example Usage

```terraform
data "dnacenter_icap_settings_device_deployments_count" "example" {
  provider           = dnacenter
  deploy_activity_id = "string"
  network_device_ids = "string"
}

output "dnacenter_icap_settings_device_deployments_count_example" {
  value = data.dnacenter_icap_settings_device_deployments_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `deploy_activity_id` (String) deployActivityId query parameter. activity from the /deploy task response
- `network_device_ids` (String) networkDeviceIds query parameter. device ids, retrievable from the id attribute in intent/api/v1/network-device

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)
