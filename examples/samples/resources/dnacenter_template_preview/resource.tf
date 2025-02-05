terraform {
  required_providers {
    dnacenter = {
      version = "1.3.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_template_preview" "example" {
  provider = dnacenter

  parameters {
    device_id       = "string"
    params          = "string"
    resource_params = ["string"]
    template_id     = "string"
  }
}
