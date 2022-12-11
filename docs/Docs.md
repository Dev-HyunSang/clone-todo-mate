## Docs clone-todo-mate 

## DataBase with SQLite
### `ToDo`
```go
// Fields of the ToDo.
func (ToDo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("todo_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("todo_context").
			Default("null"),
		field.Bool("completion").
			Default(false),
		field.Time("completed_at").
			Default(nil), // 완료하지 않으면 표시하지 않음.
		field.Time("created_at").
			Default(time.Now().AddDate(2006, 01, 02)), // Year-Mouth-Day
		field.Time("edited_at").
			Default(time.Now().AddDate(2006, 01, 02)),
	}
}
```

### `Users`
```go
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("user_email").
			Default("null"),
		field.String("user_password").
			Default("null"),
		field.String("user_nickname").
			Default("null"),
		field.Time("updated_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}
```

## Flowchart
### `cmd.JoinUser` - POST `/api/auth/join`
1. 사용자로부터 회원가입 요청을 받습니다.
   - 에러 발생시 `painc()`
2. 새로운 UUID 발행합니다.
   - 에러 발생시 JSON 구조체(`models.ErrResp{}`)으로 오류 응답합니다. 
3. 입력한 패스워드를 암호화 합니다.
   -  에러 발생시 JSON 구조체(`models.ErrResp{}`)으로 오류 응답합니다.
4. 사용자 구조체를 이용하여 정의하며 데이터베이스에 사용자 정보를 넣습니다.
   - 에러 발생시 JSON 구조체(`models.ErrResp{}`)로 오류 응답합니다.
5. 이미 가입된 계정 중 동일한 메일로 가입된 계정을 찾아봅니다.
   - 동일한 계정 발생 시 사용자에게 동일한 메일 있다고 응답합니다.
   - 에러 발생시  JSON 구조체(`models.ErrResp{}`)로 오류 응답합니다.
6. 위 기능이 완료되면 응답합니다.
   
### `cmd.LoginUser` - POST `/api/auth/login`
1. 사용자로부터 로그인 요청을 받습니다.
2. 사용자가 입력한 데이터를 토대로 데이터베이스에서 메일을 찾으며, 정보를 가지고 옵니다.
3. 가지고 온 정보 중에서 비밀번호를 가지고 오며, 비밀번호를 해쉬화하며 확인합니다.
4. JWT(JSON Web Token)를 발행합니다.
5. `c.cookie`로 JWT(JSON Web Token)를 추가합니다.
6. 위 기능이 완료되면 응답합니다.

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

### POST `/api/auth/login`
#### Request
```json
{
    "user_email": "me@hyunsang.dev",
    "user_password": "1q2w3e4r"
}
```

#### Response
```
{
    "code": "success",
    "status_code": 200,
    "success": true,
    "message": "",
    "Data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzA3NTA3ODYsInVzZXJfbmlja25hbWUiOiJIeXVuU2FuZyBQYXJrIiwidXNlcl91dWlkIjoiNTllYzk2NzUtYzE5Mi00ZDdjLWExYjUtNDk0Yjk5MWE2ZGRkIn0.WuxH1UGfRpEOirZ9FdmFsTqibx8UlhlnLDg_CXSBM_0"
    },
    "resoponded_at": "2022-12-11T17:26:26.421505+09:00"
}
```

### POST `/api/todo/create`
#### Request
```
<!-- JWT Cookie -->
jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzA3NTA3ODYsInVzZXJfbmlja25hbWUiOiJIeXVuU2FuZyBQYXJrIiwidXNlcl91dWlkIjoiNTllYzk2NzUtYzE5Mi00ZDdjLWExYjUtNDk0Yjk5MWE2ZGRkIn0.WuxH1UGfRpEOirZ9FdmFsTqibx8UlhlnLDg_CXSBM_0; Path=/;
```

```json
{
    "todo_context": "Hello World!",
    "todo_completion": false
}
```

#### Response
```json
{
    "code": "success",
    "status_code": 200,
    "success": true,
    "message": "성공적으로 할일을 생성했어요!",
    "responded_at": "2022-12-11T17:18:02.317328+09:00"
}
```