
data "dnacenter_snmpv2_read_community_credential_create" "example" {
  provider        = dnacenter
  comments        = "string"
  credential_type = "string"
  description     = "string"
  read_community  = "string"
}