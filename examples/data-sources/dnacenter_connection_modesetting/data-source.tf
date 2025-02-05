
data "dnacenter_connection_modesetting" "example" {
  provider = dnacenter
}

output "dnacenter_connection_modesetting_example" {
  value = data.dnacenter_connection_modesetting.example.item
}
