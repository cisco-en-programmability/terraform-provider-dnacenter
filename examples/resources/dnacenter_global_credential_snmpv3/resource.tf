
resource "dnacenter_global_credential_snmpv3" "example" {
  provider = dnacenter
  parameters {
    auth_password    = "string"
    auth_type        = "string"
    comments         = "string"
    credential_type  = "string"
    description      = "string"
    id               = "string"
    instanceTenantId = "string"
    instanceUuid     = "string"
    privacy_password = "string"
    privacy_type     = "string"
    snmp_mode        = "string"
    username         = "string"
  }
}

output "dnacenter_global_credential_snmpv3_example" {
  value = dnacenter_global_credential_snmpv3.example
}