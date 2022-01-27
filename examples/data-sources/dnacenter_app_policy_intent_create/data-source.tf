
data "dnacenter_app_policy_intent_create" "example" {
  provider = dnacenter
  create_list {

    advanced_policy_scope {

      advanced_policy_scope_element {

        group_id = ["string"]
        ssid     = ["string"]
      }
      name = "string"
    }
    consumer {

      scalable_group {

        id_ref = "string"
      }
    }
    contract {

      id_ref = "string"
    }
    delete_policy_status = "string"
    exclusive_contract {

      clause {

        device_removal_behavior = "string"
        host_tracking_enabled   = "false"
        relevance_level         = "string"
        type                    = "string"
      }
    }
    name         = "string"
    policy_scope = "string"
    priority     = "string"
    producer {

      scalable_group {

        id_ref = "string"
      }
    }
  }
  delete_list = ["string"]
  update_list {

    advanced_policy_scope {

      advanced_policy_scope_element {

        group_id = ["string"]
        id       = "string"
        ssid     = ["string"]
      }
      id   = "string"
      name = "string"
    }
    consumer {

      id = "string"
      scalable_group {

        id_ref = "string"
      }
    }
    contract {

      id_ref = "string"
    }
    delete_policy_status = "string"
    exclusive_contract {

      clause {

        device_removal_behavior = "string"
        host_tracking_enabled   = "false"
        id                      = "string"
        relevance_level         = "string"
        type                    = "string"
      }
      id = "string"
    }
    id           = "string"
    name         = "string"
    policy_scope = "string"
    priority     = "string"
    producer {

      id = "string"
      scalable_group {

        id_ref = "string"
      }
    }
  }
}