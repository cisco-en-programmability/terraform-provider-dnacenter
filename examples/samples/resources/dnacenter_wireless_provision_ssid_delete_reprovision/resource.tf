
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.12-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_wireless_provision_ssid_delete_reprovision" "example" {
  provider = dnacenter

  parameters {
    managed_aplocations = "string"
    ssid_name           = "string"
  }
}
