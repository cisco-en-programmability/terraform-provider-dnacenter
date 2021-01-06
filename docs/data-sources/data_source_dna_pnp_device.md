---
page_title: "dna_pnp_device Data Source - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_device data source allows you to retrieve information about a particular DNACenter PnP device.
---

# Data Source dna_pnp_device

The dna_pnp_device data source allows you to retrieve information about a particular DNACenter PnP device.

## Example Usage

```hcl
data "dna_pnp_device" "result" {
  provider = dnacenter
}
```

## Argument Reference

- `sort` - (Optional) The sort param.
- `sort_order` - (Optional) The sort order param.
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
- `mac_address` - (Optional) The mac address param.
- `hostname` - (Optional) The hostname param.
- `site_name` - (Optional) The site name param.

## Attributes Reference

The following attributes are exported.

- `items` - The item response. See [Items](#items) below for details.

### Items

- `id` - The PnP device's id.
- `day_zero_config` - The PnP device's day zero config. See [day_zero_config](#day_zero_config) below for details.
- `day_zero_config_preview` - The PnP device's day zero config preview.
- `device_info` - (Required) The PnP device's device info. See [device_info](#device_info) below for details.
- `run_summary_list` - The PnP device's run summary list. See [run_summary_list](#run_summary_list) below for details.
- `system_reset_workflow` - The PnP device's system reset workflow. See [system_reset_workflow](#workflow) below for details.
- `system_workflow` - The PnP device's system workflow. See [system_workflow](#workflow) below for details.
- `tenant_id` - The PnP device's tenant id.
- `version` - The PnP device's version.
- `workflow` - The PnP device's workflow. See [workflow](#workflow) below for details.
- `workflow_parameters` - The PnP device's workflow_parameters. See [workflow_parameters](#workflow_parameters) below for details.

#### day_zero_config

- `config` - The day zero config.

#### device_info

- `aaa_credentials` - The device info's aaa credentials.
- `added_on` - The device info's added on.
- `addn_mac_addrs` - The device info's addn mac addrs.
- `agent_type` - The device info's agent type.
- `auth_status` - The device info's auth status.
- `authenticated_mic_number` - The device info's authenticated mic number.
- `authenticated_sudi_serial_no` - The device info's authenticated sudi serial no.
- `capabilities_supported` - The device info's capabilities supported.
- `cm_state` - The device info's cm state.
- `description` - The device info's description.
- `device_sudi_serial_nos` - The device info's device sudi serial nos.
- `device_type` - The device info's device type.
- `features_supported` - The device info's features supported.
- `file_system_list` - The device info's file system list.
- `first_contact` - The device info's first contact.
- `hostname` - The device info's hostname.
- `http_headers` - The device info's http headers.
- `image_file` - The device info's image file.
- `image_version` - The device info's image version.
- `ip_interfaces` - The device info's ip interfaces.
- `last_contact` - The device info's last contact.
- `last_sync_time` - The device info's last sync time.
- `last_update_on` - The device info's last update on.
- `location` - The device info's location.
- `mac_address` - The device info's mac address.
- `mode` - The device info's mode.
- `name` - The device info's name.
- `neighbor_links` - The device info's neighbor links.
- `onb_state` - The device info's onb state.
- `pid` - The device info's pid.
- `pnp_profile_list` - The device info's pnp profile list.
- `populate_inventory` - The device info's populate inventory.
- `pre_workflow_cli_ouputs` - The device info's pre workflow cli ouputs.
- `project_id` - The device info's project id.
- `project_name` - The device info's project name.
- `reload_requested` - The device info's reload requested.
- `serial_number` - The device info's serial number.
- `site_id` - The device info's site id.
- `site_name` - The device info's site name.
- `smart_account_id` - The device info's smart account id.
- `source` - The device info's source.
- `stack` - The device info's stack.
- `stack_info` - The device info's stack info.
- `state` - The device info's state.
- `sudi_required` - The device info's sudi required.
- `tags` - The device info's tags.
- `user_mic_numbers` - The device info's user mic numbers.
- `user_sudi_serial_nos` - The device info's user sudi serial nos.
- `virtual_account_id` - The device info's virtual account id.
- `workflow_id` - The device info's workflow id.
- `workflow_name` - The device info's workflow name.

#### run_summary_list

- `details` - The list's details.
- `error_flag` - The list's error flag.
- `history_task_info` - The list's history task info.
- `timestamp` - The list's timestamp.

#### workflow

- `id` - The workflow's id.
- `add_to_inventory` - The workflow's add to inventory flag.
- `added_on` - The workflow's added on.
- `config_id` - The workflow's config id.
- `curr_task_idx` - The workflow's curr task idx.
- `description` - The workflow's description.
- `end_time` - The workflow's end time.
- `exec_time` - The workflow's exec time.
- `image_id` - The workflow's image id.
- `instance_type` - The workflow's instance type.
- `lastupdate_on` - The workflow's lastupdate on.
- `name` - The workflow's name.
- `start_time` - The workflow's start time.
- `state` - The workflow's state.
- `tasks` - The workflow's tasks. See [workflow tasks](#workflow_tasks) below for details.
- `tenant_id` - The workflow's tenant id.
- `type` - The workflow's type.
- `use_state` - The workflow's use state.
- `version` - The workflow's version.

##### workflow_tasks

- `curr_work_item_idx` - The task's curr work item index.
- `end_time` - The task's end time.
- `name` - The task's name.
- `start_time` - The task's start time.
- `state` - The task's state.
- `task_seq_no` - The task's task sequential number.
- `time_taken` - The task's time taken.
- `type` - The task's type.
- `work_item_list` - The task's work item list. See [workflow item list](#workflow_item_list) below for details.

###### workflow_item_list

- `command` - The item's command.
- `end_time` - The item's end time.
- `output_str` - The item's output string.
- `start_time` - The item's start time.
- `state` - The item's state.
- `time_taken` - The item's time taken.

#### workflow_parameters

- `config_list` - The workflow parameters' config list. See below for details.
- `license_level` - The workflow parameters' license level.
- `license_type` - The workflow parameters' license type.
- `top_of_stack_serial_number` - The workflow parameters' top of stack serial number.

##### config_list

- `config_id` - The config item's config id.
- `config_parameters` - The config item's config parameters. See below for details.

###### config_parameters

- `key` - The parameter's key.
- `value` - The parameter's value.
