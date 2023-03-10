variable "namespace" {
  default = "default"
}

variable "service_account" {
  default = "boris"
}

resource "kubernetes_service_account" "this" {
  metadata {
    name      = var.service_account
    namespace = var.namespace
  }
}

resource "kubernetes_secret" "this" {
  metadata {
    annotations = {
      "kubernetes.io/service-account.name" = kubernetes_service_account.this.metadata.0.name
    }
    namespace     = kubernetes_service_account.this.metadata.0.namespace
    generate_name = "${kubernetes_service_account.this.metadata.0.name}-token-"
  }

  type                           = "kubernetes.io/service-account-token"
  wait_for_service_account_token = true
}
