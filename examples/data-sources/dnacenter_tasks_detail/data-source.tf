
data "dnacenter_tasks_detail" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_tasks_detail_example" {
  value = data.dnacenter_tasks_detail.example.item
}
