package main

//Dependencies
import (
	"flag"
	"log"
	"os"
	"os/exec"
)

//Execution Command
func exec_Command (command string , args_arr []string) (err error) {
	//to handle many arguments
	args := args_arr
	//to execute whole command
	comm_exec := exec.Command(command,args...)
	//Program Output
	comm_exec.Stdout = os.Stdout
	//Program Error
	comm_exec.Stderr = os.Stderr 
	//Program input
	comm_exec.Stdin = os.Stdin
	//running command
	err =comm_exec.Run()
	//check_for_errors
	if err != nil {
		log.Fatal(err)
		return
	}
	
	return nil

}


func main() {
	//handle arguments
	in_face := flag.String("iface" , "" , "interface you wish to change" )
	new_Mac := flag.String("mac" , "" , "new mac address" )
	//parsing arguments
	flag.Parse()
	//shutting down the_interface , changing_new mac address and then enable it
	exec_Command("sudo" , []string{"ifconfig" , *in_face , "down"})
	exec_Command("sudo" , []string{"ifconfig" , *in_face , "hw" , "ether" , *new_Mac})
	exec_Command("sudo" , []string{"ifconfig" , *in_face ,  "up"})

}
