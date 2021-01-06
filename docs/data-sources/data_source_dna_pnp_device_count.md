---
page_title: "dna_pnp_device_count Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_device_count data source allows you to count the amount of DNACenter discoveries discovered.
---

# Data Source dna_pnp_device_count

The dna_pnp_device_count data source allows you to count the amount of DNACenter discoveries discovered.

## Example Usage

```hcl
data "dna_pnp_device_count" "response" {
  provider = dnacenter
  # name     = ["FOCTEST1"]
}
```

## Argument Reference

- `serial_number` - (Optional) The serial number param.
- `state` - (Optional) The state param.
- `onb_state` - (Optional) The onb state param.
- `cm_state` - (Optional) The cm state param.
- `name` - (Optional) The name param.
- `pid` - (Optional) The pid param.
- `source` - (Optional) The source param.
- `project_id` - (Optional) The project id param.
- `workflow_id` - (Optional) The workflow id param.
- `project_name` - (Optional) The project name param.
- `workflow_name` - (Optional) The workflow name param.
- `smart_account_id` - (Optional) The smart account id param.
- `virtual_account_id` - (Optional) The virtual account id param.
- `last_contact` - (Optional) The last contact param.

## Attributes Reference

The following attributes are exported.

- `response` - DNACenter discoveries discovered count.
