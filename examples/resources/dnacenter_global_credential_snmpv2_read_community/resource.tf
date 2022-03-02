
resource "dnacenter_global_credential_snmpv2_read_community" "example" {
  provider = dnacenter
  parameters {
    description     = "string"
    comments        = "string"
    credential_type = "string"
    read_community  = "string"
  }
}

output "dnacenter_global_credential_snmpv2_read_community_example" {
  value = dnacenter_global_credential_snmpv2_read_community.example
}