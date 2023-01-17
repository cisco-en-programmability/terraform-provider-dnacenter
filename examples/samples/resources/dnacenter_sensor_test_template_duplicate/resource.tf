terraform {
  required_providers {
    dnacenter = {
      version = "1.0.16-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}


resource "dnacenter_sensor_test_template_duplicate" "example" {
  provider = dnacenter

  parameters {
    new_template_name = "string"
    template_name     = "string"
  }
}