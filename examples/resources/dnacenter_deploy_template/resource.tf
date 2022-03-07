
resource "dnacenter_deploy_template" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    force_push_template=  "false"
    is_composite=         "false"
    member_template_deployment_info= []
    target_info {
        host_name=              "string"
        id=                     "string"
        params=                 ["string"]
        resource_params=        ["string"]
        type=                   "string"
        versioned_template_id=  "string"
      }
    template_id=         "string"
  }
}

output "dnacenter_deploy_template_example" {
  value = dnacenter_deploy_template.example
}

