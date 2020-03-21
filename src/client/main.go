package main
import (
  "log"
  "api"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)
func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial(":7777", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  c := api.NewDemoClient(conn)
  response, err := c.SayHello(context.Background(), &api.RequestMessage{Message: "Hola!! Handshake :)"})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("[CLIENT] Response from server: %s", response.Message)
}
