---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_wireless_provision_access_point Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Wireless.
          - Access Point Provision and ReProvision
---

# dnacenter_wireless_provision_access_point (Resource)

It performs create operation on Wireless.
		- Access Point Provision and ReProvision


~>**Warning:**
This resource does not represent a real-world entity in Cisco DNA Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco DNA Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.


## Example Usage

```terraform
resource "dnacenter_wireless_provision_access_point" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    custom_ap_group_name   = "string"
    custom_flex_group_name = ["string"]
    device_name            = "string"
    rf_profile             = "string"
    site_id                = "string"
    site_name_hierarchy    = "string"
    type                   = "string"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parameters** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **items** (List of Object) (see [below for nested schema](#nestedatt--items))
- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- **payload** (Block List) Array of RequestWirelessAPProvision (see [below for nested schema](#nestedblock--parameters--payload))
- **persistbapioutput** (String) __persistbapioutput header parameter. Persist bapi sync response

<a id="nestedblock--parameters--payload"></a>
### Nested Schema for `parameters.payload`

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

