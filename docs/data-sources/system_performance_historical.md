---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_system_performance_historical Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Health and Performance.
  Retrieves the average values of cluster key performance indicators (KPIs), like CPU utilization, memory utilization or
  network rates grouped by time intervals within a specified time range. The data will be available from the past 24
  hours.
---

# dnacenter_system_performance_historical (Data Source)

It performs read operation on Health and Performance.

- Retrieves the average values of cluster key performance indicators (KPIs), like CPU utilization, memory utilization or
network rates grouped by time intervals within a specified time range. The data will be available from the past 24
hours.

## Example Usage

```terraform
data "dnacenter_system_performance_historical" "example" {
  provider   = dnacenter
  end_time   = 1609459200
  kpi        = "string"
  start_time = 1609459200
}

output "dnacenter_system_performance_historical_example" {
  value = data.dnacenter_system_performance_historical.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `end_time` (Number) endTime query parameter. This is the epoch end time in milliseconds upto which performance indicator need to be fetched
- `kpi` (String) kpi query parameter. Fetch historical data for this kpi. Valid values: cpu,memory,network
- `start_time` (Number) startTime query parameter. This is the epoch start time in milliseconds from which performance indicator need to be fetched

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `host_name` (String)
- `kpis` (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis))
- `version` (String)

<a id="nestedobjatt--item--kpis"></a>
### Nested Schema for `item.kpis`

Read-Only:

- `cpu_avg` (String)
- `data` (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--data))
- `legends` (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--legends))
- `memory_avg` (String)

<a id="nestedobjatt--item--kpis--data"></a>
### Nested Schema for `item.kpis.data`

Read-Only:

- `t1` (List of String)


<a id="nestedobjatt--item--kpis--legends"></a>
### Nested Schema for `item.kpis.legends`

Read-Only:

- `cpu` (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--legends--cpu))
- `memory` (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--legends--memory))
- `network_rx_rate` (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--legends--network_rx_rate))
- `network_tx_rate` (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--legends--network_tx_rate))

<a id="nestedobjatt--item--kpis--legends--cpu"></a>
### Nested Schema for `item.kpis.legends.network_tx_rate`

Read-Only:

- `units` (String)


<a id="nestedobjatt--item--kpis--legends--memory"></a>
### Nested Schema for `item.kpis.legends.network_tx_rate`

Read-Only:

- `units` (String)


<a id="nestedobjatt--item--kpis--legends--network_rx_rate"></a>
### Nested Schema for `item.kpis.legends.network_tx_rate`

Read-Only:

- `units` (String)


<a id="nestedobjatt--item--kpis--legends--network_tx_rate"></a>
### Nested Schema for `item.kpis.legends.network_tx_rate`

Read-Only:

- `units` (String)
