package connection

import (
    "net"
    "log"
    "bufio"
    "net/http"
    "github.com/Vallerious/http_server/handlers"
)

func HandleConnection(c net.Conn) error {
    log.Println("received connection")
    reader := bufio.NewReader(c)
    req, err := http.ReadRequest(reader)
    
    if err != nil {
        return err
    }

    log.Printf("Incoming request on path: %s and method: %s\n", req.URL, req.Method)
    
    b, _ := handlers.FileHandler(req)

    log.Println(string(b))

    writer := bufio.NewWriter(c)
    writer.WriteString(string(b))
    writer.Flush()
    c.Close()

    return nil
}


