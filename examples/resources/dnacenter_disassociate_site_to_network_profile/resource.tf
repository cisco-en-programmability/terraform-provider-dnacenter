
resource "dnacenter_disassociate_site_to_network_profile" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    network_profile_id = "string"
    site_id            = "string"
  }
}