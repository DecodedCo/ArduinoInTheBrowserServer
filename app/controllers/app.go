package controllers

import (
	"bufio"
	"log"
	"io"
	"io/ioutil"
	"bytes"
	"github.com/revel/revel"
	"os"
	"os/exec"
	"strings"
)

type App struct {
	*revel.Controller
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//an object to hold responses
type Message struct {
	Message  string `json:"message"`
	Result   string `json:"result"`
	Identity string `json:"identity"`
}
//copy a file
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}
//copy an entire directory
func CopyDir(source string, dest string) (err error) {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		sourcefilepointer := source + "/" + obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()
		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				log.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

func getLibraries() string {

	dht := "#include <DHT.h>\n"
	servo := "#include <Servo.h>\n"
	softSerial := "#include <SoftwareSerial.h>\n"

	var buffer bytes.Buffer
	buffer.WriteString(dht)
	buffer.WriteString(servo)
	buffer.WriteString(softSerial)

	return buffer.String()
}

func (c App) Program() revel.Result {

	/*
		need to generate a unique idenitifier for the user, and return that to the user in the json
		so that when they request the hex we know who they are

	*/
	id, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Println("error!!")
	}
	identity := strings.TrimSpace(string(id))
	identity = strings.Trim(identity, " ")
	//out is this users id and must be returned to them!

	var qVal string
	c.Params.Bind(&qVal, "q")
	log.Println("q = ", qVal)

	var postedString string
	c.Params.Bind(&postedString, "program")
	log.Println("program = ", postedString)

	libraries := getLibraries()

	s := []string{"/srv/codefiles/", identity, ".ino"}
	tempCodeFile := strings.Join(s, "")

	f, err := os.Create(tempCodeFile)
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	finalString := []string{libraries, "\n\n", postedString}

	log.Println("final Program:\n\n", strings.Join(finalString, ""))

	n, err := w.WriteString(strings.Join(finalString, ""))

	log.Printf("wrote %d bytes\n", n)
	w.Flush()

	result := buildProject(identity)
	var m Message
	log.Println("result: ", result)
	if result == "done!" {
		m = Message{"compilation successful", string(result), identity}
	} else {
		m = Message{"compilation failed", string(result), identity}
	}
	log.Println("m: ", m)
	return c.RenderJson(m)
}

//run system commands on go routines to stop blocking
func buildProject(identity string) string {
	out, err := exec.Command("/bin/bash", os.Getenv("HOME")+"/command.sh", identity).Output()
	if err != nil {
		log.Println("Error: ", err)
		return string(out)
	}
	if strings.Contains(string(out), "Converting to firmware.hex") { //compilation was complete
		return "done!"
	} else {
		return string(out)
	}
	log.Printf("%s", out)
	return string(out)
}

func (c App) Hex() revel.Result {
	var identity string
	c.Params.Bind(&identity, "identity")
	log.Println("identity = ", identity)
	s := []string{"/srv/codefiles/", identity, ".hex"}
	dat, err := ioutil.ReadFile(strings.Join(s, ""))
	check(err)
	os.RemoveAll(strings.Join(s, ""))
	log.Println(string(dat))
	return c.RenderText(string(dat))
}

func (c App) Index() revel.Result {
	m := Message{"nothing at this location", "error 404", "NIL"}
	return c.RenderJson(m)
}
