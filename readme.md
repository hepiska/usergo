# User go

just user crud plus auth write in go


 
#enpoint
|  url |method   |query   |body   |   |
|---|---|---|---|---|
|/login  |POST   | -   | email, password  |   |
| /signup  |   POST| -  | email, password , address, name |   |
|  /me |  GET | -  |-   | |
|  /user | GET   | search,skip,limit   |-   |   |
|  /user/:id | GET   | -   |-   |   |
|  /user/:id | DELETE   | -   |-   |   |
|  /user/:id | PUT   | -   |address, name  |   |