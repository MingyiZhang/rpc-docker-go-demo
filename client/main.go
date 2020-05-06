package main

import (
  "flag"
  "log"
  "net"
  "net/rpc/jsonrpc"

  "rpc-demo/service"
)

var host = flag.String("host", "", "the host to connect")

func main() {
  flag.Parse()
  if *host == "" {
    log.Fatal("must specify host")
  }

  conn, err := net.Dial("tcp", *host)
  if err != nil {
    panic(err)
  }
  client := jsonrpc.NewClient(conn)

  var result float64
  err = client.Call(
    "DemoService.Div",
    service.Args{
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
    service.Args{
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
