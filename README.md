# omamatest
# soal 0 - Polindrom
input two numbers separated by a space character (a n) to produce a number series between a and n.
then from the series of numbers the polindrome numbers will be selected,
the output will produce a polindron number length

exp:.
    Input:
    22 25

    Output:
    1

# soal 1 - sorting kode buku berdasarkan kategori buku
random input book code series separated by spaces, book code consists of category code, book title, and book height
code book: "9Z99"
     9 Category codes
     Z Book Title
     99 book height

then the system will sort the data by category order,
then the book codes are sorted by book height, highest to lowest

exp:.
    input: 3A13 5X19 9Y20 2c18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14
    output: 0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2c18 5X19 5G14 3N21 3N20 3N14 3A13

# soal 2 - bilangan yang hilang dalam deret angka
input a series of numbers of type array string, where one of the numbers is omitted
the system will calculate the number of number series with the formula Sn = 1/2 * n * (a + un) then compare it with the number of existing numbers
(u1 + u2 + u3 + ...)
the output is a missing number

exp:. 
    deretAngka := []int{23, 24, 25, 26, 27, 28, 30}

    result:
        Input:  23242526272830
        Nilai yang hilang:  29

# soal 3 - webservice untuk soal no 2
run the application with the command go run main.go then enter,
open postman or browser and type http: // localhost: 8081
end point:
    get http://localhost:8081/index

# soal 4 - docker service
tools used:
     docker toolbox
     postman or browser
     windows powershell
     Visual studio code


create a file named Dockerfile,
change environment golang GOOS = windows to GOOS = linux via windows powershell with the command $ Env: GOOS = "linux", to check the changes use the go env command.
compile the main.go file with the go build command main.go via windows powershell or Viscode so that it produces a new file with the name "main" without extension
open the Docker Quickstart Terminal, and run the command to create docker images using the Dockerfile:
     docker build -t <name docker images> -> to create a docker image
     docker images -> to see the docker images that have been created
     docker run -it -p 8080: 8081 <name docker images> ->to run the webservice
     docker stop <name docker images> -> to stop the webservice

to access the endpoint on the browser or postman use the docker ip 192.168.100 and the port that has been declared is 8080

# test
command:
    go test

output:
if the value equals the true value(
    PASS
    ok      omamatest/<folder_name>     0.000s
)
else(
    --- FAIL: <Test_function_name> (0.00s)
    <file_name>.go:11: Expected 1.5, got  <output_value>
    FAIL
    exit status 1
    FAIL    omamatest/<folder_name>     0.036s
)