
resource "dnacenter_site_assign_credential" "example" {
  provider = dnacenter
  parameters {

    cli_id            = "string"
    http_read         = "string"
    http_write        = "string"
    site_id           = "string"
    snmp_v2_read_id   = "string"
    snmp_v2_write_id  = "string"
    snmp_v3_id        = "string"
    persistbapioutput = "false"
  }
}

output "dnacenter_site_assign_credential_example" {
  value = dnacenter_site_assign_credential.example
}