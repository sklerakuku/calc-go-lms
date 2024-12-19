# calc-go-lms
Yandex lyceum Go final sprit 0 project. no-gui calculator app 


## ---------------------
# API Простого калькулятора

###### Веб-сервис, который вычисляет арифметические выражения, пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.

## Запуск сервиса

Для запуска скопируйте команду ниже и вставьте в командную строку


## Немного о работе сервиса

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

## Примеры использования

- [[#^a0c94d|Корректный запрос - 200 (OK)]]
- [[#^bd360b|Некорректный запрос - 422 (Unprocessable Entity)]]
- [[#^59f5d0|Некорректный запрос - 500 (Internal Server Error)]]


Найдем чему равно выражение `2+2*2`, для этого отправим POST запрос сервису 

```cURL
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
	"expression": "2+2*2"
}'
```

Получим ожидаемый ответ `6` и кодом 200 (OK)

```JSON
{
    "result": "6.000000"
}
```

Если мы допустим в выражении ошибку, напишем высказывание с буквами `a+2*b` 

```cURL
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "a + 2*b"
}'
```

Получим ошибку и код 422 (Unprocessable Entity)

```JSON
{
    "error": "Expression is not valid"
}
```

В случаи иного некорректного запроса, получи ошибку  «Что-то пошло не так» 

```cURL
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    
'
```

```JSON
{
    "error": "Internal server error"
}
```

