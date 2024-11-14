
resource "dnacenter_pnp_device_import" "example" {
  provider = dnacenter
  parameters {

    id = "string"
    device_info {

      description            = "string"
      device_sudi_serial_nos = ["string"]
      hostname               = "string"
      mac_address            = "string"
      pid                    = "string"
      serial_number          = "string"
      site_id                = "string"
      stack                  = "false"
      stack_info {

        is_full_ring = "false"
        stack_member_list {

          hardware_version   = "string"
          license_level      = "string"
          license_type       = "string"
          mac_address        = "string"
          pid                = "string"
          priority           = 1.0
          role               = "string"
          serial_number      = "string"
          software_version   = "string"
          stack_number       = 1.0
          state              = "string"
          sudi_serial_number = "string"
        }
        stack_ring_protocol      = "string"
        supports_stack_workflows = "false"
        total_member_count       = 1.0
        valid_license_levels     = ["string"]
      }
      sudi_required        = "false"
      user_mic_numbers     = ["string"]
      user_sudi_serial_nos = ["string"]
      workflow_id          = "string"
      workflow_name        = "string"
    }
  }
}

output "dnacenter_pnp_device_import_example" {
  value = dnacenter_pnp_device_import.example
}