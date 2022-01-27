
resource "dnacenter_event_subscription_syslog" "example" {
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

        connector_type = "string"
      }
    }
    subscription_id = "string"
    version         = "string"
  }
}

output "dnacenter_event_subscription_syslog_example" {
  value = dnacenter_event_subscription_syslog.example
}