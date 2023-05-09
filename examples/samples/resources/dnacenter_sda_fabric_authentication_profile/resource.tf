terraform {
  required_providers {
    dnacenter = {
      version = "1.1.4-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_sda_fabric_authentication_profile" "example" {
  provider = dnacenter
  parameters {
    payload {
      authenticate_template_name    = "string"
      authentication_order          = "string"
      dot1x_to_mab_fallback_timeout = "string"
      number_of_hosts               = "string"
      site_name_hierarchy           = "string"
      wake_on_lan                   = "false"
    }
  }
}

output "dnacenter_sda_fabric_authentication_profile_example" {
  value = dnacenter_sda_fabric_authentication_profile.example
}