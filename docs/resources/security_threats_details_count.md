---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_security_threats_details_count Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Devices.
  The details count for the Rogue and aWIPS threats
---

# dnacenter_security_threats_details_count (Resource)

It performs create operation on Devices.

- The details count for the Rogue and aWIPS threats
~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "dnacenter_security_threats_details_count" "example" {
  provider = meraki
  parameters {

    end_time      = 1
    is_new_threat = "false"
    limit         = 1
    offset        = 1
    site_id       = ["string"]
    start_time    = 1
    threat_level  = ["string"]
    threat_type   = ["string"]
  }
}

output "dnacenter_security_threats_details_count_example" {
  value = dnacenter_security_threats_details_count.example
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

- `end_time` (Number) End Time
- `is_new_threat` (String) Is New Threat
- `limit` (Number) Limit
- `offset` (Number) Offset
- `site_id` (List of String) Site Id
- `start_time` (Number) Start Time
- `threat_level` (List of String) Threat Level
- `threat_type` (List of String) Threat Type


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `response` (Number)
- `version` (String)