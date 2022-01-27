
data "dnacenter_disassociate_site_to_network_profile" "example" {
  provider           = dnacenter
  network_profile_id = "string"
  site_id            = "string"
}