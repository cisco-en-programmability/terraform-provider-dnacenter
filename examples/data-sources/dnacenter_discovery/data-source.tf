
data "dnacenter_discovery" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_discovery_example" {
  value = data.dnacenter_discovery.example.item
}
