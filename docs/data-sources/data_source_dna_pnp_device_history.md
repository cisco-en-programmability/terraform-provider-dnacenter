---
page_title: "dna_pnp_device_history Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_pnp_device_history data source allows you to retrieve information about a particular DNACenter PnP device history.
---

# Data Source dna_pnp_device_history

The dna_pnp_device_history data source allows you to retrieve information about a particular DNACenter PnP device history.

## Example Usage

```hcl
data "dna_pnp_device_history" "response" {
  provider      = dnacenter
  serial_number = "FOCTEST1"
}
```

## Argument Reference

- `serial_number` - (Required) The serial_number param.
- `sort` - (Optional) The sort param. Available values are "asc" and "des".
- `sort_order` - (Optional) The sort_order param.

## Attributes Reference

The following attributes are exported.

- `items` - Items in a DNACenter PnP device history. See [Items](#items) below for details.

### Items

Each item contains the following attributes.

- `details` - The item's details.
- `error_flag` - The item's error flag.
- `history_task_info` - The item's history task info. See [history_task_info](#history_task_info) below for details.
- `timestamp` - The item's timestamp.

#### history_task_info

- `addn_details` - The task's addn details.
- `name` - The task's name.
- `time_taken` - The task's time taken.
- `type` - The task's type.
- `work_item_list` - The task's work item list.
