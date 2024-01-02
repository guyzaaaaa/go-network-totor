package main

import(
	"fmt"
	"io"
	"net"
    "os" 
	
	  
)

func handConnection(conn net.Conn){
	defer conn.Close()
	buffer := make([]byte, 1024)
	fileNameBuffer := make([]byte, 64)
	n , err := conn.Read(fileNameBuffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := string(fileNameBuffer[:n])
	fmt.Println("Receive File Name:" , fileName)


	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //close file brfore exit

	for{
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF{
				fmt.Println("Transfer Complete")
			} else{
				fmt.Println(err)
				
			}
			return
		}
		file.Write(buffer[:n])
	}
}

func main() {
	//create server
listener, err := net.Listen("tcp", ":5000")
if err != nil {
	fmt.Println(err)
	return
}
defer listener.Close()
fmt.Println("Server is listening on port 5000")
for{
	conn, err := listener.Accept()
	if err != nil{
		fmt.Println(err)
		continue
	}
	fmt.Println("Client Connected", conn.RemoteAddr())
	go handConnection(conn)
}
}