
data "dnacenter_file" "example" {
  provider = dnacenter
  dirpath  = "string"
  file_id  = "string"
}

output "dnacenter_file_example" {
  value = data.dnacenter_file.example.item
}
