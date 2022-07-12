
resource "dnacenter_event_subscription_rest" "example" {
  provider = dnacenter
  parameters {
    payload {
      description = "string"
      filter {

        categories = ["string"]
        domains_subdomains {

          domain      = "string"
          sub_domains = ["string"]
        }
        event_ids  = ["string"]
        severities = ["string"]
        site_ids   = ["string"]
        sources    = ["string"]
        types      = ["string"]
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
}

output "dnacenter_event_subscription_rest_example" {
  value = dnacenter_event_subscription_rest.example
}