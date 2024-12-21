# calc-go-lms
Yandex lyceum Go final sprit 0 project. no-gui calculator app 


## ---------------------
# API Простого калькулятора

Веб-сервис, который вычисляет арифметические выражения, пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.

## Запуск сервиса

Для запуска сервера на своём компьютере скачайте файл [server.exe](https://github.com/sklerakuku/calc-go-lms/releases/download/exe/server.exe)


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

—— [Корректный запрос - 200 (OK)](https://github.com/sklerakuku/calc-go-lms/tree/main?tab=readme-ov-file#%D0%BD%D0%B0%D0%B9%D0%B4%D0%B5%D0%BC-%D1%87%D0%B5%D0%BC%D1%83-%D1%80%D0%B0%D0%B2%D0%BD%D0%BE-%D0%B2%D1%8B%D1%80%D0%B0%D0%B6%D0%B5%D0%BD%D0%B8%D0%B5-222-%D0%B4%D0%BB%D1%8F-%D1%8D%D1%82%D0%BE%D0%B3%D0%BE-%D0%BE%D1%82%D0%BF%D1%80%D0%B0%D0%B2%D0%B8%D0%BC-post-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81-%D1%81%D0%B5%D1%80%D0%B2%D0%B8%D1%81%D1%83)

—— [Некорректный запрос - 422 (Unprocessable Entity)](https://github.com/sklerakuku/calc-go-lms/tree/main?tab=readme-ov-file#%D0%B5%D1%81%D0%BB%D0%B8-%D0%BC%D1%8B-%D0%B4%D0%BE%D0%BF%D1%83%D1%81%D1%82%D0%B8%D0%BC-%D0%B2-%D0%B2%D1%8B%D1%80%D0%B0%D0%B6%D0%B5%D0%BD%D0%B8%D0%B8-%D0%BE%D1%88%D0%B8%D0%B1%D0%BA%D1%83-%D0%BD%D0%B0%D0%BF%D0%B8%D1%88%D0%B5%D0%BC-%D0%B2%D1%8B%D1%81%D0%BA%D0%B0%D0%B7%D1%8B%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5-%D1%81-%D0%B1%D1%83%D0%BA%D0%B2%D0%B0%D0%BC%D0%B8-a2b)

—— [Некорректный запрос - 500 (Internal Server Error)](https://github.com/sklerakuku/calc-go-lms/tree/main?tab=readme-ov-file#%D0%B2-%D1%81%D0%BB%D1%83%D1%87%D0%B0%D0%B8-%D0%B8%D0%BD%D0%BE%D0%B3%D0%BE-%D0%BD%D0%B5%D0%BA%D0%BE%D1%80%D1%80%D0%B5%D0%BA%D1%82%D0%BD%D0%BE%D0%B3%D0%BE-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0-%D0%BF%D0%BE%D0%BB%D1%83%D1%87%D0%B8-%D0%BE%D1%88%D0%B8%D0%B1%D0%BA%D1%83--%D1%87%D1%82%D0%BE-%D1%82%D0%BE-%D0%BF%D0%BE%D1%88%D0%BB%D0%BE-%D0%BD%D0%B5-%D1%82%D0%B0%D0%BA)


#### Найдем чему равно выражение `2+2*2`, для этого отправим POST запрос сервису 

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


#### Если мы допустим в выражении ошибку, напишем высказывание с буквами `a+2*b` 

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


#### В случаи иного некорректного запроса, получи ошибку  «Что-то пошло не так» 

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

