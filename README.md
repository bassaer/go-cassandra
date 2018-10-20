# go-cassandra

```
cqlsh> select * from mykeyspace.users;

 user_id | fname | lname
---------+-------+-------
    1745 |  john | smith
    1744 |  john |   doe
    1746 |  john | smith

(3 rows)
```

```
❯ go run main.go
1745 john smith
1744 john doe
1746 john smith
```
