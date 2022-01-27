
data "dnacenter_task_tree" "example" {
  provider = dnacenter
  task_id  = "string"
}

output "dnacenter_task_tree_example" {
  value = data.dnacenter_task_tree.example.items
}
