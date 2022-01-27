
data "dnacenter_pnp_device_reset" "example" {
  provider = dnacenter
  device_reset_list {

    config_list {

      config_id = "string"
      config_parameters {

        key   = "string"
        value = "string"
      }
    }
    device_id                  = "string"
    license_level              = "string"
    license_type               = "string"
    top_of_stack_serial_number = "string"
  }
  project_id  = "string"
  workflow_id = "string"
}