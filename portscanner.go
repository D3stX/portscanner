package main

import (
    "fmt"
    "net"
    "strings"
)

const asciiArt = `
$$\   $$\           $$\           $$$$$$\                      
$$$\  $$ |          \__|         $$  __$$\                     
$$$$\ $$ | $$$$$$\  $$\  $$$$$$\ $$ /  \__| $$$$$$\   $$$$$$$\ 
$$ $$\$$ |$$  __$$\ $$ |$$  __$$\\$$$$$$\  $$  __$$\ $$  _____|
$$ \$$$$ |$$ /  $$ |$$ |$$ |  \__|\____$$\ $$$$$$$$ |$$ /      
$$ |\$$$ |$$ |  $$ |$$ |$$ |     $$\   $$ |$$   ____|$$ |      
$$ | \$$ |\$$$$$$  |$$ |$$ |     \$$$$$$  |\$$$$$$$\ \$$$$$$$\ 
\__|  \__| \______/ \__|\__|      \______/  \_______| \_______|
                                                               
                                                               
`

func main() {
    fmt.Printf("\033[35m%s\033[0m\n", asciiArt)

    var targets string
    var ports int

    fmt.Print("[*] Enter Targets To Scan (split them by ,): ")
    fmt.Scanln(&targets)

    fmt.Print("[*] Enter How Many Ports You Want To Scan: ")
    fmt.Scanln(&ports)

    if strings.Contains(targets, ",") {
        fmt.Println("[*] Scanning Multiple Targets")
        ips := strings.Split(targets, ",")
        for _, ip := range ips {
            scan(strings.TrimSpace(ip), ports)
        }
    } else {
        scan(targets, ports)
    }
}

func scan(target string, ports int) {
    fmt.Printf("\n Starting Scan For %s\n", target)
    for port := 1; port <= ports; port++ {
        scanPort(target, port)
    }
}

func scanPort(ipaddress string, port int) {
    target := fmt.Sprintf("%s:%d", ipaddress, port)
    conn, err := net.Dial("tcp", target)
    if err != nil {
        return
    }
    conn.Close()
    fmt.Printf("[+] Port Opened %d\n", port)
}
