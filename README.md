# go ddd api server
- [参考](https://mintaku-blog.net/go-ddd/)
- [mysql接続文字列を直接書かないようにする](https://zenn.dev/mstn_/articles/75667657fa5aed)

## dbの準備
```bash
docker run -it --rm --name go-ddd-mysql -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 -d mysql:latest
docker exec -it go-ddd-mysql mysql -p
```

```sql
create database test_db;
create table test_db.users (id int PRIMARY KEY AUTO_INCREMENT, name varchar(10), created_at datetime DEFAULT CURRENT_TIMESTAMP);
insert into test_db.users (name) values ('taro');
```

## 動作確認
```bash
# get
curl http://localhost:8080/api/users -X GET

# post
curl http://localhost:8080/api/users -X POST -H 'Content-Type: application/json' -d '{"name": "miko"}'
```