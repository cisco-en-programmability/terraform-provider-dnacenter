
resource "dnacenter_sda_port_channels" "example" {
  provider = dnacenter

  parameters {

    connected_device_type = "string"
    description           = "string"
    fabric_id             = "string"
    id                    = "string"
    interface_names       = ["string"]
    network_device_id     = "string"
    port_channel_name     = "string"
    protocol              = "string"
  }
}

output "dnacenter_sda_port_channels_example" {
  value = dnacenter_sda_port_channels.example
}