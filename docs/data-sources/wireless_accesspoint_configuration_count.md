---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_wireless_accesspoint_configuration_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Wireless.
  Get Access Point Configuration Count
---

# dnacenter_wireless_accesspoint_configuration_count (Data Source)

It performs read operation on Wireless.

- Get Access Point Configuration Count

## Example Usage

```terraform
data "dnacenter_wireless_accesspoint_configuration_count" "example" {
  provider       = dnacenter
  ap_mode        = "string"
  ap_model       = "string"
  mesh_role      = "string"
  provisioned    = "string"
  wlc_ip_address = "string"
}

output "dnacenter_wireless_accesspoint_configuration_count_example" {
  value = data.dnacenter_wireless_accesspoint_configuration_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ap_mode` (String) apMode query parameter. AP Mode. Allowed values are Local, Bridge, Monitor, FlexConnect, Sniffer, Rogue Detector, SE-Connect, Flex+Bridge, Sensor.
- `ap_model` (String) apModel query parameter. AP Model
- `mesh_role` (String) meshRole query parameter. Mesh Role. Allowed values are RAP or MAP
- `provisioned` (String) provisioned query parameter. Indicate whether AP provisioned or not. Allowed values are True or False
- `wlc_ip_address` (String) wlcIpAddress query parameter. WLC IP Address

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)
