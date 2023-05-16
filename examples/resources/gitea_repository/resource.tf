resource "gitea_repository" "repo" {
  name         = "repo-name"
  owner        = "org owner"
  allow_rebase = false
  has_issues   = true
  template     = true
}