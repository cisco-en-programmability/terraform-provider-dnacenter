terraform {
  required_providers {
    dnacenter = {
      version = "1.1.3-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_sda_virtual_network" "example" {
  provider = dnacenter
  parameters {
    payload {
      site_name_hierarchy  = "Global/New Jersey/MurrayHill/test/TestFloor"
      virtual_network_name = "GUEST_VN"
    }
  }
}

# resource "dnacenter_sda_virtual_network" "example2" {
#   provider = dnacenter
#   depends_on = [
#     dnacenter_sda_virtual_network.example
#   ]
#   parameters {
#     payload {
#       site_name_hierarchy  = "Global/New Jersey/MurrayHill/test/TestFloor"
#       virtual_network_name = " TEST_VNs"
#     }
#   }
# }

output "dnacenter_sda_virtual_network_example" {
  value = dnacenter_sda_virtual_network.example
}