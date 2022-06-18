
resource "dnacenter_event_subscription_email" "example" {
  provider = dnacenter
  parameters {

    description = "string"
    filter {

      categories = ["string"]
      domains_subdomains {

        domain      = "string"
        sub_domains = ["string"]
      }
      event_ids  = ["string"]
      severities = [1]
      site_ids   = ["string"]
      sources    = ["string"]
      types      = ["string"]
    }
    name = "string"
    subscription_endpoints {

      instance_id = "string"
      subscription_details {

        connector_type     = "string"
        description        = "string"
        from_email_address = "string"
        name               = "string"
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