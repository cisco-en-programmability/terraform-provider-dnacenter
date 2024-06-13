terraform {
  required_providers {
    dnacenter = {
      version = "1.1.32-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_configuration_template_clone" "example" {
  provider = dnacenter

  parameters {
    name        = "hola"
    project_id  = "42f4a526-7498-4593-88f5-d45e22e924ea"
    template_id = "fcfd4d19-99e2-494e-9c6f-0d85cf3094e5"
  }
}
