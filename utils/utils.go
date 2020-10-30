package utils

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"time"
)

// https://studygolang.com/articles/7196

type UpdateInfo struct  {
	Name			string 	`json:"name"`
	GameCategory	[]int	`json:"gameCategory"`
	GameMaster		string	`json:"gameMaster"`
	GameSpec		string	`json:"gameSpec"`
	Line			int		`json:"line"`
	Bonus			int		`json:"bonus"`
	PlatformProduct	string 	`json:"platform_product"`
}

func LoadUpdateInfoByJson() []UpdateInfo{
	// loading file : json
	f, err := ioutil.ReadFile("./updateInfo.json")
	CheckErr(err)
	// create an object and give it values
	var updateInfos []UpdateInfo
	json.Unmarshal( f, &updateInfos )
	return updateInfos
}

func publicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}

func ConnectWithKeyReturnClient(user, keyPath, host string, port int) (*ssh.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, publicKeyFile(keyPath))
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         10 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	return client, nil
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}