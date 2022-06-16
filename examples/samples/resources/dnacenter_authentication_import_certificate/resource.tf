

terraform {
  required_providers {
    dnacenter = {
      version = "0.3.0"
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
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    cert_file_path = "string"
    cert_file_name = "string"
    pk_file_name   = "string"
    list_of_users  = ["string"]
    pk_file_path   = "string"
    pk_password    = "******"
  }
}