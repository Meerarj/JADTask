 Creating a REST based application having GET method using net/http package

Step 1 : Using your text editor, create a file called request.go in the product_api directory. You’ll write your Go code in this file.
         Let’s make the current directory the root of a module by using command "go mod init".
         The go.mod file only appears in the root of the module. 
         Packages in subdirectories have import paths consisting of the module path plus the path to the subdirectory.

Step 2 : Into request.go, at the top of the file, paste the following package declaration.
            package main
         A standalone program (as opposed to a library) is always in package main.

Step 3 : Beneath the package declaration import the packages
           "fmt"
	   "io/ioutil"
	   "net/http" 
	   "reflect"

Step 4 : Assign the handler function to an endpoint path.
            func main(){
            }
         This sets up an association in which handles requests to the endpoint path.

Step 5 : Assign URL to read data from a server.
         
         URL to read this data:(This API is in particular used in the Open Food Facts.)
 
           https://world.openfoodfacts.org/api/v0/product/[barcode].json

         URl used in our code 

           url := "https://world.openfoodfacts.org/api/v0/product/737628064502.json"
         
         You can also get the result in XML by using .xml

           url = "https://world.openfoodfacts.org/api/v0/product/737628064502.xml" 

step 6 : The net/http package we imported has a Get function used for making GET requests.
         This package provides us with all the utilities we need to make HTTP requests.
         Update your HTTP request code so that instead of using http.get to make a request to the server, 
         you use http.NewRequest and http.DefaultClient’s Do method

Step 7 : Once the http.NewRequest is created and configured, used the Do method of http.DefaultClient to send the request to the server.
         The http.DefaultClient value is Go’s default HTTP client, the same you’ve been using with http.Get.
         The Do method of the HTTP client returns the same values you received from the http.Get function so that you can handle the response in the same way.

Step 8 : By simply checking for an error before doing anything with the response we can conclude what we can and cannot do next.

         If the error is not nil, we have an error.Print an error.
         If the error is nil, we don’t have an error. We can freely use the response as we originally intended.

Step 9 : Defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns.
         The response body should be closed after the response is fully read. This is done to prevent resource leak of connections.

Step 10 :To read the body of the response, we need to access its Body property first.
         We can access the Body property of a response using the ioutil.ReadAll() method.

Step 11 :Stored the data of body in a variable.

Step 12 :The fmt.Printf() function in Go language formats according to a format specifier and writes to standard output.

Step 13 :For more readability we used reflect package
         It allows us to examine types at runtime.

Step 14 :Now save the file and run the program using command "go run".



          
