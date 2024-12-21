## -------------------------------------------------------------------/______/..
# 🔢 API for Simple Calculator

Web service that calculates arithmetic expressions. The user sends an arithmetic expression via HTTP and receives the result in response.

<br><br>

## 💻 Running the Service

To run the server on your computer, download and run the file [server.exe](https://github.com/sklerakuku/calc-go-lms/releases/download/exe/server.exe)

### Alternative Method

Download the entire `calc-go-lms` folder and run the server using `go run ./...`

<br><br>

## 🧮 About the Service

The service has 1 endpoint with the URL `/api/v1/calculate`

```
/api/v1/calculate
```

To get the solution, send a POST request to this URL with the body:

```json
{
"expression": "the expression to be solved"
}
```

<br><br>

## 🖼 Examples of Use

—— [Correct Request - 200 (OK)](https://github.com/sklerakuku/calc-go-lms/tree/main?tab=readme-ov-file#find-out-what-the-expression-is-equal-to-for-this-we-will-send-a-post-request-to-the-service)

—— [Incorrect Request - 422 (Unprocessable Entity)](https://github.com/sklerakuku/calc-go-lms/tree/main?tab=readme-ov-file#if-we-make-an-error-in-the-expression-write-a-statement-with-letters-a+2*b)

—— [Incorrect Request - 500 (Internal Server Error)](https://github.com/sklerakuku/calc-go-lms/tree/main?tab=readme-ov-file#in-case-of-another-incorrect-request-receive-an-error-«something-went-wrong»)

<br><br>

#### Find out what the expression `2+2*2` is equal to. To do this, send a POST request to the service.

```cURL
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2*2"
}'
```

Receive the expected response 6 with code 200 (OK)
```json
{
    "result": "6.000000"
}
```

If we make an error in the expression, write a statement with letters a+2*b.

```
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "a + 2*b"
}'
```

Receive an error with code 422 (Unprocessable Entity)
```json
{
    "error": "Expression is not valid"
}
```

In case of another incorrect request, receive an error «something went wrong».
```text
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    
}'
```
```json
{
    "error": "Internal server error"
}
```


<table> <tr> <td> Status </td> <td> Request </td> <td> Response </td> </tr> <tr> <td> 200 (OK) </td> <td>
  
```c
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2*2"
}'
```
  
</td> <td>
  
```json
{
    "result": "6.000000"
}
```
  
</td> <tr> <td> 422 </td> <td>
  
```c
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "a + 2*b"
}'
```
  
</td> <td>
  
```json
{
    "error": "Expression is not valid"
}
  ```

</td> <tr> <td> 500 </td> <td>
  
```c
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    
}'
```
</td> <td>
  
```json
{
    "error": "Internal server error"
}
  ```
</td> </table>
