# clone-todo-mate
[todo mate](https://www.todomate.net/#/) 서비스의 백엔드를 클론 코딩해 보는 레파지스토리입니다.

## Tech Stack:
- **Golang**
  - `gofiber/fiber`
  - `ent/ent`
  - `google/uuid`
  - `golang-jwt/jwt`
- **DataBase:**
  - SQLite

## ToDo
### DataBase
- `Schema`
  - [X] 새로운 `User` 스키마 생성
  - [X] 새로운 `ToDo` 스키마 생성
- `Create`
  - [X] 새로운 유저 생성
  - [ ] 새로운 할일 생성
- `Read`
  - [ ] 생성되어 있는 할일 불러오기
- `Update`
  - [ ] 생성되어 있는 할일 수정
  - [ ] 생성되어 있는 할일 업데이트
- `Delete`
  - [ ] 생성되어 있는 유저 삭제
  - [ ] 생성되어 있는 할일 삭제
## Auth
- [ ] 새로운 JWT 토큰 발행

## Dcos - REST API
### POST `/api/auth/join`
#### Reqeuset
```json
{
    "user_email": "me@hyunsang.dev",
    "user_password": "1q2w3e4r",
    "user_nickname": "HyunSang Park"
}
```
#### Response
```json
{
    "code": "success",
    "status_code": 200,
    "success": true,
    "message": "어서오세요. 하루를 더 체계적으로 계획적인 하루를 보내 봐요.",
    "data": {
        "id": 1,
        "user_uuid": "99aefc46-df76-4fcf-99fe-e6949106f8b1",
        "user_email": "me@hyunsang.dev",
        "user_password": "$2a$10$SEDy7nuaBtxiR6H0bKPcde7MzGfuedM/y0fO9u9G.aqFfPuUF3VEO",
        "user_nickname": "HyunSang Park",
        "updated_at": "2022-12-10T16:30:38.548958+09:00",
        "created_at": "2022-12-10T16:30:38.548958+09:00"
    },
    "responded_at": "2022-12-10T16:30:38.550683+09:00"
}
```