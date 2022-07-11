
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.4-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_authentication_import_certificate_p12" "example" {
  provider = dnacenter
 
  parameters {
    file_name     = "string"
    list_of_users = ["string"]
    p12_file_path = "string"
    p12_password  = "******"
    pk_password   = "******"
  }
}