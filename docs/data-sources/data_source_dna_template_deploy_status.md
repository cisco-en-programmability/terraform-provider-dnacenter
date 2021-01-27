---
page_title: "dna_template_deploy_status Data Source - terraform-provider-dnacenter"
subcategory: "Configuration Templates"
description: |-
  The dna_template_deploy_status data source allows you to retrieve information about a particular Cisco DNA Center template deploy status.
---

# Data Source dna_template_deploy_status

The dna_template_deploy_status data source allows you to retrieve information about a particular Cisco DNA Center template deploy status.

## Example Usage

```hcl
data "dna_template_deploy_status" "response" {
  provider      = dnacenter
  deployment_id = data.dna_template_deploy.response.item.0.deployment_id
}
```

## Argument Reference

- `deployment_id` - (Required) The deployment_id param.

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

- `deployment_id` - The item's deployment id.
- `deployment_name` - The item's deployment name.
- `duration` - The item's duration.
- `end_time` - The item's end time.
- `project_name` - The item's project name.
- `start_time` - The item's start time.
- `status` - The item's status.
- `template_name` - The item's template name.
- `template_version` - The item's template version.
- `devices` - The item's devices. See [Devices](#devices) below for details.

#### Devices

- `device_id` - The device id.
- `duration` - The device's duration.
- `end_time` - The device's end time.
- `ip_address` - The device's ip address.
- `name` - The device's name.
- `start_time` - The device's start time.
- `status` - The device's status.
