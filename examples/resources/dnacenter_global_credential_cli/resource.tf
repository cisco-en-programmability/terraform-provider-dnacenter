provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_global_credential_cli" "example" {
  provider = dnacenter
  parameters {
    comments           = "string"
    credential_type    = "string"
    description        = "string"
    enable_password    = "string"
    id                 = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    password           = "string"
    username           = "string"
  }
}

output "dnacenter_global_credential_cli_example" {
  value     = dnacenter_global_credential_cli.example
  sensitive = true
}