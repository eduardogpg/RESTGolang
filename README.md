###Methods

> curl -X GET http://localhost:3000/api/v1/users/
> curl -X GET http://localhost:3000/api/v1/users/1 -i
> curl -i -H "Content-Type: application/json" -X POST -d '{"username":"New User", "password": "password123"}' http://localhost:3000/api/v1/users/


> curl -i -H "Content-Type: application/json" -X PUT -d '{"username":"Lalo", "password": "change123"}' http://localhost:3000/api/v1/users/2 

> curl -i -H "Content-Type: application/json" -X PUT -d '{"username":"Lalo", "password": "change123"}' http://localhost:3000/api/v1/users/3

> curl -X DELETE http://localhost:3000/api/v1/users/1 -i