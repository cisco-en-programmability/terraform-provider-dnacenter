
resource "dnacenter_interfaces_members_associations_query" "example" {
  provider = dnacenter
  parameters {

    ids = ["string"]
  }
}

output "dnacenter_interfaces_members_associations_query_example" {
  value = dnacenter_interfaces_members_associations_query.example
}