
data "dnacenter_nfv_provision_detail" "example" {
  provider  = dnacenter
  device_ip = "string"
}

output "dnacenter_nfv_provision_detail_example" {
  value = data.dnacenter_nfv_provision_detail.example.item
}
