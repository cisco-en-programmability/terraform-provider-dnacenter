terraform {
  required_providers {
    dnacenter = {
      version = "1.1.17-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_configuration_template_import_project" "example" {
  provider = dnacenter

  parameters {
    do_version = "false"
  }
}
