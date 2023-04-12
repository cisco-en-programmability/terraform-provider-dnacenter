
resource "dnacenter_credential_to_site_by_siteid_create_v2" "example" {
  provider = dnacenter
  parameters {

    cli_id           = "string"
    http_read        = "string"
    http_write       = "string"
    site_id          = "string"
    snmp_v2_read_id  = "string"
    snmp_v2_write_id = "string"
    snmp_v3_id       = "string"
  }
}

output "dnacenter_credential_to_site_by_siteid_create_v2_example" {
  value = dnacenter_credential_to_site_by_siteid_create_v2.example
}