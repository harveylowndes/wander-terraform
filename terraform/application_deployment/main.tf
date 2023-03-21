resource "helm_release" "hello_app" {
  chart   = "../../app/chart/hello-app-0.1.0.tgz"
  name    = "hello-app"
  version = "0.1.0"

  set {
    name  = "project"
    value = var.project_id
  }
}