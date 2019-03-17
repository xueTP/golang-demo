package main

import (
	"crypto/tls"
	"log"
	"io/ioutil"
	"crypto/x509"
	"fmt"
)

func main() {
	cret, err := tls.LoadX509KeyPair("tls/client/client.pem", "tls/client/client.key")
	if err != nil {
		log.Println(err)
		return
	}

	clientPem, err := ioutil.ReadFile("tls/client/client.pem")
	if err != nil {
		panic(err)
	}

	clientPemPool := x509.NewCertPool()
	ok := clientPemPool.AppendCertsFromPEM(clientPem)
	if !ok {
		panic("unable to append client pem to pool")
	}
	config := &tls.Config{
		RootCAs: clientPemPool,
		Certificates: []tls.Certificate{cret},
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:9995", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("你好\n"))
	if err != nil {
		log.Println(n, err)
		return
	}
	buff := make([]byte, 100)
	n, err = conn.Read(buff)
	if err != nil {
		log.Println(n, err)
		return
	}
	fmt.Println(string(buff[:n]))
}
