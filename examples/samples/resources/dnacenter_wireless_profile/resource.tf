
terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}


resource "dnacenter_wireless_profile" "example" {
  provider = dnacenter

  parameters {
    profile_details {
      name  = "Test22"
      sites = ["Global/CR"]



      ssid_details {
        enable_fabric = "true"
        flex_connect {
          enable_flex_connect = "false"
          local_to_vlan       = 0
        }
        interface_name = "management"
        name           = "BTest22"
        type           = "string"
      }
      ssid_details {
        enable_fabric = "true"
        flex_connect {
          enable_flex_connect = "false"
          local_to_vlan       = 0
        }
        interface_name = "management"
        name           = "ATest222"
        type           = "eduroam"
      }
      ssid_details {
        enable_fabric = "true"
        flex_connect {
          enable_flex_connect = "false"
          local_to_vlan       = 0
        }
        interface_name = "management"
        name           = "new"
        type           = "string"
      }
    }
  }
}


# output "dnacenter_wireless_profile_example" {
#   value = dnacenter_wireless_profile.example
# }
