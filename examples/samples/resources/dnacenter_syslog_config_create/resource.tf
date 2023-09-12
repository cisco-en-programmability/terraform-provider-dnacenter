terraform {
  required_providers {
    dnacenter = {
      version = "1.1.16-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_syslog_config_create" "example" {
  provider = dnacenter
  parameters {

    //config_id = "string"
    description = "string"
    host        = "string"
    name        = "string"
    port        = "string"
    protocol    = "string"
  }
}

output "dnacenter_syslog_config_create_example" {
  value = dnacenter_syslog_config_create.example
}
