# calc-go-lms 🦺🎇
Yandex lyceum Go final sprit 0 project. no-gui calculator app 

[Englishh readme ver.](https://github.com/sklerakuku/calc-go-lms/blob/main/README.en.md)


- [x] b（￣▽￣）d　
- [ ] 🐱‍🏍

## -------------------------------------------------------------------/______/..
# 🔢 API Простого калькулятора

Веб-сервис, который вычисляет арифметические выражения, пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.

Калькулятор поддерживает базовые арифметические операции (+, -, *, /) над положительными рациональными числами (дробная часть записывается через точку `1.22+3`), а также использования скобок для обозначения порядка выполнения операций. 

Калькулятор не поддерживает отрицательные числа ! 😄

Использование унарного минуса(или плюса) также приведёт к некорректной работе программы

<br><br>

## 💻 Запуск сервиса

Для запуска сервера на своём компьютере скачайте и запустите файл [server.exe](https://github.com/sklerakuku/calc-go-lms/releases/download/exe/server.exe)

### Альтернативный способ

Скачать всю папку calc-go-lms и запустить сервер `go run ./...`

```
git clone https://github.com/sklerakuku/calc-go-lms.git
```
```
go run ./...
```
<b>Для работы программы необходимо установить Golang https://go.dev/doc/install</b>

<br><br>

## 🧮 Немного о работе сервиса

У сервиса 1 endpoint с url-ом /api/v1/calculate

```
/api/v1/calculate
```

Для получения решения нужно оправить на этот url POST-запрос с телом:

```
{
    "expression": "выражение, которое нужно решить"
}
```

<br><br>

## 🖼 Примеры использования

—— [Корректный запрос - 200 (OK)](https://github.com/sklerakuku/calc-go-lms/tree/main?tab=readme-ov-file#%D0%BD%D0%B0%D0%B9%D0%B4%D0%B5%D0%BC-%D1%87%D0%B5%D0%BC%D1%83-%D1%80%D0%B0%D0%B2%D0%BD%D0%BE-%D0%B2%D1%8B%D1%80%D0%B0%D0%B6%D0%B5%D0%BD%D0%B8%D0%B5-222-%D0%B4%D0%BB%D1%8F-%D1%8D%D1%82%D0%BE%D0%B3%D0%BE-%D0%BE%D1%82%D0%BF%D1%80%D0%B0%D0%B2%D0%B8%D0%BC-post-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81-%D1%81%D0%B5%D1%80%D0%B2%D0%B8%D1%81%D1%83)

—— [Некорректный запрос - 422 (Unprocessable Entity)](https://github.com/sklerakuku/calc-go-lms/tree/main?tab=readme-ov-file#%D0%B5%D1%81%D0%BB%D0%B8-%D0%BC%D1%8B-%D0%B4%D0%BE%D0%BF%D1%83%D1%81%D1%82%D0%B8%D0%BC-%D0%B2-%D0%B2%D1%8B%D1%80%D0%B0%D0%B6%D0%B5%D0%BD%D0%B8%D0%B8-%D0%BE%D1%88%D0%B8%D0%B1%D0%BA%D1%83-%D0%BD%D0%B0%D0%BF%D0%B8%D1%88%D0%B5%D0%BC-%D0%B2%D1%8B%D1%81%D0%BA%D0%B0%D0%B7%D1%8B%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5-%D1%81-%D0%B1%D1%83%D0%BA%D0%B2%D0%B0%D0%BC%D0%B8-a2b)

—— [Некорректный запрос - 500 (Internal Server Error)](https://github.com/sklerakuku/calc-go-lms/tree/main?tab=readme-ov-file#%D0%B2-%D1%81%D0%BB%D1%83%D1%87%D0%B0%D0%B8-%D0%B8%D0%BD%D0%BE%D0%B3%D0%BE-%D0%BD%D0%B5%D0%BA%D0%BE%D1%80%D1%80%D0%B5%D0%BA%D1%82%D0%BD%D0%BE%D0%B3%D0%BE-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0-%D0%BF%D0%BE%D0%BB%D1%83%D1%87%D0%B8-%D0%BE%D1%88%D0%B8%D0%B1%D0%BA%D1%83--%D1%87%D1%82%D0%BE-%D1%82%D0%BE-%D0%BF%D0%BE%D1%88%D0%BB%D0%BE-%D0%BD%D0%B5-%D1%82%D0%B0%D0%BA)

<br><br>

#### Найдем чему равно выражение `2+2*2`, для этого отправим POST запрос сервису 

```cURL
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
	"expression": "2+2*2"
}'
```

Получим ожидаемый ответ `6` и кодом 200 (OK)

```json
{
    "result": "6.000000"
}
```

<br><br>


#### Если мы допустим в выражении ошибку, напишем высказывание с буквами `a+2*b` 

```cURL
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
	"expression": "2+2*2"
}'
```

Получим ошибку и код 422 (Unprocessable Entity)

```json
{
    "error": "Expression is not valid"
}
```

<br><br>


#### В случаи иного некорректного запроса, получи ошибку  «Что-то пошло не так» 

```cURL
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    
'
```

```json
{
    "error": "Internal server error"
}
```

<br><br><br>

<table>
<tr>
<td> Status </td> <td> Request </td> <td> Response </td>
</tr>
<tr>
<td> 200 (OK) </td>
<td>


```c
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "a + 2*b"
}'
```


</td>
<td>


```json
{
    "result": "6.000000"
}
```


</td>

<tr>
<td> 422 </td>
<td>


```c
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "a + 2*b"
}'
```


</td>
<td>


```json
{
    "error": "Expression is not valid"
}
```


</td>
<tr>
<td> 500 </td>
<td>


```c
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    
'
```


</td>
<td>


```json
{
    "error": "Internal server error"
}
```


</td>
</table>

<br>

![image](https://github.com/user-attachments/assets/f0e9462a-ae13-454c-a450-9e3d1fc85a06)



## Описание структуры проекта (под котиком т.к. котик лучше, чем ЭТО)

```
calc-go-lms 
├─ cmd
│  ├─ main.go                           - вход в программу
├─ internal
│  ├─ application
│  │  ├─ application_test.go            - тесты application
│  │  ├─ application.go                 - http сервер ... 
├─ pkg
│  ├─ calculation       
│  │  ├─ calculation_test.go            - тесты calculation
│  │  ├─ calculation.go                 - логика вычислений 
├─ go.mod                               - файл с зависимостями (наркоман)
├─ .gitignore
├─ README.en.md
├─ README.md
├─ server.exe                           - исполнительный файл win
```
