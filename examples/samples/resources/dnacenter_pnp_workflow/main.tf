

terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_pnp_workflow" "example" {
  provider = dnacenter
  parameters {

    id               = "string"
    add_to_inventory = "false"
    added_on         = 1
    config_id        = "string"
    curr_task_idx    = 1
    description      = "string"
    end_time         = 1
    exec_time        = 1
    image_id         = "string"
    instance_type    = "string"
    lastupdate_on    = 1
    name             = "string"
    start_time       = 1
    state            = "string"
    tasks {

      curr_work_item_idx = 1
      end_time           = 1
      name               = "string"
      start_time         = 1
      state              = "string"
      task_seq_no        = 1
      time_taken         = 1
      type               = "string"
      work_item_list {

        command    = "string"
        end_time   = 1
        output_str = "string"
        start_time = 1
        state      = "string"
        time_taken = 1
      }
    }
    tenant_id = "string"
    type      = "string"
    use_state = "string"
    version   = 1
  }
}

output "dnacenter_pnp_workflow_example" {
  value = dnacenter_pnp_workflow.example
}
