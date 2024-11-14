
resource "dnacenter_sda_layer3_virtual_networks" "example" {
  provider = dnacenter

  parameters {

    anchored_site_id     = "string"
    fabric_ids           = ["string"]
    id                   = "string"
    virtual_network_name = "string"
  }
}

output "dnacenter_sda_layer3_virtual_networks_example" {
  value = dnacenter_sda_layer3_virtual_networks.example
}