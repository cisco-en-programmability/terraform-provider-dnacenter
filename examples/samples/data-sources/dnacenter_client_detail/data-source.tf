terraform {
  required_providers {
    dnacenter = {
      version = "0.1.0-beta.1"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_client_detail" "example" {
  provider    = dnacenter
  mac_address = "string"
  timestamp   = "string"
}

output "dnacenter_client_detail_example" {
  value = data.dnacenter_client_detail.example.item
}