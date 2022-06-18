
resource "dnacenter_nfv_provision_detail" "example" {
  provider = dnacenter
  parameters {

    device_ip = "string"
  }
}

output "dnacenter_nfv_provision_detail_example" {
  value = dnacenter_nfv_provision_detail.example
}