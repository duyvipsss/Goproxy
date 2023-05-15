
package main

import (
	"log"
	"io"
  "net"
)

func handle(conn net.Conn){
  dst, err := net.Dial("tcp", "google.com:80")
  if err != nil{
    log.Fatal("Không thể kết nối đến website")
  }
  defer dst.Close()
  go func(){
    if _, err := io.Copy(conn,dst); err != nil{
      log.Fatal(err)
    }
  }()
  if _, err := io.Copy(dst,conn); err != nil{
      log.Fatal(err)
  }
}

func main(){
  listener, err := net.Listen("tcp", ":90")
  if err != nil{
    log.Fatal("Không thể lắng nghe tại cồng này")
  }
  for{
    conn, err := listener.Accept()
    if err != nil{
      log.Fatal("Thực hiện kết nối với client không thành công")
    }
    go handle(conn)
  }
}
