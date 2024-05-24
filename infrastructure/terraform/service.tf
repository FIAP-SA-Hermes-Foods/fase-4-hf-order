resource "kubernetes_service" "hf-order-service" { 
    metadata { 
        name = "hf-order-service"
        namespace = "dev"
    }

    spec { 
        type = "LoadBalancer"
        selector = { 
            app = "hf-order-go-app"
        }
        port { 
            protocol = "TCP"
            name = "hf-order-http-port"
            port = 8080
            target_port = 8080
        }
        port { 
            protocol = "TCP"
            name = "hf-order-rpc-port"
            port = 8070
            target_port = 8070
        }
    }
}
