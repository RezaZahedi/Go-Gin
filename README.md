# Go:  Gin/Micro


#### simple go app using gin, wire, and micro.

This application stores book names, descriptions, and IDs; it also has simple authentication mechanism using cookies,
the authentication is needed to create new books.  
The application computes Fibonacci numbers using recursive function memoization with concurrency safety, to speed up
answering to parallel requests and cache repeated function calls.  
The Fibonacci service is designed with the miro-service architecture in mind, somehow. As Gin REST API talks with it using 
RPG and Protocol Buffers.    
This module contains some redundant abstractions and redirections, like in user and product packages; they are mostly for the sake of illustration 
of Wire.  

The tests only for database package are complete.

<br>

##### how to run code

Visit `localhost:8080`. Register, then create some books and see them in the home page. Go to the Fibonacci tab to get
a Fibonacci number computed. It works fine with numbers less than 100,000.

The server supports HTML, JSON, and XML formats as answers.  
For example you can <br>
 `curl -H "Accept: application/json" 'localhost:8080/fibo/' -v -F 'number=8'`
 <br> or simply <br>
 `curl -H "Accept: application/json" 'localhost:8080/fibo/' -F 'number=8'`
<br> to get a Fibonacci number as an answer.  

<br>

##### Refrences

- Gin framework patterns are mostly from: [[1]](https://github.com/demo-apps/go-gin-app)
and [[2]](https://github.com/hellokoding/hellokoding-courses/tree/master/golang-examples/rest-gin-gorm/product)
- Templates are from: [[1]](https://github.com/demo-apps/go-gin-app)

- Wire concept is from: [[2]](https://github.com/hellokoding/hellokoding-courses/tree/master/golang-examples/rest-gin-gorm/product)
and [[3]](https://blog.golang.org/wire)

- Micro concepts from: [[4]](https://itnext.io/micro-in-action-part-3-calling-a-service-55d865928f11)

- Function memoization concepts are from: [[5]](http://www.gopl.io/)
or [[6]](https://github.com/adonovan/gopl.io/tree/master/ch9/memo5)  

- And other concepts from some other books and articles that I absolutely can't say which one is from which  
<br>

[1] : https://github.com/demo-apps/go-gin-app  
[2] : https://github.com/hellokoding/hellokoding-courses/tree/master/golang-examples/rest-gin-gorm/product  
[3] : https://blog.golang.org/wire  
[4] : https://itnext.io/micro-in-action-part-3-calling-a-service-55d865928f11  
[5] : http://www.gopl.io/  
[6] : https://github.com/adonovan/gopl.io/tree/master/ch9/memo5