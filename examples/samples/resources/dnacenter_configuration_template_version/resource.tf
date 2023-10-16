terraform {
  required_providers {
    dnacenter = {
      version = "1.1.23-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_configuration_template_version" "example" {
  provider = dnacenter
  parameters {
    # comments = "string"
    template_id = "524d6bfd-45df-4399-9bbf-3e6f4006b009"
  }
}

output "dnacenter_configuration_template_version_example" {
  value = dnacenter_configuration_template_version.example
}
