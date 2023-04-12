
resource "dnacenter_integration_settings_instances_itsm" "example" {
  provider = dnacenter

  parameters {

    data {

      connection_settings {

        auth_password  = "string"
        auth_user_name = "string"
        url            = "string"
      }
    }
    description = "string"
    dyp_name    = "string"
    instance_id = "string"
    name        = "string"
  }
}

output "dnacenter_integration_settings_instances_itsm_example" {
  value = dnacenter_integration_settings_instances_itsm.example
}