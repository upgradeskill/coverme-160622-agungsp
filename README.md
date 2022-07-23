[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-c66648af7eb3fe8bc4f294546bfd86ef473780cde1dea487d3c4ff354943c9ae.svg)](https://classroom.github.com/online_ide?assignment_repo_id=8046162&assignment_repo_type=AssignmentRepo)

# Task 2 - Majoo Golang Bootcamp

<hr>

## Build repo to executable file

`goexec build`
<small style="color: red;"><i>\*This command only for Windows users</i></small>

## Build and start server

`goexec run`
<small style="color: red;"><i>\*This command only for Windows users</i></small>

## Running the test

`goexec test`
<small style="color: red;"><i>\*This command only for Windows users</i></small>

## Endpoint list

> **Get all data**
> [![GET - /products](https://img.shields.io/badge/GET-%2Fproducts-2ea44f?style=for-the-badge)](#)

> **Get data by ID**
> [![GET - /products/{id}](https://img.shields.io/badge/GET-%2Fproducts%2F{id}-2ea400?style=for-the-badge)](#)

> **Create data**
> [![POST - /products](https://img.shields.io/badge/POST-%2Fproducts-2ea4f0?style=for-the-badge)](#)
>
> ```
> {
>    "name": "mouse",
>    "unit": "pcs",
>    "price": 75000,
>    "stock": 2
> }`
> ```

> **Update data**  
> [![PUT - /products/{id}](https://img.shields.io/badge/PUT-%2Fproducts%2F{id}-2ea4A0?style=for-the-badge)](#)
>
> ```
> {
>    "name": "pensil",
>    "unit": "pcs",
>    "price": 5000,
>    "stock": 10
> }`
> ```

> **Delete data**
> [![DELETE - /products/{id}](https://img.shields.io/badge/DELETE-%2Fproducts%2F{id}-FF1212?style=for-the-badge)](#)
