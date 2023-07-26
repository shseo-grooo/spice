# Spice V3

가독성
한 함수를 여러 함수로 쪼개야 할 때는 domain으로 빼자.

### 변경점
1. interface 사용금지
2. 폴더구조 단순화
3. orm : gorm -> entgo
4. router : 한 곳에서 관리, toString 지원
5. error : nil 체크 제거 -> panic으로 처리할 것.
6. serializer : client에서 요청한 데이터만 반환한다.

### entgo 명령어
```bash
$ go run -mod=mod entgo.io/ent/cmd/ent new 객체명
$ go generate ./ent
```

### entgo 플러그인 명령어
1. init : 기본 폴더구조 생성
2. go generate : ent schema, controller, domain, repository, service, config, test CRUD 자동생성

### rollback 처리
### soft_delete
### join
### serializer
### go generate