// internal/gateway/registry_consul.go
package gateway

// import (
//   "github.com/hashicorp/consul/api"
//   "log"
// )

// func LookupService(service string) (string, error) {
//   c, err := api.NewClient(api.DefaultConfig())
//   if err != nil {
//     return "", err
//   }

//   srvs, _, err := c.Health().Service(service, "", true, nil)
//   if err != nil {
//     return "", err
//   }
//   if len(srvs) == 0 {
//     return "", fmt.Errorf("no healthy instances")
//   }
//   // pick first or do round-robin
//   addr := fmt.Sprintf("%s:%d", srvs[0].Service.Address, srvs[0].Service.Port)
//   log.Printf("resolved %s -> %s", service, addr)
//   return addr, nil
// }
