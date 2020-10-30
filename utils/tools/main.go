package main

import (
	"fmt"
	"uploadGaming/utils"
)





func main() {


	var updateInfos []utils.UpdateInfo
	updateInfos = utils.LoadUpdateInfoByJson()

	for i, element :=range updateInfos {
		fmt.Printf("%d, %s\n", i, element.Name)
	}

	sshClient, err := utils.ConnectWithKeyReturnClient("sam", "./sshKeys/vmHome/id_rsa", "172.20.10.3", 22)
	utils.CheckErr(err)
	defer sshClient.Close()

	sshSession, err := sshClient.NewSession()
	utils.CheckErr(err)
	defer sshSession.Close()

	combo, err := sshSession.CombinedOutput("ls -al")
	utils.CheckErr(err)
	fmt.Print(string(combo))



}
