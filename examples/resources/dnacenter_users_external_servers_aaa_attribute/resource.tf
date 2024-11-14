
resource "dnacenter_users_external_servers_aaa_attribute" "example" {
  provider = dnacenter

}

output "dnacenter_users_external_servers_aaa_attribute_example" {
  value = dnacenter_users_external_servers_aaa_attribute.example
}