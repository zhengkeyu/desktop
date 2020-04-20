package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	var cliarray = make(map[string]net.Conn)
	listener,err := net.Listen("tcp","localhost:6000")
	if err != nil{
		fmt.Println(err)
	}
	for  {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}else{
			cliarray[conn.RemoteAddr().String()] = conn
			if len(cliarray) != 0{
				for k,v := range cliarray {
					if v.RemoteAddr().String() != conn.RemoteAddr().String(){
						v.Write([]byte("服务器："+conn.RemoteAddr().String()+"已上线"))
					}
						if k != conn.RemoteAddr().String(){
							conn.Write([]byte("服务器："+k+"已上线"))
						}
				}
			}
			go func() {
				for  {
					var climsg = make([]byte,1024)
					count,err := conn.Read(climsg)
					if err != nil && err != io.EOF{
						fmt.Println(err)
						break
					}
					if string(climsg[:count]) == "exit"{
						delete(cliarray,conn.RemoteAddr().String())
						for _,v := range cliarray {
								v.Write([]byte("服务器："+conn.RemoteAddr().String()+"已下线"))
						}
						break
					}else{
						for k,v := range cliarray {
							if k != conn.RemoteAddr().String(){
								_,err := v.Write([]byte(conn.RemoteAddr().String()+"说: "+string(climsg[:count])))
								if err != nil{
									fmt.Println(err)
								}
							}
						}
					}
				}
			}()
		}

	}
}
