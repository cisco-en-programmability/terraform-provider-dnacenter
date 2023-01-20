
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.17-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_global_credential_cli" "example" {
  provider = dnacenter
  parameters {
    comments        = "test4_comment"
    credential_type = "GLOBAL"
    description     = "string"
    enable_password = "string"
    #id= "string"
    #instance_tenant_id= "string"
    #instance_uuid= "string"
    password = "string"
    username = "TEST4"
  }
}

output "dnacenter_global_credential_cli_example" {
  value     = dnacenter_global_credential_cli.example
  sensitive = true
}
