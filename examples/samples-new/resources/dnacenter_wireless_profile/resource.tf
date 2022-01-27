
terraform {
  required_providers {
  dnacenter = {
    version = "0.0.3"
    source  = "hashicorp.com/edu/dnacenter"
    # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
  }
  }
}

resource "dnacenter_wireless_profile" "example" {
  provider = dnacenter
  parameters {

    profile_details {
      name  = "Test2"
      sites = ["Global/CR"]
      ssid_details {
        enable_fabric = "true"
        flex_connect {
          enable_flex_connect = "false"
          local_to_vlan       = 0
        }
        interface_name = "management"
        name           = "Test2"
        type           = "string"
      }
    }
    #wireless_profile_name = "string"
  }
}

output "dnacenter_wireless_profile_example" {
  value = dnacenter_wireless_profile.example
}