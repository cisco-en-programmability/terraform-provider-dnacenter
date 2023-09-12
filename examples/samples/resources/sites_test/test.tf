
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.16-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

#variable "username" {
#  type = string
#}

#variable "password" {
#  type = string
#}

# Configure provider with your Cisco DNA Center SDK credentials


resource "dnacenter_area" "test_area_1" {
  provider = dnacenter
  parameters {

    site {

      area {

        name        = "Test Area 1"
        parent_name = "Global"
      }
    }
    type = "area"
  }
}

resource "dnacenter_building" "test_building_1" {
  provider = dnacenter
  parameters {

    site {

      building {

        address     = "8 AV JEAN MEDECIN 06000 NICE FRANCE"
        name        = "Test Building 1"
        parent_name = dnacenter_area.test_area_1.item[0].site_name_hierarchy
      }
    }
    type    = "building"
  }
}

resource "dnacenter_floor" "test_floor_0" {
  provider = dnacenter
  parameters {

    site {

      floor {

        floor_number = 0
        height       = 10
        length       = 10
        name         = "Test Floor 0"
        parent_name  = "Global/Test Area 1/Test Building 1"
        rf_model     = "Cubes And Walled Offices"
        width        = 10
      }
    }
    type    = "floor"
  }
}


resource "dnacenter_floor" "test_floor_1" {
  provider = dnacenter
  parameters {

    site {

      floor {

        floor_number = 1 
        height       = 10 
        length       = 10
        name         = "Test Floor 1"
        parent_name  = dnacenter_building.test_building_1.item[0].site_name_hierarchy
        rf_model     = "Cubes And Walled Offices"
        width        = 10
      }
    }
    type    = "floor"
  }
}

