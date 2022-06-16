
resource "dnacenter_event_webhook_create" "example" {
  provider = dnacenter
  parameters {

    description = "string"
    headers {

      default_value = "string"
      encrypt       = "false"
      name          = "string"
      value         = "string"
    }
    method     = "string"
    name       = "string"
    trust_cert = "false"
    url        = "string"
    webhook_id = "string"
  }
}

output "dnacenter_event_webhook_create_example" {
  value = dnacenter_event_webhook_create.example
}