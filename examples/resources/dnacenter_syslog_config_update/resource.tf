
resource "dnacenter_syslog_config_update" "example" {
  provider = dnacenter
  parameters {

    config_id   = "string"
    description = "string"
    host        = "string"
    name        = "string"
    port        = "string"
    protocol    = "string"
  }
}

output "dnacenter_syslog_config_update_example" {
  value = dnacenter_syslog_config_update.example
}