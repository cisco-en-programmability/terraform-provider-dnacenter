
resource "dnacenter_ipam_server_setting" "example" {
  provider = dnacenter

  parameters {

    password    = "******"
    provider    = "string"
    server_name = "string"
    server_url  = "string"
    sync_view   = "false"
    user_name   = "string"
    view        = "string"
  }
}

output "dnacenter_ipam_server_setting_example" {
  value = dnacenter_ipam_server_setting.example
}