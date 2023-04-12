
data "dnacenter_dnac_packages" "example" {
  provider = dnacenter
}

output "dnacenter_dnac_packages_example" {
  value = data.dnacenter_dnac_packages.example.items
}
