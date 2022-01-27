
resource "dnacenter_sda_virtual_network_v2" "example" {
  provider = dnacenter
  parameters {

    is_guest_virtual_network = "false"
    scalable_group_names     = ["string"]
    virtual_network_name     = "string"
    virtual_network_type     = "string"
  }
}

output "dnacenter_sda_virtual_network_v2_example" {
  value = dnacenter_sda_virtual_network_v2.example
}