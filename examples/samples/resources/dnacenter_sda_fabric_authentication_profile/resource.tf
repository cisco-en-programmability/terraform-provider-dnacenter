terraform {
  required_providers {
    dnacenter = {
      version = "1.1.15-beta"
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
      authenticate_template_name    = "Open Authentication"
      authentication_order          = "dot1x"
      dot1x_to_mab_fallback_timeout = "21"
      number_of_hosts               = "Unlimited"
      site_name_hierarchy           = "Global/San Francisco"
      wake_on_lan                   = "false"
    }
  }
}

output "dnacenter_sda_fabric_authentication_profile_example" {
  value = dnacenter_sda_fabric_authentication_profile.example
}
