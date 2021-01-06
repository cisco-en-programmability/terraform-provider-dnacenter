---
page_title: "dna_discovery_job Data Source - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_discovery_job data source allows you to retrieve information about a particular DNACenter discovery job by id.
---

# Data Source dna_discovery_job

The dna_discovery_job data source allows you to retrieve information about a particular DNACenter discovery job by id.

Returns the list of discovery jobs for the given Discovery ID. The results can be optionally filtered based on IP.

## Example Usage

```hcl
data "dna_discovery_job" "response" {
  provider = dnacenter
  id = var.discovery_id
}
```

## Argument Reference

- `id` - (Required) The id param.
- `offset` - (Optional) The offset param.
- `limit` - (Optional) The limit param.
- `ip_address` - (Optional) The ip_address param.

## Attributes Reference

The following attributes are exported.

- `items` - Items in a DNACenter discovery job. See [Items](#items) below for details.

### Items

Each item contains the following attributes.

- `attribute_info` - The item's attribute info.
- `cli_status` - The item's cli status.
- `discovery_status` - The item's discovery status.
- `end_time` - The item's end time.
- `http_status` - The item's http status.
- `id` - The item's id.
- `inventory_collection_status` - The item's inventory collection status.
- `inventory_reachability_status` - The item's inventory reachability status.
- `ip_address` - The item's ip address.
- `job_status` - The item's job status.
- `name` - The item's name.
- `netconf_status` - The item's netconf status.
- `ping_status` - The item's ping status.
- `snmp_status` - The item's snmp status.
- `start_time` - The item's start time.
- `task_id` - The item's task id.
