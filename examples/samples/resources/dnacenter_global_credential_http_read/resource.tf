
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.5-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_global_credential_http_read" "example" {
  provider = dnacenter
  parameters {
    secure      = true
    username    = "userTF3"
    password    = "123"
    port        = 23
    description = "This is a test for tf"
    #comments= null
    credential_type = "APP"
    #instance_tenant_id= "6168b750e7a2701a37d64526"
    #instance_uuid= "aed1c6d9-e32d-47b5-a7c4-9e8cb15060a1"
    #id= "aed1c6d9-e32d-47b5-a7c4-9e8cb15060a1"
  }
}

output "dnacenter_global_credential_http_read_example" {
  value     = dnacenter_global_credential_http_read.example
  sensitive = true
}
