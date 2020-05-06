package main

import (
  "flag"
  "fmt"
  "log"
  "net"
  "net/rpc"
  "net/rpc/jsonrpc"

  "rpc-demo/service"
)

var port = flag.Int("port", 0, "the port to listening on")

func main() {
  flag.Parse()
  if *port == 0 {
    log.Fatal("must specify a port")
  }
  err := rpc.Register(service.DemoService{})
  if err != nil {
    panic(err)
  }
  listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
  if err != nil {
    panic(err)
  }
  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Printf("accept error: %v", err)
      continue
    }
    go jsonrpc.ServeConn(conn)
  }
}
