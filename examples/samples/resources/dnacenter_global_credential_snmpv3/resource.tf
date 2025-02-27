
terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_global_credential_snmpv3" "example" {
  provider = dnacenter
  parameters {
    auth_password   = "12345678"
    auth_type       = "SHA"
    comments        = "Test"
    credential_type = "APP"
    description     = "Description 3"
    #id= "string"
    #instanceTenantId= "string"
    #instanceUuid= "string"
    privacy_password = "privacy_password_test"
    privacy_type     = "AES128"
    snmp_mode        = "AUTHPRIV"
    username         = "Global_credential_test3"
  }
}

output "dnacenter_global_credential_snmpv3_example" {
  value = dnacenter_global_credential_snmpv3.example
}
