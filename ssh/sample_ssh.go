package main

import (
	"log"
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"flag"
)

func main() {
	var user string
	var addr string
	var key string
	flag.StringVar(&user, "user", "", "ssh username")
	flag.StringVar(&addr, "host", "", "ssh host with port")
	flag.StringVar(&key, "key", "", "path to ssh private key")
	flag.Parse()

	content, err := ioutil.ReadFile(key)
	if err != nil {
		log.Fatal("read id_rsa failed", err)
	}
	signer, err := ssh.ParsePrivateKey(content)
	if err != nil {
		log.Fatal("parse id_rsa failed", err)
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/bin/ls /"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
