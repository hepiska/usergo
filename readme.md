# User go

just user crud plus auth write in go


 
#enpoint
|  url |method   |query   |body   |   |
|---|---|---|---|---|
|/login  |POST   | -   | email, password  |   |
| /signup  |   POST| -  | email, password , address, name |   |
|  /me |  GET | -  |-   | |
|  /users | GET   | search,skip,limit   |-   |   |
|  /users/:id | GET   | -   |-   |   |
|  /users/:id | DELETE   | -   |-   |   |
|  /users/:id | PUT   | -   |address, name  |   |

