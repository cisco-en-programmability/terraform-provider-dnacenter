
resource "dnacenter_syslog_config_create" "example" {
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

output "dnacenter_syslog_config_create_example" {
  value = dnacenter_syslog_config_create.example
}