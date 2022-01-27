
resource "dnacenter_sda_fabric_authentication_profile" "example" {
  provider = dnacenter
  parameters {

    authenticate_template_name    = "string"
    authentication_order          = "string"
    dot1x_to_mab_fallback_timeout = "string"
    number_of_hosts               = "string"
    site_name_hierarchy           = "string"
    wake_on_lan                   = "false"
  }
}

output "dnacenter_sda_fabric_authentication_profile_example" {
  value = dnacenter_sda_fabric_authentication_profile.example
}