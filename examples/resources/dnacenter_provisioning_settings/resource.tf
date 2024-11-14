
resource "dnacenter_provisioning_settings" "example" {
  provider = dnacenter

  parameters {

    require_itsm_approval = "false"
    require_preview       = "false"
  }
}

output "dnacenter_provisioning_settings_example" {
  value = dnacenter_provisioning_settings.example
}