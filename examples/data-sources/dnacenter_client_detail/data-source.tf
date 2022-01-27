
data "dnacenter_client_detail" "example" {
  provider    = dnacenter
  mac_address = "string"
  timestamp   = "string"
}

output "dnacenter_client_detail_example" {
  value = data.dnacenter_client_detail.example.item
}
