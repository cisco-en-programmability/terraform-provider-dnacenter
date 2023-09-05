terraform {
  required_providers {
    dnacenter = {
      version = "1.1.13-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacdnacenter_network_device_sync" "example" {
  provider   = dnacenter
  force_sync = "false"
}
