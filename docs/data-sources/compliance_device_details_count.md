---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_compliance_device_details_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Compliance.
  Return  Compliance Count Detail
---

# dnacenter_compliance_device_details_count (Data Source)

It performs read operation on Compliance.

- Return  Compliance Count Detail

## Example Usage

```terraform
data "dnacenter_compliance_device_details_count" "example" {
  provider          = dnacenter
  compliance_status = "string"
  compliance_type   = "string"
}

output "dnacenter_compliance_device_details_count_example" {
  value = data.dnacenter_compliance_device_details_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **compliance_status** (String) complianceStatus query parameter. Compliance status can have value among 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'
- **compliance_type** (String) complianceType query parameter. complianceType can have any value among 'NETWORK_PROFILE', 'IMAGE', 'APPLICATION_VISIBILITY', 'FABRIC', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **response** (Number)
- **version** (String)

