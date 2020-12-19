---
page_title: "dna_discovery_summary Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_discovery_summary data source allows you to count the devices from a discovery job based on given filters.
---

# Data Source dna_discovery_summary

The applications data source allows you to count the devices from a discovery job based on given filters.

## Example Usage

```hcl
data "dna_discovery_summary" "amount" {
  provider = dnacenter
  id = discovery.id
}
```

## Argument Reference

- `id` - (Required) The id param.
- `task_id` - (Optional) The task_id param.
- `sort_by` - (Optional) The sort_by param.
- `sort_order` - (Optional) The sort_order param.
- `ip_address` - (Optional) The ip_address param.
- `ping_status` - (Optional) The ping_status param.
- `snmp_status` - (Optional) The snmp_status param.
- `cli_status` - (Optional) The cli_status param.
- `netconf_status` - (Optional) The netconf_status param.
- `http_status` - (Optional) The http_status param.

## Attributes Reference

The following attributes are exported.

- `response` - DNACenter discovery devices count.
