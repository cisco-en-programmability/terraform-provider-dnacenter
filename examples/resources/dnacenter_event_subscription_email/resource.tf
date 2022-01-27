
resource "dnacenter_event_subscription_email" "example" {
  provider = dnacenter
  parameters {

    description = "string"
    filter {

      event_ids = ["string"]
    }
    name = "string"
    subscription_endpoints {

      instance_id = "string"
      subscription_details {

        connector_type     = "string"
        from_email_address = "string"
        subject            = "string"
        to_email_addresses = ["string"]
      }
    }
    subscription_id = "string"
    version         = "string"
  }
}

output "dnacenter_event_subscription_email_example" {
  value = dnacenter_event_subscription_email.example
}