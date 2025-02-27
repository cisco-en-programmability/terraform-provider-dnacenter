---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network_bugs_results_network_devices_network_device_id Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Compliance.
  Get network bug device by device id
---

# dnacenter_network_bugs_results_network_devices_network_device_id (Data Source)

It performs read operation on Compliance.

- Get network bug device by device id

## Example Usage

```terraform
data "dnacenter_network_bugs_results_network_devices_network_device_id" "example" {
  provider          = dnacenter
  network_device_id = "string"
}

output "dnacenter_network_bugs_results_network_devices_network_device_id_example" {
  value = data.dnacenter_network_bugs_results_network_devices_network_device_id.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_device_id` (String) networkDeviceId path parameter. Id of the network device

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `response` (List of Object) (see [below for nested schema](#nestedobjatt--items--response))
- `version` (String)

<a id="nestedobjatt--items--response"></a>
### Nested Schema for `items.response`

Read-Only:

- `bug_count` (Number)
- `comments` (String)
- `last_scan_time` (Number)
- `network_device_id` (String)
- `scan_mode` (String)
- `scan_status` (String)
