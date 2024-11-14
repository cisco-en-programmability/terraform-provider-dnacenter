
resource "dnacenter_sda_multicast_v1_update" "example" {
  provider = dnacenter
  parameters {

    fabric_id        = "string"
    replication_mode = "string"
  }
}

output "dnacenter_sda_multicast_v1_update_example" {
  value = dnacenter_sda_multicast_v1_update.example
}