terraform {
  required_version = ">= 1.4.6"

  required_providers {
    gitea = {
      source  = "terraform.local/local/gitea"
      version = "1.1.0"
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
