# Тестовое задание на стажировку Avito 2023
## Launch
To build the project, use the command: docker-compose up --build
## Methods
### User Requests
+ POST http://localhost:8080/api/v1/create_user
+ DEL http://localhost:8080/api/v1/delete_user
+ GET http://localhost:8080/api/v1/get_segments/{id}
+ POST http://localhost:8080/api/v1/add_segments/{id}
+ DEL http://localhost:8080/api/v1/remove_segments/{id}
### Segment Requests
+ POST http://localhost:8080/api/v1/create_segment
+ DEL http://localhost:8080/api/v1/delete_segment/{id}
