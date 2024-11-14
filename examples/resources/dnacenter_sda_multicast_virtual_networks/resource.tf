
resource "dnacenter_sda_multicast_virtual_networks" "example" {
  provider = dnacenter

  parameters {

    fabric_id       = "string"
    id              = "string"
    ip_pool_name    = "string"
    ipv4_ssm_ranges = ["string"]
    multicast_r_ps {

      ipv4_address       = "string"
      ipv4_asm_ranges    = ["string"]
      ipv6_address       = "string"
      ipv6_asm_ranges    = ["string"]
      is_default_v4_rp   = "false"
      is_default_v6_rp   = "false"
      network_device_ids = ["string"]
      rp_device_location = "string"
    }
    virtual_network_name = "string"
  }
}

output "dnacenter_sda_multicast_virtual_networks_example" {
  value = dnacenter_sda_multicast_virtual_networks.example
}