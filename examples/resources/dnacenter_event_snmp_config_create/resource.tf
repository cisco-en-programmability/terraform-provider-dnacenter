
resource "dnacenter_event_snmp_config_create" "example" {
  provider = dnacenter
  parameters {

    auth_password     = "string"
    community         = "string"
    description       = "string"
    ip_address        = "string"
    name              = "string"
    port              = "string"
    privacy_password  = "string"
    snmp_auth_type    = "string"
    snmp_mode         = "string"
    snmp_privacy_type = "string"
    snmp_version      = "string"
    user_name         = "string"
  }
}

output "dnacenter_event_snmp_config_create_example" {
  value = dnacenter_event_snmp_config_create.example
}