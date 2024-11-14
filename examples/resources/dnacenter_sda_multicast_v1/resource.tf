
resource "dnacenter_sda_multicast_v1" "example" {
  provider = dnacenter

  parameters {

    fabric_id        = "string"
    replication_mode = "string"
  }
}

output "dnacenter_sda_multicast_v1_example" {
  value = dnacenter_sda_multicast_v1.example
}