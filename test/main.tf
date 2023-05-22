terraform {
  required_version = ">= 1.4.6"

  required_providers {
    gitea = {
      source  = "terraform.local/local/gitea"
      version = "1.1.2"
    }
  }
}

variable "base_url" {
  type = string
}

variable "token" {
  type = string
}

provider "gitea" {
  base_url       = var.base_url
  token          = var.token
  debug          = true
  debug_log_path = "/mnt/d/trace.txt"
}

resource "gitea_org" "testorg" {
  name = "test_org"
}

# resource "gitea_team" "testteam" {
#   name = "test_team"
#   organization = gitea_org.testorg.name
# }

# resource "gitea_repo" "testrepo" {
#   owner = gitea_org.testorg.name
#   name = "test_repo"
# }

# resource "gitea_branch_protection" "test_bp" {
#   owner = gitea_org.testorg.name
#   repo = gitea_repo.testrepo.name
#   branch_name = "master"
# }