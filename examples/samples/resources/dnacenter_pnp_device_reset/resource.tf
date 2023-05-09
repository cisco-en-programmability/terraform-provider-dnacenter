
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.4-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_pnp_device_reset" "example" {
  provider = dnacenter

  parameters {
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
      project_id                 = "string"
      workflow_id                = "string"
    }

  }
}