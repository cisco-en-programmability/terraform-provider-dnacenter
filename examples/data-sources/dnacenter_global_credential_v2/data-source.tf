
data "dnacenter_global_credential_v2" "example" {
  provider = dnacenter
}

output "dnacenter_global_credential_v2_example" {
  value = data.dnacenter_global_credential_v2.example.item
}
