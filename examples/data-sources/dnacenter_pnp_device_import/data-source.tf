
data "dnacenter_pnp_device_import" "example" {
  provider = dnacenter
  id       = "string"
  device_info {

    aaa_credentials {

      password = "******"
      username = "string"
    }
    added_on                     = 1
    addn_mac_addrs               = ["string"]
    agent_type                   = "string"
    auth_status                  = "string"
    authenticated_sudi_serial_no = "string"
    capabilities_supported       = ["string"]
    cm_state                     = "string"
    description                  = "string"
    device_sudi_serial_nos       = ["string"]
    device_type                  = "string"
    features_supported           = ["string"]
    file_system_list {

      freespace = 1
      name      = "string"
      readable  = "false"
      size      = 1
      type      = "string"
      writeable = "false"
    }
    first_contact = 1
    hostname      = "string"
    http_headers {

      key   = "string"
      value = "string"
    }
    image_file    = "string"
    image_version = "string"
    ip_interfaces {

      ipv4_address      = ["string"]
      ipv6_address_list = ["string"]
      mac_address       = "string"
      name              = "string"
      status            = "string"
    }
    last_contact   = 1
    last_sync_time = 1
    last_update_on = 1
    location {

      address   = "string"
      altitude  = "string"
      latitude  = "string"
      longitude = "string"
      site_id   = "string"
    }
    mac_address = "string"
    mode        = "string"
    name        = "string"
    neighbor_links {

      local_interface_name        = "string"
      local_mac_address           = "string"
      local_short_interface_name  = "string"
      remote_device_name          = "string"
      remote_interface_name       = "string"
      remote_mac_address          = "string"
      remote_platform             = "string"
      remote_short_interface_name = "string"
      remote_version              = "string"
    }
    onb_state = "string"
    pid       = "string"
    pnp_profile_list {

      created_by        = "string"
      discovery_created = "false"
      primary_endpoint {

        certificate  = "string"
        fqdn         = "string"
        ipv4_address = ["string"]
        ipv6_address = ["string"]
        port         = 1
        protocol     = "string"
      }
      profile_name = "string"
      secondary_endpoint {

        certificate  = "string"
        fqdn         = "string"
        ipv4_address = ["string"]
        ipv6_address = ["string"]
        port         = 1
        protocol     = "string"
      }
    }
    populate_inventory = "false"
    pre_workflow_cli_ouputs {

      cli        = "string"
      cli_output = "string"
    }
    project_id       = "string"
    project_name     = "string"
    reload_requested = "false"
    serial_number    = "string"
    smart_account_id = "string"
    source           = "string"
    stack            = "false"
    stack_info {

      is_full_ring = "false"
      stack_member_list {

        hardware_version   = "string"
        license_level      = "string"
        license_type       = "string"
        mac_address        = "string"
        pid                = "string"
        priority           = 1
        role               = "string"
        serial_number      = "string"
        software_version   = "string"
        stack_number       = 1
        state              = "string"
        sudi_serial_number = "string"
      }
      stack_ring_protocol      = "string"
      supports_stack_workflows = "false"
      total_member_count       = 1
      valid_license_levels     = ["string"]
    }
    state                = "string"
    sudi_required        = "false"
    tags                 = ["string"]
    user_sudi_serial_nos = ["string"]
    virtual_account_id   = "string"
    workflow_id          = "string"
    workflow_name        = "string"
  }
  run_summary_list {

    details    = "string"
    error_flag = "false"
    history_task_info {

      addn_details {

        key   = "string"
        value = "string"
      }
      name       = "string"
      time_taken = 1
      type       = "string"
      work_item_list {

        command    = "string"
        end_time   = 1
        output_str = "string"
        start_time = 1
        state      = "string"
        time_taken = 1
      }
    }
    timestamp = 1
  }
  system_reset_workflow {

    id               = "string"
    add_to_inventory = "false"
    added_on         = 1
    config_id        = "string"
    curr_task_idx    = 1
    description      = "string"
    end_time         = 1
    exec_time        = 1
    image_id         = "string"
    instance_type    = "string"
    lastupdate_on    = 1
    name             = "string"
    start_time       = 1
    state            = "string"
    tasks {

      curr_work_item_idx = 1
      end_time           = 1
      name               = "string"
      start_time         = 1
      state              = "string"
      task_seq_no        = 1
      time_taken         = 1
      type               = "string"
      work_item_list {

        command    = "string"
        end_time   = 1
        output_str = "string"
        start_time = 1
        state      = "string"
        time_taken = 1
      }
    }
    tenant_id = "string"
    type      = "string"
    use_state = "string"
    version   = 1
  }
  system_workflow {

    id               = "string"
    add_to_inventory = "false"
    added_on         = 1
    config_id        = "string"
    curr_task_idx    = 1
    description      = "string"
    end_time         = 1
    exec_time        = 1
    image_id         = "string"
    instance_type    = "string"
    lastupdate_on    = 1
    name             = "string"
    start_time       = 1
    state            = "string"
    tasks {

      curr_work_item_idx = 1
      end_time           = 1
      name               = "string"
      start_time         = 1
      state              = "string"
      task_seq_no        = 1
      time_taken         = 1
      type               = "string"
      work_item_list {

        command    = "string"
        end_time   = 1
        output_str = "string"
        start_time = 1
        state      = "string"
        time_taken = 1
      }
    }
    tenant_id = "string"
    type      = "string"
    use_state = "string"
    version   = 1
  }
  tenant_id = "string"
  version   = 1
  workflow {

    id               = "string"
    add_to_inventory = "false"
    added_on         = 1
    config_id        = "string"
    curr_task_idx    = 1
    description      = "string"
    end_time         = 1
    exec_time        = 1
    image_id         = "string"
    instance_type    = "string"
    lastupdate_on    = 1
    name             = "string"
    start_time       = 1
    state            = "string"
    tasks {

      curr_work_item_idx = 1
      end_time           = 1
      name               = "string"
      start_time         = 1
      state              = "string"
      task_seq_no        = 1
      time_taken         = 1
      type               = "string"
      work_item_list {

        command    = "string"
        end_time   = 1
        output_str = "string"
        start_time = 1
        state      = "string"
        time_taken = 1
      }
    }
    tenant_id = "string"
    type      = "string"
    use_state = "string"
    version   = 1
  }
  workflow_parameters {

    config_list {

      config_id = "string"
      config_parameters {

        key   = "string"
        value = "string"
      }
    }
    license_level              = "string"
    license_type               = "string"
    top_of_stack_serial_number = "string"
  }
}