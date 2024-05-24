resource "kubernetes_deployment" "hf-order-deployment" {

    metadata {  
        name = "hf-order-deployment"
        namespace = "dev"
    }
    spec {  
        selector { 
            match_labels = {
              app = "hf-order-go-app"
            }
        }
        template { 
            metadata { 
                labels = {  
                    app = "hf-order-go-app"
                }
            }
            spec { 
                container { 
                    name = "hf-order-go-app-http"
                    image = "${data.external.env_vars.result.image_api_http_url}:${data.external.env_vars.result.image_tag}"
                    image_pull_policy = "Always"
                    env_from { 
                        secret_ref { 
                            name = "hf-deploy-secret"
                        }
                    }
                    port { 
                        container_port = 8080
                    }
                }
                container { 
                    name = "hf-order-go-app-rpc"
                    image = "${data.external.env_vars.result.image_api_rpc_url}:${data.external.env_vars.result.image_tag}"
                    image_pull_policy = "Always"
                    env_from { 
                        secret_ref { 
                            name = "hf-deploy-secret"
                        }
                    }
                    port { 
                        container_port = 8080
                    }
                }
            image_pull_secrets { 
                    name = "hfregcred"
                }
            }
        }
    }
}

