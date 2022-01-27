
data "dnacenter_disasterrecovery_system_status" "example" {
  provider = dnacenter
}

output "dnacenter_disasterrecovery_system_status_example" {
  value = data.dnacenter_disasterrecovery_system_status.example.item
}
