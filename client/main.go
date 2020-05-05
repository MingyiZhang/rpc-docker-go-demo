package main

import (
  "log"
  "net"
  "net/rpc/jsonrpc"
  "os"

  rpc_demo "rpc-demo"
)

func main() {
  serverHost := os.Getenv("SERVER_HOST")
  conn, err := net.Dial("tcp", serverHost)
  if err != nil {
    panic(err)
  }
  client := jsonrpc.NewClient(conn)

  var result float64
  err = client.Call(
    "DemoService.Div",
    rpc_demo.Args{
      A: 10,
      B: 3,
    },
    &result)
  if err != nil {
    log.Println(err)
  } else {
    log.Println(result)
  }

  err = client.Call(
    "DemoService.Div",
    rpc_demo.Args{
      A: 10,
      B: 0,
    },
    &result)
  if err != nil {
    log.Println(err)
  } else {
    log.Println(result)
  }

}
