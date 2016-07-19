package main

import (
  client "k8s.io/kubernetes/pkg/client/unversioned"
  "log"
  "os"
  "fmt"
  "k8s.io/kubernetes/pkg/api"
  "k8s.io/kubernetes/pkg/client/restclient"
)

func DisplayMap(m map[string]int ){
  // Display the map keys/value
  for n, v := range m {
    fmt.Printf("      %s: %d\n", n, v)
  }
}

func checkErr(err error, m string){
  if err != nil {
      log.Fatalln(m, err)
  }
}

func main() {
  log.SetOutput(os.Stdout)

  token := "djxCNG1qdrl47vzrE8tZw4nrb2uJzael"

  config := restclient.Config{
      Host:        "https://kubetest.com:8443",
      BearerToken: string(token),
      Insecure: true,
  }
  c, err := client.New(&config)
  checkErr(err, "Can't connect to Kubernetes API:")

  log.Printf("%s", config.Host)


    // Nodes
    nodes, err := c.Nodes().List(api.ListOptions{})
    checkErr(err, "Can't get nodes:")

    available_nodes := 0
    for _, node := range nodes.Items {
        // Name
        fmt.Printf("Name: %s\n", node.Name)

        // Labels
        fmt.Println("Labels:")
        for k, v := range node.Labels {
            fmt.Printf("   %s=%s\n", k,v)
        }

        //fmt.Printf("Status: %v\n", node.Status.NodeInfo.OSImage)
        fmt.Println("Status:")
        for _, c := range node.Status.Conditions {
          fmt.Printf("   %s %s\n", c.Type, c.Status)
          if ((c.Type == "Ready") && (c.Status == "True")) {
              available_nodes += 1
          }

        }
        fmt.Println("---------------------")
    }

    fmt.Printf("Nodes available: %d/%d\n", available_nodes, len(nodes.Items))

    // Namespaces
    fmt.Println("Namespaces:")
    ns, err := c.Namespaces().List(api.ListOptions{})
    checkErr(err, "Can't get namespaces:")

    for _, n := range ns.Items {
      fmt.Printf("   %s\n", n.Name)
    }


    // Pods
    fmt.Println("Pods:")

    pods, err := c.Pods(api.NamespaceAll).List(api.ListOptions{})
    checkErr(err, "Can't get pods:")

    // Pods by namespaces
    pna := make(map[string]int)
    // Pods by nodes
    pno := make(map[string]int)
    for _, p := range pods.Items {
      pna[p.Namespace]+=1
      pno[p.Spec.NodeName]+=1
    }
    fmt.Printf("   Total: %d\n", len(pods.Items))
    // Number pods by namespaces
    fmt.Println("   Per namespaces:")
    DisplayMap(pna)

    // Number pods by namespaces
    fmt.Println("   Per nodes:")
    DisplayMap(pno)

    // Services
    svc, err := c.Services(api.NamespaceAll).List(api.ListOptions{})
    checkErr(err, "Can't get service:")

    fmt.Printf("Services:\n   Total: %d\n", len(svc.Items))
    // svc by namespaces
    sn := make(map[string]int)
    for _, s := range svc.Items {
        sn[s.Namespace]+=1
        //fmt.Printf("%s\n", s.Name)
    }
    fmt.Println("   Per namespaces:")
    DisplayMap(sn)

    // Route / Ingress
    ing, err := c.Ingress(api.NamespaceAll).List(api.ListOptions{})
    checkErr(err, "Can't get ingress:")
    fmt.Printf("Ingress:\n   Total: %d\n", len(ing.Items))

    in := make(map[string]int)
    for _, i := range ing.Items {
        in[i.Namespace]+=1
    }

    fmt.Println("   Per namespaces:")
    DisplayMap(in)

}
