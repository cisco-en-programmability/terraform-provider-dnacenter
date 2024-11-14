---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network_devices_issues_remediation_provision Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Compliance.
  Remediates configuration compliance issues. Compliance issues related to 'Routing', 'HA Remediation', 'Software
  Image', 'Securities Advisories', 'SD-Access Unsupported Configuration', 'Workflow', etc. will not be addressed by this
  API.
  Warning: Fixing compliance mismatches could result in a possible network flap.
---

# dnacenter_network_devices_issues_remediation_provision (Resource)

It performs create operation on Compliance.

- Remediates configuration compliance issues. Compliance issues related to 'Routing', 'HA Remediation', 'Software
Image', 'Securities Advisories', 'SD-Access Unsupported Configuration', 'Workflow', etc. will not be addressed by this
API.
Warning: Fixing compliance mismatches could result in a possible network flap.
~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "dnacenter_network_devices_issues_remediation_provision" "example" {
  provider = meraki
  id       = "string"
  parameters {

  }
}

output "dnacenter_network_devices_issues_remediation_provision_example" {
  value = dnacenter_network_devices_issues_remediation_provision.example
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

Required:

- `id` (String) id path parameter. Network device identifier


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)