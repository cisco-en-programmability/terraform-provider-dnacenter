
resource "dnacenter_sda_multicast" "example" {
  provider = dnacenter
  parameters {

    multicast_method = "string"
    multicast_vn_info {

      external_rp_ip_address = "string"
      ip_pool_name           = "string"
      ssm_group_range        = "string"
      ssm_info               = ["string"]
      ssm_wildcard_mask      = "string"
      virtual_network_name   = "string"
    }
    muticast_type       = "string"
    site_name_hierarchy = "string"
  }
}

output "dnacenter_sda_multicast_example" {
  value = dnacenter_sda_multicast.example
}