# backend-code

this is backen for Fflix that manage api for CRUD for user.

###  How to use it?

build
```go
cd backend-code
dep ensure
go build main
```

run
```go
./backend-code
```

### Backend Path

create new user
```
POST http://localhost:3000/v1/account/register

json body:
{
    "email":"xxx.yyy.zzz@gmail.com"
}
======================================================
Return:
status : 201
body :
{
  "message": "account created",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImhhcmlzLmRlLm1vcmF0YUBnbWFpbC5jb20iLCJleHAiOjE1NjcwOTg5NjJ9.tIGRLE0MK6aZ9IGv3-87_zFnriYAVGhME0GgPibJaLk"
}
```
login user
```
POST http://localhost:3000/v1/account/login

json body:
{
    "email":"xxx.yyy.zzz@gmail.com",
    "password":"aaaaaa"
}
=======================================================
status : 202
return :
{
  "message": "logged in",
  "status": false,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjAsIk5hbWUiOiIiLCJFbWFpbCI6ImFiZHVsLmhhcmlzLmRqYWZhckBnbWFpbC5jb20iLCJleHAiOjE1NjkxNTk4NTJ9.HkNJyFLJyoXH9Aiwkw_BLm4hr1voF15coEWizz7Femk",
  "user": {
    "Email": "xxx.yyy.zzz@gmail.com",
    "Password": "$2a$10$Pa3NTRtDCQj43HEHdwH5S.cxETmYK.nRoO//laazijNODZNlPbxVW"
  }
}
```
update user
```
PUT http://localhost:3000/auth/v1/account

json body:
{
	"email":"haris.de.morata@gmail.com",
	"password":"robonson11",
	"fullname":"abdul haris jafar",
	"birthday":"27/27/2009",
	"gender":"Male",
	"country":"country",
	"plan":"Free",
	"newsletter":""
}
=======================================================
Return:
{
  "message": "User Updated"
}
status : 200
header :
    x-access-token : Token from login
```
Delete user
```
POST http://localhost:3000/auth/v1/account/delete

json body:
{
    "email":"xxx.yyy.zzz@gmail.com"
}
=======================================================
Return:
{
  "message": "User Updated"
}
status : 200
header :
    x-access-token : Token from login
body :
{
    "password":"robonson11"
}
```
List Movies
```
GET http://localhost:3000/auth/v1/movies?count=number

=======================================================
Return:
[
  {
    "NetflixId": "70236771",
    "Title": "Criminal Justice",
    "Image": "https://occ-0-753-1360.1.nflxso.net/dnm/api/v6/XsrytRUxks8BtTRf9HNlZkW2tvY/AAAABSVph8Jb2ZgT7LZLRw5UVqIeukFtnLA21MUr_MtZp7KNFN3h6cAP6uaXs4j_gJwWyM1ivorAYqvN2SR4f6nezub0Z54KOUY.jpg?r=645",
    "Synopsis": "Each season of this award-winning BBC series follows a single homicide defendant through Britain&#39;s criminal justice system.<br><b>From 1 to 2 Seasons</b><br>2019-07-18 16:00:17",
    "Type": "",
    "Released": "2008",
    "Runtime": "",
    "LargeImage": "",
    "Unogsdate": "2019-07-12",
    "Imbid": "tt1188927",
    "Download": "0",
    "Rating": "7.8"
  },
  .
  .
  .
  .
 ]
status : 200
header :
    x-access-token : Token from login
```
Get Full Account
```
POST http://localhost:3000/auth/v1/account/full
=======================================================
body :
{
    "email":"kotekaman.ahay@uhuy.com"
}
Return:
{
  "Email": "haris.de.morata@gmail.com",
  "Password": "$2a$10$MpxKTkSAfDC4ReoutIztMepxcBBkiXuv1.Z7XXTpiGjdJbL54EqSy",
  "Fullname": "abdul haris jafar",
  "Birthday": "27/27/2009",
  "Gender": "Male",
  "Country": "country",
  "Plan": "Free",
  "Newsletter": ""
}
status : 200
header :
    x-access-token : Token from login


```
Get Refresh Count Data
```
GET http://localhost:3000/auth/v1/refresh?email=xxx@gmail.com
Return:
{
    "refresh":"10"
}
header :
    x-access-token : Token from login
```
edit config.toml
```toml
[postgres]
url="hostdb"
Port="portdb"
db="dbname"
user="dbusr"
password="dbpassword"

[mail]
smtp="smtp.host.com"
port="587"
from="aa.bb.cc@host.com"
password="passwordemail"

```