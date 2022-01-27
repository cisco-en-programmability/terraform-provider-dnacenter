
data "dnacenter_file_namespace_files" "example" {
  provider   = dnacenter
  name_space = "string"
}

output "dnacenter_file_namespace_files_example" {
  value = data.dnacenter_file_namespace_files.example.items
}
