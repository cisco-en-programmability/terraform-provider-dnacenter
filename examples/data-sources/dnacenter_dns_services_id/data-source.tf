
data "dnacenter_dns_services_id" "example" {
  provider    = dnacenter
  end_time    = 1609459200
  id          = "string"
  start_time  = 1609459200
  xca_lle_rid = "string"
}

output "dnacenter_dns_services_id_example" {
  value = data.dnacenter_dns_services_id.example.item
}
