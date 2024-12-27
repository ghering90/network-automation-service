package main
 
import (
	"bytes"
	"encoding/json"
	"os"
	"fmt"
	"log"
	"golang.org/x/crypto/ssh"
)

// Define a struct to hold the credentials
type SSHCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func loadCredentials(filename string) (*SSHCredentials, error) {
	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open secrets file: %v", err)
	}
	defer file.Close()

	// Decode the JSON file into the struct
	var creds SSHCredentials
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&creds); err != nil {
		return nil, fmt.Errorf("failed to decode secrets file: %v", err)
	}
	return &creds, nil
}


func main()  {
	// Load credentials from the secrets file
	creds, err := loadCredentials("secrets.json")
	if err != nil {
		log.Fatalf("Error loading credentials: %v", err)
	}


	// SSH client configuration
	config := &ssh.ClientConfig{
		User: creds.Username, // Replace with your username
		Auth: []ssh.AuthMethod{
			ssh.Password(creds.Password), // Replace with your password
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Use a secure callback in production
	}

	// Connect to the remote host
	address := fmt.Sprintf("%s:%s", creds.Host, creds.Port)
	client, err := ssh.Dial("tcp", address, config) 
	if err != nil {
		log.Fatalf("Failed to connect to remote host: %v", err)
	}
	defer client.Close()

	// Start a new session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create SSH session: %v", err)
	}
	defer session.Close()

	// Buffer to capture command output
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	// Run the `date` command
	if err := session.Run("date"); err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}

	// Print the command output
	fmt.Println("Command Output:")
	fmt.Println(stdoutBuf.String())
	
}



// func count(x,y int) int {
// 	return x + y
// }
// func loop_list(items []byte) {
// 	for i , item := range items {
// 		fmt.Printf("item %d: %d\n", i, item)
// 	}
// }

// func sum_list(items []byte) int {
// 	sum_dig := 0
// 	for _ , item := range items {
// 		sum_dig += int(item)
// 	}
// 	return sum_dig
// }


// func main() {
// 	fmt.Println("hi")
// 	fmt.Println(count(5,6))
// 	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	if err == nil {
// 		//fmt.Printf("Response: %+v\n", resp)
// 		body, _ := io.ReadAll(resp.Body)
// 		//fmt.Println(body)
// 		loop_list(body)
// 		fmt.Println(sum_list(body))
// 	}

// }