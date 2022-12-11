# clone-todo-mate
[todo mate](https://www.todomate.net/#/) 서비스의 백엔드를 클론 코딩해 보는 레파지스토리입니다.  
**개발한 내용에 대해서 [`Docs.md`](./docs/Docs.md) 확인하실 수 있으십니다.**

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