# projector-hsa-08 Web Servers

## How to launch

Execute `docker compose up --build` and that's all.

## Report

General Count:

```
SELECT COUNT(*) FROM users;
```

**40000000**


1. Select all adult users;

```
SELECT * from users where year(now()) - year(birth_date) > 18
```

(execution: 11 s, fetching: 42 ms)

2. Select users from 15 to 18 years;

```
SELECT * from users where year(now()) - year(birth_date)  between 15 and 18;
```

(execution: 11 s, fetching: 23 ms)

3. Find exactly one user by its birth_date

```
SELECT * from users where birth_date = '2008-06-19 22:19:30'
```

(execution: 19 s 353 ms, fetching: 9 ms)

### Btree index:

``
CREATE INDEX birth_date_idx ON users (birth_date);
``

It took around 4 mins to apply index; Lets make some queries:

1. Select all adult users;

```
SELECT * from users where year(now()) - year(birth_date) > 18
```

(execution: 10 ms, fetching: 42 ms)

2. Select users from 15 to 18 years;

```
SELECT * from users where year(now()) - year(birth_date)  between 15 and 18;
```

(execution: 3 ms, fetching: 23 ms)

3. Find exactly one user by its birth_date

```
SELECT * from users where birth_date = '2008-06-19 22:19:30'
```

(execution: 3 ms, fetching: 17 ms)

### Hash index

Its not possible to create Hash index explicitly in MySQL. MySQL use it under the hood for tiny table that fits
to the memory entirely, to have some performance benefits.

```
If a table fits almost entirely in main memory, a hash index speeds up queries by enabling direct lookup of any element, turning the index value into a sort of pointer. InnoDB has a mechanism that monitors index searches. If InnoDB notices that queries could benefit from building a hash index, it does so automatically.
```

https://dev.mysql.com/doc/refman/8.0/en/innodb-adaptive-hash.html

In all other cases, BTree index is better then Hash index, so MySQL use it by default.
