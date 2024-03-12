package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)


func exec_Command (command string , args_arr []string) (err error) {

	args := args_arr
	comm_exec := exec.Command(command,args...)
	comm_exec.Stdout = os.Stdout
	comm_exec.Stderr = os.Stderr 
	comm_exec.Stdin = os.Stdin
	
	err =comm_exec.Run()
	
	if err != nil {
		log.Fatal(err)
		return
	}
	
	return nil

}


func main() {

	in_face := flag.String("iface" , "" , "interface you wish to change" )
	new_Mac := flag.String("mac" , "" , "new mac address" )

	flag.Parse()

	exec_Command("sudo" , []string{"ifconfig" , *in_face , "down"})
	exec_Command("sudo" , []string{"ifconfig" , *in_face , "hw" , "ether" , *new_Mac})
	exec_Command("sudo" , []string{"ifconfig" , *in_face ,  "up"})

}