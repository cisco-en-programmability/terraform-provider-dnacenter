
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.11-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_site_assign_credential" "example" {
  provider = dnacenter

  parameters {
    site_id          = "string"
    cli_id           = "string"
    http_read        = "string"
    http_write       = "string"
    snmp_v2_read_id  = "string"
    snmp_v2_write_id = "string"
    snmp_v3_id       = "string"
  }
}