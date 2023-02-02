
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.18-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_pnp_device" "example" {
  provider = dnacenter
  # lifecycle {
  #   create_before_destroy = true
  # }
  parameters {

    # id = "61f1c861264f342e4fa1a78e"
    device_info {
      agent_type         = "POSIX"
      cm_state           = "Not Contacted"
      onb_state          = "Not Contacted"
      populate_inventory = "false"
      reload_requested   = "true"
      serial_number      = "FLM2213W05R"
      stack              = "false"
      # state= "Unclaimed"
      state         = "Planned"
      sudi_required = "false"
      name          = "FLM2213W05R"
      last_contact  = 1
    }
  }
}

output "dnacenter_pnp_device_example" {
  value = dnacenter_pnp_device.example
}
