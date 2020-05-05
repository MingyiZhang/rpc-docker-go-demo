package main

import (
  "log"
  "net"
  "net/rpc"
  "net/rpc/jsonrpc"

  rpc_demo "rpc-demo"
)

func main() {
  rpc.Register(rpc_demo.DemoService{})
  listener, err := net.Listen("tcp", ":1234")
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
