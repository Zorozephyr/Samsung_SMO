Beginner Level Commands:
    go run (Compiles and executes one or two files)
    go build (Compiles and builds an executable)
    go fmt (Format all the code in the repository)
    go install (Compiles and installs a package)
    go get (Downloads the raw source code of someone else's package)
    go test (Runs any tests associated with the current project)

Package == Project == Workspace
Executable -> Generates a file that we can run
Reusable -> Code used as 'helpers'

package main -> Creates executable file(Must have a function called main)

Static Typing :var variableName variableType
Static Typed Inference: variableName := value (Infers type from variable)

To reassigning an existing variable, dont use colon, only use equal..:= is only for initial assignment

global variables can be declared but not assigned values outside of functions

In main.go:

package main
 
func main() {
    printState()
}


In a separate file called state.go:

package main
 
import "fmt"
 
func printState() {
    fmt.Println("California")
}
Your answer is correct because in Go, files that share the same package can directly access each other's functions without needing to import them, which allows the main.go file to call printState() from state.go seamlessly. This demonstrates a fundamental aspect of package organization in Go.



Arrays In Go:
1.Array(Fixed length list of things)
2.Slice(Array that can grow and shrink)
Every element in the slice and array should be of same type

cards := []string{<Place the elements in the slice array>}

To add new element:
cards = append(cards, newElement) //Append doesnt edit the original slice, it creates new slice while appending

for i,card := range cards{

}

To create a new type:
    type deck []string

    func (d deck) print() {
        for i, card := range d {
            fmt.Println(i, card)
        }
    }

    These functions are called recievers...Now every deck type variable can access this function
    By convention we call it with one letter like d in this case...But it is not compulsory

    slice[startIndexIncluding:upToNotIncluding]

From a function u can return multiple return values

func funcName(arguments) (returnTypes){}

ioutil package
func WriteFile(filename string, data []byte, perm os.FileMode) error

[]byte -> slice of bytes
byte slice is a way to represent a string

[]byte(StringVariable) -> Type conversion

go mod init cards

To create a test file, we need to create file ending with _test.go

Struct Declaration:
type structname struct{
    property1 type
    property2 type
}

Struct Initialization:
eg: alex:= person{"Alex", "Anderson"}
    alex:= person{firstName: "alex", lastName: "Anderson"}
eg2: var alex person
     If we do not add actual values for the properties, then it gets the zero values
     string => ""
     int => 0
     float => 0
     bool => false
eg3: var alex person
     alex.firstName = "Alex"
     alex.lastName = "Anderson"
eg4:
    type contactInfo struct{
        email string
        zipcode int
    }

    type person struct{
        firstName string
        lastName string
        contact contactInfo
    }

    jim := person{
        contact: contactInfo{
            email: "dsads@gmail.com"
            zipCode: 94000,
        },
    }
fmt.Printf("%+v",alex), will printout field names and its values

func (p person) print(){ -> Reciever for struct
    fmt.Printf("%+v",p)
}

func (p person) updateName(newFirstName string){
    p.firstName = newFirstName
} This will not update the original person object

Go is a pass by value language, the copied value is sent to the function 

jimPointer := &jim
jimPointer.updateName("jimmy)

func(pointerToPerson *person) updateName(newFirstName string){
    *pointerToPerson.firstName = newFirstName
}

& gives memory address
*pointer -> give value in that address

jimPointer := &jim
jimPointer.updateName("jimmy) -> Easier way to right the code

jim.updateNmae("jimmy") will work automatically even if the argument of the reciever is pointerToPerson

Maps:
key value pair
All keys should be of same type and all values have to be of same type

colors:= map[string]string{
    "red":"#ff0000",
    "green":"ff0011",   
}

var colors map[string]string //Initialization
colors:=make(map[string]string) //Initialization

colors["white"]="#ffffff"
delete(colors, keyName)

Iterating over maps:
func printMap (c map[string]string){
    for key,value := range c{

    }
}

Interfaces:
Every value has a type
Every function has to specify the type of its arguments

Interfaces are implicit
Interfaces are not generic types
Interfaces are a contract to help us manage types


Response Struct
Status(string)
StatusCode(int)
Body(io.ReadCloser)

ReadCloser Interface
Reader
Closer

Reader Interface
Read([]byte(int,error))

Closer Interface
Close()(error)

bs:=make([]byte,99999)
Read function doesnt automatically resize the slice while reading into it

io.Copy(os.Stdout, resp.Body)
Writer interface channels data outside our program

func Copy(dst Writer, src Reader) (written int64, err error)

os.Stdout is a File type, File has a function called 'Write', therefore it implements Write interface

Go Routines:
Our running program is a single go routine

We can use go keyword
eg:
for _,link:=range links{
    go checkLink(link)
}
New go routine is created...And if there is any sort of blocking call then control is given back to another go routine

Go Scheduler -> Uses only one cpu core by default
Scheduler detects blocking call and pauses go routine and runs another go routine

In case of multiple cpu core, it assigns accordingly and parallely

If Main routine ends, all child routines also stops, resulting in the program not running properly

Channels:
To communicate between different go routines
Channel also has a type, and messages send through channel can be of that specific type it was initialized with

channel <- 5 (send the value '5' into the value)
myNumber <- channel (Wait for a value to be sent into channel, when we get one assign to myNumber)
fmt.Println(<-channel>)

Recieving messages from a channel is a blocking  call

for l:= range channel{
    go checkLink(l,c)
}

function literal
go func(link string){
    time.Sleep(5*time.second)
    checkLink(link,c)
}(l)








