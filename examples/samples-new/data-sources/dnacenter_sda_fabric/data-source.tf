terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_sda_fabric" "example" {
  provider    = dnacenter
  fabric_name = "DNAC_Guide_Fabric"
}

output "dnacenter_sda_fabric_example" {
  value = data.dnacenter_sda_fabric.example
}
