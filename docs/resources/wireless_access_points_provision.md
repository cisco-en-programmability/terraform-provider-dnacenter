---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_wireless_access_points_provision Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Wireless.
  This data source action is used to provision access points
---

# dnacenter_wireless_access_points_provision (Resource)

It performs create operation on Wireless.

- This data source action is used to provision access points
~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "dnacenter_wireless_access_points_provision" "example" {
  provider = meraki
  parameters {

    ap_zone_name = "string"
    network_devices {

      device_id = "string"
      mesh_role = "string"
    }
    rf_profile_name = "string"
    site_id         = "string"
  }
}

output "dnacenter_wireless_access_points_provision_example" {
  value = dnacenter_wireless_access_points_provision.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `ap_zone_name` (String) AP Zone Name. A custom AP Zone should be passed if no rfProfileName is provided.
- `network_devices` (Block List) (see [below for nested schema](#nestedblock--parameters--network_devices))
- `rf_profile_name` (String) RF Profile Name. RF Profile is not allowed for custom AP Zones.
- `site_id` (String) Site ID

<a id="nestedblock--parameters--network_devices"></a>
### Nested Schema for `parameters.network_devices`

Optional:

- `device_id` (String) Network device ID of access points
- `mesh_role` (String) Mesh Role (Applicable only when AP is in Bridge Mode)



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)