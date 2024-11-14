
resource "dnacenter_event_syslog_config" "example" {
  provider = dnacenter

  parameters {

    config_id   = "string"
    description = "string"
    host        = "string"
    name        = "string"
    port        = 1
    protocol    = "string"
  }
}

output "dnacenter_event_syslog_config_example" {
  value = dnacenter_event_syslog_config.example
}