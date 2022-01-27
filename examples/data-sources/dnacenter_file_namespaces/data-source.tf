
data "dnacenter_file_namespaces" "example" {
  provider = dnacenter
}

output "dnacenter_file_namespaces_example" {
  value = data.dnacenter_file_namespaces.example.items
}
