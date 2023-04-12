
data "dnacenter_event_email_config" "example" {
  provider = dnacenter
}

output "dnacenter_event_email_config_example" {
  value = data.dnacenter_event_email_config.example.items
}
