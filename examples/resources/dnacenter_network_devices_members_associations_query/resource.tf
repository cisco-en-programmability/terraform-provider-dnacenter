
resource "dnacenter_network_devices_members_associations_query" "example" {
  provider = dnacenter
  parameters {

    ids = ["string"]
  }
}

output "dnacenter_network_devices_members_associations_query_example" {
  value = dnacenter_network_devices_members_associations_query.example
}