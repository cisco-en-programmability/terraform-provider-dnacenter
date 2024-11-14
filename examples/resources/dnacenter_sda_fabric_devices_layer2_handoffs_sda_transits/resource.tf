
resource "dnacenter_sda_fabric_devices_layer2_handoffs_sda_transits" "example" {
  provider = dnacenter

  parameters {

    affinity_id_decider               = 1
    affinity_id_prime                 = 1
    connected_to_internet             = "false"
    fabric_id                         = "string"
    is_multicast_over_transit_enabled = "false"
    network_device_id                 = "string"
    transit_network_id                = "string"
  }
}

output "dnacenter_sda_fabric_devices_layer2_handoffs_sda_transits_example" {
  value = dnacenter_sda_fabric_devices_layer2_handoffs_sda_transits.example
}