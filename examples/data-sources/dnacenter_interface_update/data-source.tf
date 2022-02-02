
data "dnacinterface_update" "example" {
  provider        = dnac
  deployment_mode = "string"
  interface_uuid  = "string"
  admin_status {

    type = "string"
  }
  description {

    type = "string"
  }
  item {

    properties {

      # task_id = ------
      url {

        # type = ------
      }
    }
    # required = [------]
    # type = ------
  }
  vlan_id {

    type = "string"
  }
}