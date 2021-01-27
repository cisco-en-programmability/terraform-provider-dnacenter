---
page_title: "dna_pnp_device Resource - terraform-provider-dnacenter"
subcategory: "Device Onboarding PnP"
description: |-
  The dna_pnp_device resource allows you to configure a Cisco DNA Center PnP device.
---

# Resource dna_pnp_device

The dna_pnp_device resource allows you to configure a Cisco DNA Center PnP device.

## Example Usage

```hcl
resource "dna_pnp_device" "response" {
  provider = dnacenter
  item {
    device_info {
      serial_number = "FOCTEST2"
    }
  }
}
```

## Argument Reference

- `item` - (Required) Item in a Cisco DNA Center PnP device. See [Device item](#device-item) below for details.

### Device item

- `id` - (Optional) The PnP device's id.
- `day_zero_config` - (Optional) The PnP device's day zero config. See [day_zero_config](#day_zero_config) below for details.
- `day_zero_config_preview` - (Optional) The PnP device's day zero config preview.
- `device_info` - (Required) The PnP device's device info. See [device_info](#device_info) below for details.
- `run_summary_list` - (Optional) The PnP device's run summary list. See [run_summary_list](#run_summary_list) below for details.
- `system_reset_workflow` - (Optional) The PnP device's system reset workflow. See [system_reset_workflow](#workflow) below for details.
- `system_workflow` - (Optional) The PnP device's system workflow. See [system_workflow](#workflow) below for details.
- `tenant_id` - (Optional) The PnP device's tenant id.
- `version` - (Optional) The PnP device's version.
- `workflow` - (Optional) The PnP device's workflow. See [workflow](#workflow) below for details.
- `workflow_parameters` - (Optional) The PnP device's workflow_parameters. See [workflow_parameters](#workflow_parameters) below for details.

### day_zero_config

- `config` - (Optional) The day zero config.

### device_info

- `aaa_credentials` - (Optional) The device info's aaa credentials.
- `added_on` - (Optional) The device info's added on.
- `addn_mac_addrs` - (Optional) The device info's addn mac addrs.
- `agent_type` - (Optional) The device info's agent type.
- `auth_status` - (Optional) The device info's auth status.
- `authenticated_mic_number` - (Optional) The device info's authenticated mic number.
- `authenticated_sudi_serial_no` - (Optional) The device info's authenticated sudi serial no.
- `capabilities_supported` - (Optional) The device info's capabilities supported.
- `cm_state` - (Optional) The device info's cm state.
- `description` - (Optional) The device info's description.
- `device_sudi_serial_nos` - (Optional) The device info's device sudi serial nos.
- `device_type` - (Optional) The device info's device type.
- `features_supported` - (Optional) The device info's features supported.
- `file_system_list` - (Optional) The device info's file system list.
- `first_contact` - (Optional) The device info's first contact.
- `hostname` - (Optional) The device info's hostname.
- `http_headers` - (Optional) The device info's http headers.
- `image_file` - (Optional) The device info's image file.
- `image_version` - (Optional) The device info's image version.
- `ip_interfaces` - (Optional) The device info's ip interfaces.
- `last_contact` - (Optional) The device info's last contact.
- `last_sync_time` - (Optional) The device info's last sync time.
- `last_update_on` - (Optional) The device info's last update on.
- `location` - (Optional) The device info's location.
- `mac_address` - (Optional) The device info's mac address.
- `mode` - (Optional) The device info's mode.
- `name` - (Optional) The device info's name.
- `neighbor_links` - (Optional) The device info's neighbor links.
- `onb_state` - (Optional) The device info's onb state.
- `pid` - (Optional) The device info's pid.
- `pnp_profile_list` - (Optional) The device info's pnp profile list.
- `populate_inventory` - (Optional) The device info's populate inventory.
- `pre_workflow_cli_ouputs` - (Optional) The device info's pre workflow cli ouputs.
- `project_id` - (Optional) The device info's project id.
- `project_name` - (Optional) The device info's project name.
- `reload_requested` - (Optional) The device info's reload requested.
- `serial_number` - (Optional) The device info's serial number.
- `site_id` - (Optional) The device info's site id.
- `site_name` - (Optional) The device info's site name.
- `smart_account_id` - (Optional) The device info's smart account id.
- `source` - (Optional) The device info's source.
- `stack` - (Optional) The device info's stack.
- `stack_info` - (Optional) The device info's stack info.
- `state` - (Optional) The device info's state.
- `sudi_required` - (Optional) The device info's sudi required.
- `tags` - (Optional) The device info's tags.
- `user_mic_numbers` - (Optional) The device info's user mic numbers.
- `user_sudi_serial_nos` - (Optional) The device info's user sudi serial nos.
- `virtual_account_id` - (Optional) The device info's virtual account id.
- `workflow_id` - (Optional) The device info's workflow id.
- `workflow_name` - (Optional) The device info's workflow name.

### run_summary_list

- `details` - (Optional) The list's details.
- `error_flag` - (Optional) The list's error flag.
- `history_task_info` - (Optional) The list's history task info.
- `timestamp` - (Optional) The list's timestamp.

### workflow

- `id` - (Optional) The workflow's id.
- `add_to_inventory` - (Optional) The workflow's add to inventory flag.
- `added_on` - (Optional) The workflow's added on.
- `config_id` - (Optional) The workflow's config id.
- `curr_task_idx` - (Optional) The workflow's curr task idx.
- `description` - (Optional) The workflow's description.
- `end_time` - (Optional) The workflow's end time.
- `exec_time` - (Optional) The workflow's exec time.
- `image_id` - (Optional) The workflow's image id.
- `instance_type` - (Optional) The workflow's instance type.
- `lastupdate_on` - (Optional) The workflow's lastupdate on.
- `name` - (Optional) The workflow's name.
- `start_time` - (Optional) The workflow's start time.
- `state` - (Optional) The workflow's state.
- `tasks` - (Optional) The workflow's tasks. See [workflow tasks](#workflow_tasks) below for details.
- `tenant_id` - (Optional) The workflow's tenant id.
- `type` - (Optional) The workflow's type.
- `use_state` - (Optional) The workflow's use state.
- `version` - (Optional) The workflow's version.

#### workflow_tasks

- `curr_work_item_idx` - (Optional) The task's curr work item index.
- `end_time` - (Optional) The task's end time.
- `name` - (Optional) The task's name.
- `start_time` - (Optional) The task's start time.
- `state` - (Optional) The task's state.
- `task_seq_no` - (Optional) The task's task sequential number.
- `time_taken` - (Optional) The task's time taken.
- `type` - (Optional) The task's type.
- `work_item_list` - (Optional) The task's work item list. See [workflow item list](#workflow_item_list) below for details.

##### workflow_item_list

- `command` - (Optional) The item's command.
- `end_time` - (Optional) The item's end time.
- `output_str` - (Optional) The item's output string.
- `start_time` - (Optional) The item's start time.
- `state` - (Optional) The item's state.
- `time_taken` - (Optional) The item's time taken.

### workflow_parameters

- `config_list` - (Optional) The workflow parameters' config list. See below for details.
- `license_level` - (Optional) The workflow parameters' license level.
- `license_type` - (Optional) The workflow parameters' license type.
- `top_of_stack_serial_number` - (Optional) The workflow parameters' top of stack serial number.

#### config_list

- `config_id` - (Optional) The config item's config id.
- `config_parameters` - (Optional) The config item's config parameters. See below for details.

##### config_parameters

- `key` - (Optional) The parameter's key.
- `value` - (Optional) The parameter's value.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The PnP device's updated time with format RFC850.
