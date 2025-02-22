---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_wireless_controllers_managed_ap_locations_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Wireless.
  Retrieves the count of Managed AP locations, including Primary Managed AP Locations, Secondary Managed AP Locations,
  and Anchor Managed AP Locations, associated with the specific Wireless Controller.
---

# dnacenter_wireless_controllers_managed_ap_locations_count (Data Source)

It performs read operation on Wireless.

- Retrieves the count of Managed AP locations, including Primary Managed AP Locations, Secondary Managed AP Locations,
and Anchor Managed AP Locations, associated with the specific Wireless Controller.

## Example Usage

```terraform
data "dnacenter_wireless_controllers_managed_ap_locations_count" "example" {
  provider          = dnacenter
  network_device_id = "string"
}

output "dnacenter_wireless_controllers_managed_ap_locations_count_example" {
  value = data.dnacenter_wireless_controllers_managed_ap_locations_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_device_id` (String) networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `anchor_managed_ap_locations_count` (Number)
- `primary_managed_ap_locations_count` (Number)
- `secondary_managed_ap_locations_count` (Number)
