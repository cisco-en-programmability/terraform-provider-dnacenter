---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_license_device_license_details Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Licenses.
  Get detailed license information of a device.
---

# dnacenter_license_device_license_details (Data Source)

It performs read operation on Licenses.

- Get detailed license information of a device.

## Example Usage

```terraform
data "dnacenter_license_device_license_details" "example" {
  provider    = dnacenter
  device_uuid = "string"
}

output "dnacenter_license_device_license_details_example" {
  value = data.dnacenter_license_device_license_details.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_uuid` (String) device_uuid path parameter. Id of device

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `access_points` (List of Object) (see [below for nested schema](#nestedobjatt--items--access_points))
- `chassis_details` (List of Object) (see [below for nested schema](#nestedobjatt--items--chassis_details))
- `device_name` (String)
- `device_type` (String)
- `device_uuid` (String)
- `dna_level` (String)
- `evaluation_license_expiry` (String)
- `feature_license` (List of String)
- `has_sup_cards` (String)
- `ip_address` (String)
- `is_license_expired` (String)
- `is_stacked_device` (String)
- `license_mode` (String)
- `mac_address` (String)
- `model` (String)
- `network_license` (String)
- `site` (String)
- `sntc_status` (String)
- `software_version` (String)
- `stacked_devices` (List of Object) (see [below for nested schema](#nestedobjatt--items--stacked_devices))
- `udi` (String)
- `virtual_account_name` (String)

<a id="nestedobjatt--items--access_points"></a>
### Nested Schema for `items.access_points`

Read-Only:

- `ap_type` (String)
- `count` (String)


<a id="nestedobjatt--items--chassis_details"></a>
### Nested Schema for `items.chassis_details`

Read-Only:

- `board_serial_number` (String)
- `modules` (List of Object) (see [below for nested schema](#nestedobjatt--items--chassis_details--modules))
- `port` (Number)
- `supervisor_cards` (List of Object) (see [below for nested schema](#nestedobjatt--items--chassis_details--supervisor_cards))

<a id="nestedobjatt--items--chassis_details--modules"></a>
### Nested Schema for `items.chassis_details.modules`

Read-Only:

- `id` (String)
- `module_name` (String)
- `module_type` (String)
- `serial_number` (String)


<a id="nestedobjatt--items--chassis_details--supervisor_cards"></a>
### Nested Schema for `items.chassis_details.supervisor_cards`

Read-Only:

- `serial_number` (String)
- `status` (String)
- `supervisor_card_type` (String)



<a id="nestedobjatt--items--stacked_devices"></a>
### Nested Schema for `items.stacked_devices`

Read-Only:

- `id` (String)
- `mac_address` (String)
- `role` (String)
- `serial_number` (String)


