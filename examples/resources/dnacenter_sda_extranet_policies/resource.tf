
resource "dnacenter_sda_extranet_policies" "example" {
  provider = dnacenter

  parameters {

    extranet_policy_name             = "string"
    fabric_ids                       = ["string"]
    id                               = "string"
    provider_virtual_network_name    = "string"
    subscriber_virtual_network_names = ["string"]
  }
}

output "dnacenter_sda_extranet_policies_example" {
  value = dnacenter_sda_extranet_policies.example
}