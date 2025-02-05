terraform {
  required_providers {
    dnacenter = {
      version = "1.3.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_sensor_test_run" "example" {
  provider      = dnacenter
  template_name = "Test"
}
