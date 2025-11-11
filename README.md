# Project: CRM_Backend

### <ins>Description:</ins>

This project involves building the backend (i.e. server-side portion) of a CRM application. The backend will allow the user make HTTP requests to the Postman server to perform CRUD operations (Create, Read, Update, and Delete). The mock customer data will allow user to perform CRUD operations in Postman application.

**Visual Studio Code** is the application software where the project is made, edited, and tested. The **Postman** tool is used as the frontend of the CRM application when running and testing the project. **GitHub Desktop** is the application software that gives access to the project from _GitHub_ and allows it to be edited in _Visual Studio Code_.

One challenge that has been faced is defining the following import statement in the project:
```
"github.com/gorilla/mux"
```
This challenge was overcome after doing some research and doing some trial and error with commands that were given by other people on _Stack Overflow_ until it was found that the following text below would define the import statement.
```
go env -w GO111MODULE=auto
go mod init
go mod tidy
```
Another challenge that has been faced is encoding the map holding the customer data. The map that was supposed to be encoded was ```map[uint32]struct``` type, not ```map[string]string``` type, making it a challenge to encode the map. After doing some research, editing the project code, and doing some trial and error with the code, it was discovered that maps with ```map[uint32]struct``` type need to be customly encoded and cannot be encoded using the built-in encoding method. This challenge was overcome by creating a new map with ```map[string]string``` type that held the same customer data as the map with ```map[uint32]struct``` type and encoding the map with ```map[string]string``` type using the built-in encoding method to keep the code simple and encode the map. 

Another challenge that has been faced is ensuring that the correct chosen customer data was updated and/or removed. Sometimes after the customer was chosen to be updated or deleted, the wrong customer would be updated or deleted after it was found in the customer map. This challenge was overcome by adding code that displays the customer map in order by their IDs in the ```getAllCustomer()``` function.

### <ins>Creation Date:</ins>

> 10/26/2025

### <ins>Languages:</ins>

**GO (Golang)**

<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a>

### <ins>Tools:</ins>

**Postman**

<img src="https://www.vectorlogo.zone/logos/getpostman/getpostman-icon.svg" alt="postman" width="40" height="40"/>


## How to Install and Run the Project

1. Clone project
2. Open Visual Studio Code _(optional)_
3. Open Terminal in Visual Studio Code or local Command Prompt
4. Change directory to where "main.go" file is being kept
5. _If this is the **first time** running the project_, use the following text below to allow project to run
```
go env -w GO111MODULE=auto
go mod init
go mod tidy
```
6. Use the following text below to run the project
```
go run main.go
```
7. Hit **Ctrl+C** to exit project


## How to Use the Project

1. Open Postman
2. _If collection does not exist_, create new collection to hold request
3. Create new request
4. Enter URL that matches chosen request
5. Choose Method that matches chosen request (GET, POST, PUT, DELETE)
6. Hit the "Send" button to send the request


## Credits
###### References used while making project

Antolín-Camarena, Omar, and Jeremy Wall. “How to Use Arbitrary Length Sequences of Values as Map Keys in Go?” _Stack Overflow_, 12 Apr. 2012, stackoverflow.com/questions/10116476/how-to-use-arbitrary-length-sequences-of-values-as-map-keys-in-go. Accessed 27 Oct. 2025.

asker152, et al. “Error Message ‘Go: Go.Mod File Not Found in Current Directory or Any Parent Directory; See “Go Help Modules.”’” Edited by DailyLearner and Peter Mortensen, _Stack Overflow_, 31 Mar. 2021, stackoverflow.com/questions/66894200/error-message-go-go-mod-file-not-found-in-current-directory-or-any-parent-dire. Accessed 28 Oct. 2025.

B., Clive. “Concatenating Strings in Go.” _Sentry_, 15 Oct. 2024, sentry.io/answers/concatenating-strings-in-go/. Accessed 6 Nov. 2025.

“Basic Writing and Formatting Syntax.” _GitHub Docs_, docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax. Accessed 9 Nov. 2025.

BenjB, et al. “How to Use a `switch` Statement in Go.” _Stack Overflow_, 21 Dec. 2021, stackoverflow.com/questions/70438741/how-to-use-a-switch-statement-in-go. Accessed 6 Nov. 2025.

Biswas, Pradip, and fuglede. “Issues in Go with JSON Unmarshal of Map[String]Int64.” Edited by Fuglede, _Stack Overflow_, 9 Oct. 2019, stackoverflow.com/questions/58307298/issues-in-go-with-json-unmarshal-of-mapstringint64#:~:text=As%20you’ve%20noted%20yourself,18.3k3%2062%20107. Accessed 4 Nov. 2025.

Chris, Kolade. “CRUD Operations – What Is CRUD?” _freeCodeCamp.Org_, freeCodeCamp.org, 15 June 2022, www.freecodecamp.org/news/crud-operations-explained/. Accessed 9 Nov. 2025.

“Constants.” _Http Package - Net/Http - Go Packages_, pkg.go.dev/net/http#pkg-constants. Accessed 28 Oct. 2025.

Dave, and Juan Carlos Garcia. “Sorting Unsigned Ints in Go.” _Stack Overflow_, 12 Jan. 2017, stackoverflow.com/questions/41624074/sorting-unsigned-ints-in-go. Accessed 7 Nov. 2025.

“Delete Elements in a Slice in Golang.” _GeeksforGeeks_, GeeksforGeeks, 23 July 2025, www.geeksforgeeks.org/go-language/delete-elements-in-a-slice-in-golang/. Accessed 8 Nov. 2025.

Demir, David. “What Are HTTP Methods (GET, POST, PUT, DELETE).” _Apidog_, 31 July 2025, apidog.com/blog/http-methods/. Accessed 8 Nov. 2025.

“Different Ways to Concatenate Two Strings in Golang.” _GeeksforGeeks_, GeeksforGeeks, 28 Oct. 2024, www.geeksforgeeks.org/go-language/different-ways-to-concatenate-two-strings-in-golang/#using-the-operator. Accessed 6 Nov. 2025.

“Different Ways to Find the Type of Variable in Golang.” _GeeksforGeeks_, GeeksforGeeks, 12 July 2025, www.geeksforgeeks.org/go-language/different-ways-to-find-the-type-of-variable-in-golang/. Accessed 29 Oct. 2025.

DJElbow, and ANisus. “Sorting a Map of Structs - GOLANG.” _Stack Overflow_, 13 Nov. 2013, stackoverflow.com/questions/19946992/sorting-a-map-of-structs-golang. Accessed 7 Nov. 2025.

“func Slice.” _Sort Package - Sort - Go Packages_, pkg.go.dev/sort@master#Slice. Accessed 7 Nov. 2025.

“GitHub Profile README Generator.” _GitHub Profile Readme Generator | GitHub Profile Readme Generator_, rahuldkjain.github.io/gh-profile-readme-generator/. Accessed 9 Nov. 2025.

“Go by Example: For.” _Go by Example: For_, gobyexample.com/for. Accessed 28 Oct. 2025.

“Go by Example: Switch.” _Go by Example: Switch_, gobyexample.com/switch. Accessed 6 Nov. 2025.

Han, and Arin Yazilim. “Go Error: Go : Go.Mod File Not Found in Current Directory or Any Parent Directory; (Working on GOPATH/Src).” _Stack Overflow_, 1 Sept. 1965, stackoverflow.com/questions/67929883/go-error-go-go-mod-file-not-found-in-current-directory-or-any-parent-director/68508392#68508392. Accessed 28 Oct. 2025.

hey, and julienc. “How to Convert Uint32 to String?” Edited by Nik and Fuz, _Stack Overflow_, 22 July 2014, stackoverflow.com/questions/24886015/how-to-convert-uint32-to-string. Accessed 5 Nov. 2025.

hhh, et al. “Simple Way of Getting Key Depending on Value from Hashmap in Golang.” Edited by User142162, _Stack Overflow_, 14 Nov. 2015, stackoverflow.com/questions/33701828/simple-way-of-getting-key-depending-on-value-from-hashmap-in-golang. Accessed 30 Oct. 2025.

“How to Sort Golang Map By Keys or Values?” _GeeksforGeeks_, GeeksforGeeks, 4 Apr. 2024, www.geeksforgeeks.org/go-language/how-to-sort-golang-map-by-keys-or-values/. Accessed 7 Nov. 2025.

“How to Take Input from the User in Golang?” _GeeksforGeeks_, GeeksforGeeks, 10 May 2020, www.geeksforgeeks.org/go-language/how-to-take-input-from-the-user-in-golang/. Accessed 28 Oct. 2025.

Ihab, Naguib, et al. “Return Nil for a Struct in Go.” Edited by Jonathan Hall, _Stack Overflow_, 5 June 2018, stackoverflow.com/questions/50697914/return-nil-for-a-struct-in-go. Accessed 30 Oct. 2025.

Joiner, Matt, and maerics. “Convert String to Integer Type in Go?” Edited by Icza and Pallav Agarwal, _Stack Overflow_, 9 July 2020, stackoverflow.com/questions/4278430/convert-string-to-integer-type-in-go. Accessed 3 Nov. 2025.

Kelche, Kevin. “How to Accept User Input with Spaces in Golang.” _Kelche_, 29 July 2023, www.kelche.co/blog/go/user-input-with-spaces/. Accessed 29 Oct. 2025.

Kfir, and mbuechmann. “Use different structs as a value in map golang.” _Stack Overflow_, 8 Nov. 2018, stackoverflow.com/questions/53204138/use-different-structs-as-a-value-in-map-golang. Accessed 27 Oct. 2025.

Koharim67, and Ezequiel Muns. “Find an Item in Array of Maps and Delete It.” _Stack Overflow_, 26 Oct. 2021, stackoverflow.com/questions/69722669/find-an-item-in-array-of-maps-and-delete-it. Accessed 8 Nov. 2025.

“MIT License.” _Choose a License_, 26 Oct. 2025, choosealicense.com/licenses/mit/. Accessed 9 Nov. 2025.

“MLA Works Cited: Electronic Sources (Web Publications).” _MLA Works Cited: Electronic Sources - Purdue OWL® - Purdue University_, owl.purdue.edu/owl/research_and_citation/mla_style/mla_formatting_and_style_guide/mla_works_cited_electronic_sources.html. Accessed 26 Oct 2025.

Naresh, and peterSO. “Convert String to Uint in Go Lang.” Edited by Danack, _Stack Overflow_, 2 Feb. 2016, stackoverflow.com/questions/35154875/convert-string-to-uint-in-go-lang. Accessed 3 Nov. 2025.

Nyakundi, Hillary. “How to Write a Good README File for Your GitHub Project.” _freeCodeCamp.Org_, freeCodeCamp.org, 8 Dec. 2021, www.freecodecamp.org/news/how-to-write-a-good-readme-file/. Accessed 9 Nov. 2025.

“Online Compiler and Debugger for C/C++.” _GDB Online Debugger_, www.onlinegdb.com/. Accessed 28 Oct. 2025.

Philip, and Jeff Emanuel. “Using Count to Find the Last Element in Struct / Map.” _Go Forum_, 4 Feb. 2022, forum.golangbridge.org/t/using-count-to-find-the-last-element-in-struct-map/26367. Accessed 1 Nov. 2025.

r/golang. “Convert Uint to String.” _Reddit_, 2024, www.reddit.com/r/golang/comments/1gjlerl/convert_uint_to_string/. Accessed 3 Nov. 2025.

r/golang, and edgmnt_net. “CRUD, Return Empty Struct or Nil?” _Reddit_, 2022, www.reddit.com/r/golang/comments/vd1upf/crud_return_empty_struct_or_nil/. Accessed 30 Oct. 2025.

r/golang, and phlatphrog. “How to Pass a Generic Map Parameter to Function.” _Reddit_, 2017, www.reddit.com/r/golang/comments/7p879p/how_to_pass_a_generic_map_parameter_to_function/. Accessed 27 Oct. 2025.

“README 101.” _Make a README_, GitHub, www.makeareadme.com/. Accessed 9 Nov. 2025.

Rees, Lucas. “Working with Null Values in Go: A Complete Guide.” _Go/Golang Programming and Scripting_, 2 Feb. 2021, golang.howtos.io/working-with-null-values-in-go-a-complete-guide/. Accessed 30 Oct. 2025.

RoboTamer, and Ekkehard.Horner. “Golang Map Prints out of Order.” _Stack Overflow_, 24 Aug. 2012, stackoverflow.com/questions/12108215/golang-map-prints-out-of-order. Accessed 6 Nov. 2025.

Sanjeev. “Custom Golang HTTP Router.” _Medium_, Medium, 13 June 2021, sanjeevsiva.medium.com/custom-golang-http-router-970a309531d7. Accessed 28 Oct. 2025.

Sarraf, Anchal, and T. Claverie. “How to Delete an Element from a Slice in Golang.” Edited by Carson and Community, _Stack Overflow_, 19 May 2016, stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang. Accessed 8 Nov. 2025.

serpent7655, and uhhd. “R/Golang on Reddit: Anyone Knows How Can Go Sort Struct by Multiple Fields?” _Reddit_, 2024, www.reddit.com/r/golang/comments/1fnyydy/anyone_knows_how_can_go_sort_struct_by_multiple/. Accessed 7 Nov. 2025.

“Strconv.” _Strconv Package - Strconv - Go Packages_, pkg.go.dev/strconv. Accessed 3 Nov. 2025.

“Switch.” _A Tour of Go_, go.dev/tour/flowcontrol/9. Accessed 6 Nov. 2025.

“Switch Statement in Go.” _GeeksforGeeks_, GeeksforGeeks, 6 Nov. 2024, www.geeksforgeeks.org/go-language/switch-statement-in-go/. Accessed 6 Nov. 2025.

Wada, Kengo. “Go Routing 101: Handling and Grouping Routes with Net/Http.” _DEV Community_, 2 Nov. 2024, dev.to/kengowada/go-routing-101-handling-and-grouping-routes-with-nethttp-4k0e. Accessed 28 Oct. 2025.

Warmth, Vinyl, and BadZen. “Why Is My Go Lang String Comparison Not Working as Expected?” _Stack Overflow_, 5 Nov. 2016, stackoverflow.com/questions/40442971/why-is-my-go-lang-string-comparison-not-working-as-expected. Accessed 29 Oct. 2025.

“Yuku,” Randy Sugianto, et al. “How to Efficiently Concatenate Strings in Go.” _Stack Overflow_, 15 Dec. 2015, stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go. Accessed 6 Nov. 2025.


## License

[MIT](https://choosealicense.com/licenses/mit/)
