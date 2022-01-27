terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # hashicorp.com/edu/dnacenter is the local built source change to cisco-en-programmability/dnacenter to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_app_policy_intent_create" "example" {
  provider          = dnacenter
  id                = "02284a52-d136-4956-b5af-ba5daf9165e7"
  instanceId        = 336349
  displayName       = 336349
  instanceCreatedOn = 1634335935102
  instanceUpdatedOn = 1634335935102
  instanceVersion   = 0
  createTime        = 1634335934772
  deployed          = false
  isSeeded          = false
  isStale           = false
  lastUpdateTime    = 1634335934772
  name              = "Application-set-default-policy"
  # namespace= 
  # policy=
  # application=
  provisioningState  = "UNKNOWN"
  qualifier          = "application"
  resourceVersion    = 0
  targetIdList       = []
  type               = "policy"
  cfsChangeInfo      = []
  customProvisions   = []
  deletePolicyStatus = "NONE"
  internal           = true
  isDeleted          = false
  isEnabled          = false
  isScopeStale       = false
  iseReserved        = false
  policyStatus       = "ENABLED"
  priority           = 100
  pushed             = false
  contractList       = []
  exclusiveContract = {
    id                = "b3b295a2-eaad-4d92-b43c-36c1be651eca"
    instanceId        = 342350
    displayName       = 342350
    instanceCreatedOn = 1634335935102
    instanceUpdatedOn = 1634335935102
    instanceVersion   = 0
    clause = [
      {
        id                = "1efb3499-a2be-471c-8f2b-cf10934437e0"
        instanceId        = 337350
        displayName       = 337350
        instanceCreatedOn = 1634335935102
        instanceUpdatedOn = 1634335935102
        instanceVersion   = 0
        priority          = 1
        type              = "BUSINESS_RELEVANCE"
        relevanceLevel    = "BUSINESS_IRRELEVANT"
      }
    ]
  }
  identitySource = {
    id                = "1df01df6-d3ce-4378-a6e2-ef7c9f16d942"
    instanceId        = 14014
    displayName       = 14014
    instanceCreatedOn = 1634257706769
    instanceUpdatedOn = 1634257706769
    instanceVersion   = 0
    state             = INACTIVE
    type              = NBAR
  }
  producer = {
    id                = "14cd6459-a159-4752-a04c-d17b2848c422"
    instanceId        = 343351
    displayName       = "consumer-social-networking"
    instanceCreatedOn = 1634335935102
    instanceUpdatedOn = 1634335935102
    instanceVersion   = 0
    scalableGroup = [
      {
        idRef = "76bab4ae-04c2-490a-a9c2-f7a284d5ba85"
      }
    ]
  }
}