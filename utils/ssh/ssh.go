package ssh

import "golang.org/x/crypto/ssh"


//创建ssh客户端的结构体
type SSHClient struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	IpAddress string  `json:"ipaddress"`
	Port      int     `json:"port"`
	Session   *ssh.Session
	Client    *ssh.Client
	channel   ssh.Channel

}

//创建新的ssh客户端时,默认用户名为root,端口为22
func NewSSHClient() SSHClient{
	client := SSHClient{}
	client.Username = "root"
	client.Port = 22
	return client
}
