package main

import (
	"bytes"
	"container/list"
	"crypto/sha512"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	// stringsExercises()
	// ioExercises()
	// fileExercises()
	// containersExercises()
	// hashesExercises()
	serverExercises()
	// commandLineExercises()
}

// STRINGS
func stringsExercises() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Repeat("*", 40))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))

	arr := []byte("test")
	fmt.Println(arr)
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)
}

// FILES
func ioExercises() {
	var buf bytes.Buffer
	fmt.Println(buf, &buf)
	buf.Write([]byte("test"))
	fmt.Println(buf, &buf)
}

func fileExercises() {
	bs, err := ioutil.ReadFile("assets/ch8-test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	file, err := os.Create("assets/ch8-test-output.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString(str)

	dir, err := os.Open("./assets/")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

// CONTAINERS
type Person struct {
	Name string
	Age  int
}
type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func containersExercises() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}

// HASHES
func getHash(filename string) ([]byte, error) {
	// open the file
	f, err1 := os.Open(filename)
	if err1 != nil {
		return nil, err1
	}
	defer f.Close()
	// h := crc32.NewIEEE()
	h := sha512.New()
	_, err2 := io.Copy(h, f)
	if err2 != nil {
		return nil, err2
	}
	return h.Sum([]byte{}), nil
}

func hashesExercises() {
	h1, err := getHash("assets/ch8-test.txt")
	if err != nil {
		return
	}
	h2, err := getHash("assets/ch8-test-output.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, string(h1) == string(h2))
}

// SERVER NOTES
func tcpServer() {
	// listen on a port
	ln, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}
func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}
func tcpClient() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// send the message
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Hello, World</title>
			</head>
			<body>
			Hello, World!
			</body>
		</html>`,
	)
}

type RPCServer struct{}

func (this *RPCServer) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}
func rpcServer() {
	rpc.Register(new(RPCServer))
	ln, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
func rpcClient() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}

func tcpServerExercise() {
	// these examples seem to be broken
	// dial tcp 127.0.0.1:9999: connect: connection refused

	go tcpServer()
	go tcpClient()
	var input string
	fmt.Scanln(&input)
}

func httpServerExercise() {
	http.HandleFunc("/hello", hello)
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
	http.ListenAndServe(":9000", nil)
}

func rpcServerExercise() {
	// these examples seem to be broken
	// dial tcp 127.0.0.1:9999: connect: connection refused
	go rpcServer()
	go rpcClient()
	var input string
	fmt.Scanln(&input)
}

func serverExercises() {
	// tcpServerExercise()
	httpServerExercise()
	// rpcServerExercise()
}

// Parsing Command-Line arguments
func commandLineExercises() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}

// Creating packages
// see src/ch8
// todo: need to get this working

// Documentation

/* Exercises
1: re-use code between applications
+ break up larger programs into smaller, easier to understand and easier to maintain programs.
2: Capital casing are exported and accessible from other packages
3: package alias for a package, specified when importing; import m "package-name/sub-dir"; m.Average(3, 5)
4: see example-packages/math/math.go
5: Adding a comment above the functions
*/
