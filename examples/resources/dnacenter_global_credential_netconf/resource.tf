
resource "dnacenter_global_credential_netconf" "example" {
  provider = dnacenter
  parameters {
    description        = "string"
    comments           = "string"
    credential_type    = "string"
    netconf_port       = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    id                 = "string"
  }
}

output "dnacenter_global_credential_netconf_example" {
  value = dnacenter_global_credential_netconf.example
}