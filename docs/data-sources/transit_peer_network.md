---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_transit_peer_network Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation.
  Get Transit Peer Network Info from SD-Access
---

# dnacenter_transit_peer_network (Data Source)

It performs read operation.

- Get Transit Peer Network Info from SD-Access



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **transit_peer_network_name** (String) transitPeerNetworkName query parameter. Transit or Peer Network Name

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **ip_transit_settings** (List of Object) (see [below for nested schema](#nestedobjatt--item--ip_transit_settings))
- **sda_transit_settings** (List of Object) (see [below for nested schema](#nestedobjatt--item--sda_transit_settings))
- **transit_peer_network_name** (String)
- **transit_peer_network_type** (String)

<a id="nestedobjatt--item--ip_transit_settings"></a>
### Nested Schema for `item.ip_transit_settings`

Read-Only:

- **autonomous_system_number** (String)
- **routing_protocol_name** (String)


<a id="nestedobjatt--item--sda_transit_settings"></a>
### Nested Schema for `item.sda_transit_settings`

Read-Only:

- **transit_control_plane_settings** (List of Object) (see [below for nested schema](#nestedobjatt--item--sda_transit_settings--transit_control_plane_settings))

<a id="nestedobjatt--item--sda_transit_settings--transit_control_plane_settings"></a>
### Nested Schema for `item.sda_transit_settings.transit_control_plane_settings`

Read-Only:

- **device_management_ip_address** (String)
- **site_name_hierarchy** (String)

