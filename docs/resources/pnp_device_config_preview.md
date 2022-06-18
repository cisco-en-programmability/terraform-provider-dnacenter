---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_pnp_device_config_preview Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Device Onboarding (PnP).
  Triggers a preview for site-based Day 0 Configuration
---

# dnacenter_pnp_device_config_preview (Resource)

It performs create operation on Device Onboarding (PnP).

- Triggers a preview for site-based Day 0 Configuration

~>**Warning:**
This resource does not represent a real-world entity in Cisco DNA Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco DNA Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_pnp_device_config_preview" "example" {
  provider = dnacenter
  parameters {

    device_id = "string"
    site_id   = "string"
    type      = "string"
  }
}

output "dnacenter_pnp_device_config_preview_example" {
  value = dnacenter_pnp_device_config_preview.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parameters** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- **device_id** (String)
- **site_id** (String)
- **type** (String)


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **complete** (String)
- **config** (String)
- **error** (String)
- **error_message** (String)
- **expired_time** (Number)
- **rf_profile** (String)
- **sensor_profile** (String)
- **site_id** (String)
- **start_time** (Number)
- **task_id** (String)

