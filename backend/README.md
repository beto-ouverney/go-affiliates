# Go Affiliates API#

Go Affiliates API is a API to manager content producers and affiliates sales.

## Conteúdos

- [General View](#general-view)
    - [The challenge](#the-chalenge)
- [The development](#the-development-process)
    - [Tools used](#tools-used)
    - [Lessons learned](#lessons-learned)
- [Usage](#usage)
- [Author](#author)

## General View

This project is a API to manager content producers and affiliates sales. 

### The Challennge

Send a file with the sales of the content producers and affiliates, normalize the data and store it in a
relational database. In this case is PostgreSQL.

**The user will be able to**

  - endpoint /api/v1/swagger/index.html
    * See the swagger documentation

  - [POST]  /api/v1/sales/upload
    * Upload a file with the sales of the content producers and affiliates, 
    will normalize the data and store it in a relational database. 
    * Return status 200 with the message "Sales added successfully" if all sales was stored
  Example:
   ```json
   {
     "message": "Sales added successfully"
}
```
- Validations
* The file must be in the format .txt
* The lines in the file must have this format: "12022-03-03T13:12:16-03:00DESENVOLVEDOR FULL STACK      0000155000ELIANA NOGUEIRA"

- Errors Return:
- If file is not in the format .txt. Return Status Code 400 with the message:
```json
{
  "message": "File must be a text/plain"
}
```
- If a line in the file is not in the correct format. Return Status Code 400 with the message:
```json
{
  "message": "Line 12: Incorrect format."
}
```
- If a line in the file dont have a product. Return Status Code 400 with the message:
```json
{
  "message": "Line 3: Product is in incorrect format."
}
```
- If a line have a name empty. Return Status Code 400 with the message:

```json
{
      "message": "Line 3: Name is in incorrect format."
}    
```
- If a line have a product with more than 30 characters. Return Status Code 400 with the message:
```json
{
  "message": "Line 3: Product must have length 30"
}
```
- If a line have a name with more than 20 characters. Return Status Code 400 with the message:
```json
{
  "message": "Line 3: Seller is too long, must be less than 21 characters."
}
```
- if a line have a value with more than 10 characters. Return Status Code 400 with the message:
```json
{
  "message": "Line 3:  Value must have 10 numbers."
}
```
  - [GET] /api/v1/sales
    * Return all sales stored in the database
    * Return status 200 with empty array if there is no sales stored
    * Return status 200 and a json with all sales
    
  Example:
  ```json
[
	{
		"product": "CURSO DE GOLANG",
		"producer": "ALBERTO PAZ",
		"affiliate": "",
		"value": 12750,
		"commission": 0,
		"date": "2022-01-15T19:20:30-03:00"
	 },
	 {
		"product": "CURSO DE BEM-ESTAR",
		"producer": "THIAGO OLIVEIRA",
		"affiliate": "JOSE CARLOS",
		"value": 12750,
		"commission": 4500,
		"date": "2022-01-16T14:13:54-03:00"
     }
  ]
```
## The development process

Optei colocar nas tabelas product e affiliates somente os nomes dos mesmos e o id do producer porque são
os únicos juntos com o date que sao imutaveis ao longo do tempo, pois poderia ter um mesmo produto de um producer com 
diferente values, o mesmo vale para o affiliate. Um affiliate poderia ter uma comissoes diferentes
ao longo do tempo caso o arquivo enviado fosse toda a base de dados existente grande seria a possibilidade de 
ter alteração de valores de comissão e o valor do produto.

### Tools Used

- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Swagger](https://swagger.io/)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [Docker](https://www.docker.com/)

## Author

- LinkedIn - [Alberto Ouverney Paz](https://www.linkedin.com/in/beto-ouverney-paz/)
