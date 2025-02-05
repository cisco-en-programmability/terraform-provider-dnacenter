
resource "dnacenter_endpoint_analytics_endpoints" "example" {
  provider = dnacenter
  parameters {

    device_type           = "string"
    ep_id                 = "string"
    hardware_manufacturer = "string"
    hardware_model        = "string"
    mac_address           = "string"
  }
}

output "dnacenter_endpoint_analytics_endpoints_example" {
  value = dnacenter_endpoint_analytics_endpoints.example
}
