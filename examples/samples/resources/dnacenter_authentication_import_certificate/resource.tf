

terraform {
  required_providers {
    dnacenter = {
      version = "1.0.14-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_authentication_import_certificate" "example" {
  provider = dnacenter

  parameters {
    cert_file_path = "string"
    cert_file_name = "string"
    pk_file_name   = "string"
    list_of_users  = ["string"]
    pk_file_path   = "string"
    pk_password    = "******"
  }
}