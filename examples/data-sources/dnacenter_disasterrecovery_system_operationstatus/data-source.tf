
data "dnacenter_disasterrecovery_system_operationstatus" "example" {
  provider = dnacenter
}

output "dnacenter_disasterrecovery_system_operationstatus_example" {
  value = data.dnacenter_disasterrecovery_system_operationstatus.example.item
}
