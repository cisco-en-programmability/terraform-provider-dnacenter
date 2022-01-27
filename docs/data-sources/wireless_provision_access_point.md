---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_wireless_provision_access_point Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Wireless.
  Access Point Provision and ReProvision
---

# dnacenter_wireless_provision_access_point (Data Source)

It performs create operation on Wireless.

- Access Point Provision and ReProvision

## Example Usage

```terraform
data "dnacenter_wireless_provision_access_point" "example" {
  provider               = dnacenter
  custom_ap_group_name   = "string"
  custom_flex_group_name = ["string"]
  device_name            = "string"
  rf_profile             = "string"
  site_id                = "string"
  site_name_hierarchy    = "string"
  type                   = "string"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **payload** (Block List) Array of RequestWirelessAPProvision (see [below for nested schema](#nestedblock--payload))
- **persistbapioutput** (String) __persistbapioutput header parameter. Persist bapi sync response

### Read-Only

- **items** (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedblock--payload"></a>
### Nested Schema for `payload`

Optional:

- **custom_ap_group_name** (String) Custom AP group name
- **custom_flex_group_name** (List of String) ["Custom flex group name"]
- **device_name** (String) Device name
- **rf_profile** (String) Radio frequency profile name
- **site_id** (String) Site name hierarchy(ex: Global/...)
- **site_name_hierarchy** (String) Site name hierarchy(ex: Global/...)
- **type** (String) ApWirelessConfiguration


<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **execution_id** (String)
- **execution_url** (String)
- **message** (String)

