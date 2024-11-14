
data "dnacenter_client_detail" "example" {
  provider    = dnacenter
  mac_address = "string"
  timestamp   = 1.0
}

output "dnacenter_client_detail_example" {
  value = data.dnacenter_client_detail.example.item
}
