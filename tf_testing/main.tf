terraform {
  required_providers {
    okta = {
      source  = "okta/okta"
      version = "~> 3.20"
    }
  }
}

variable "okta_api_token" {
  type        = string
  description = "this is the token for okta"
}

variable "okta_org_name" {
  type        = string
  description = "okta org name"
}

# # Configure the Okta Provider
provider "okta" {
  org_name  = var.okta_org_name
  base_url  = "okta.com"
  api_token = var.okta_api_token
}

resource "okta_user" "test" {
  first_name = "TestA3324"
  last_name  = "Smith"
  login      = "testAcc-replace_with_uuid@example.com"
  email      = "testAcc-replace_with_uuid@example.com"
  department = "graph, sec, other"
}
