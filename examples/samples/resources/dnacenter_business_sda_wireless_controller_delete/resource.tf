terraform {
  required_providers {
    dnacenter = {
      version = "1.1.19-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_business_sda_wireless_controller_delete" "example" {
  provider = dnacenter

  parameters {
    device_ipaddress  = "string"
    persistbapioutput = "true"
  }
}
