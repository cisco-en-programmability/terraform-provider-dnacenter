terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source   = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

resource "dna_sda_fabric_ip_pool_in_vn" "ip_pool_guest" {
  provider   = dnacenter
  virtual_network_name = "Test_VN"
  ip_pool_name = "Test_Pool2" 
  traffic_type = "Data"
  authentication_policy_name = "Test_policy"
  scalable_group_name = ""
#   is_l2_flooding_enabled = false
#   is_this_critical_pool = false
#   pool_type = "Generic"
}

output "ip_pool_guest" {
    value = dna_sda_fabric_ip_pool_in_vn.ip_pool_guest
}