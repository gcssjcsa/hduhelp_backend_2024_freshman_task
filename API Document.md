## **API Document**

### Users

**用户类型结构**
```
type User struct {
	Id         int
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       Role
	Email      string `json:"email"`
	CreateDate string
}
```

**权限类型**
```
type Role int

const (
	Admin Role = iota
	Student
	Guest
)
```

**POST 账户注册**  

`POST /register`  

注册账户，要求用户名和邮箱不大于63个字符，用户名唯一，邮箱符合admin@example.com格式

> 请求参数

```
{
    "username": string,
    "password": string,
    "email": string
}
```
|名称|类型|必选|说明|
|----|----|----|----|
|username|string|是|不超过63个字符|
|password|string|是| |
|email|string|是|不超过63个字符|

> 返回结果  

200 OK  
- 注册成功：`{"status": "Register sucessfully!"}`-

400 Bad Request  
- 用户名或密码为空：`{"error": "Username or password is empty"}`  
- 用户名重复：`{"error": "User already exists"}`  
- 用户名超过63个字符：`{"error": "Username is too long"}`  
- 邮箱超过63个字符：`"error": "Email is too long"`  
- 邮箱不符合正则判定：`{"error": "Email is invalid"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**POST 账户登录**

`POST /login`

> 请求参数

```
{
    "username": string,
    "password": string
}
```
|名称|类型|必选|说明|
|----|----|----|----|
|username|string|是|不超过63个字符|
|password|string|是| |

> 返回结果  

200 OK  
- 登录成功：`{"status": "Login sucessfully!"}`  

400 Bad Request  
- 用户名或密码为空：`{"error": "Username or password is empty"}`  
- 用户名或密码错误：`{"error": "Invalid username or password"}`  

500 Internal Server Error  
- 生成token失败：`{"error": "Couldn't generate token"}`  
- 服务器错误：`{"error": "Internal server error"}`

**POST 账户登出**  

`POST /logout`

设置token过期以达到登出效果

> 返回结果

200 OK
- 登出成功：`{"status": "Logout sucessfully!"}`

**POST 获取账户信息**

`POST /user`

需登录，依据token中的id查询信息，未登录时返回Role信息（为Guest）

> 返回结果

200 OK
- 获取成功：
	```
	{
		"id": string,
		"username": string,
		"role": Role,
		"email": string
	}
	```

401 Unauthorized
- 未登录：
	```
	{
		"error": "You haven't logged in",
		"role": Role
	}
	```

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**PUT 修改账户信息**

`PUT /user`

需登录，可修改用户名和邮箱，要求同注册

> 请求参数

```
{
    "username": string,
    "email": string
}
```
|名称|类型|必选|说明|
|----|----|----|----|
|username|string|是|不超过63个字符|
|email|string|是|不超过63个字符|

> 返回结果

200 OK
- 修改成功：`{"status": "Update sucessfully!"}`

400 Bad Request  
- 用户名或邮箱为空：`{"error": "Info can not be empty"}`  
- 用户名重复：`{"error": "User already exists"}`  
- 用户名超过63个字符：`{"error": "Username is too long"}`  
- 邮箱超过63个字符：`"error": "Email is too long"`  
- 邮箱不符合正则判定：`{"error": "Email is invalid"}`

401 Unauthorized
- 未登录：`{"error": "You haven't logged in"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**PATCH 修改账户密码**

`PATCH /user`

需登录，若更改密码成功则会清除token重新登录

> 请求参数

```
{
	"originalPwd": string,
	"newPwd": string
}
```
|名称|类型|必选|说明|
|----|----|----|----|
|originalPwd|string|是|原密码|
|newPwd|string|是|新密码|

> 返回结果

200 OK
- 修改成功：`{"status": "Change password sucessfully!"}`

400 Bad Request  
- 请求参数有误：`{"error": "Invalid data"}`

401 Unauthorized
- 未登录：`{"error": "You haven't logged in"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**DELETE 删除账户**

`DELETE /user`

需登录，删除后清除token

> 返回结果

200 OK
- 删除成功：`{"Status": "Delete sucessfully! Bye!"}`

401 Unauthorized
- 未登录：`{"error": "You haven't logged in"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

### Questions

**问题类型结构**
```
type Question struct {
	Id         int    `json:"id"`
	Author     string `json:"author"`
	Permission int    `json:"permission"`
	AuthorId   int    `json:"author_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	PostDate   string `json:"post_date"`
	ModifyDate string `json:"modify_date"`
	Likes      int    `json:"likes"`
}
```

**POST 创建问题**

`POST /api/question`

需登录

> 请求参数

```
{
	"title": string,
	"content": string,
	"permission": int,
}
```
|名称|类型|必选|说明|
|----|----|----|----|
|title|string|是|问题标题|
|content|string|是|问题内容|
|permission|int|是|可见范围，小于等于该值的用户都可见|

> 返回结果

200 OK
- 创建成功：`{"status": "Create question successfully"}`

400 Bad Request  
- 标题或内容为空：`{"error": "Title or content is empty"}`

401 Unauthorized  
- 用户未登录：`{"error": "You haven't logged in"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**GET 获取问题**

`GET /api/question/:qid`

> 返回结果

200 OK  
- 获取成功，用户未登录：
	```
	{
		"question": 
		{
    		"id": int,
    		"author": string,
    		"permission": int,
    		"author_id": int,
    		"title": string,
    		"content": string,
    		"post_date": "2006-01-02 15:04:05",
    		"modify_date": "2006-01-02 15:04:05",
    		"likes": int
  		},
		"userRole": int
	}
	```
- 获取成功，用户已登录：
	```
	{
		"question": 
		{
    		"id": int,
    		"author": string,
    		"permission": int,
    		"author_id": int,
    		"title": string,
    		"content": string,
    		"post_date": "2006-01-02 15:04:05",
    		"modify_date": "2006-01-02 15:04:05",
    		"likes": int
  		},
		"userId": int,
		"userRole": int
	}
	```

400 Bad Request  
- qid不为数字：`{"error": "Id must be integer"}`

403 Forbidden
- 问题权限小于用户Role权限：`{"error": "You do not have permission to view this question"}`  

404 Not Found
- 问题不存在：`{"error": "question not found"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**PUT 修改问题**

`PUT /api/question/:qid`

需登录

> 请求参数

```
{
	"id": int,
	"title": string,
	"content": string,
	"permission": int
}
```
|名称|类型|必选|说明|
|----|----|----|----|
|id|int|是|问题id|
|title|string|是|问题标题|
|content|string|是|问题内容|
|permission|string|是|问题可见范围|

> 返回结果

200 OK  
- 修改成功：`{"status": "Update question successfully"}`

400 Bad Request  
- qid不为数字：`{"error": "Id must be integer"}`
- 标题或内容为空：`{"error": "Title or content is empty"}`

401 Unauthorized  
- 用户未登录：`{"error": "You haven't logged in"}`

403 Forbidden
- 问题并非该用户创建且用户并非管理员：`{"error": "You do not have permission to modify this question"}`

404 Not Found
- 问题不存在：`{"error": "question not found"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**DELETE 删除问题**

`DELETE /api/question/:qid`

需登录

> 返回结果

200 OK  
- 删除成功：`{"status": "Delete question successfully"}`

400 Bad Request  
- qid不为数字：`{"error": "Id must be integer"}`

401 Unauthorized  
- 用户未登录：`{"error": "You haven't logged in"}`

403 Forbidden  
- 问题并非该用户创建：`{"error": "You do not have permission to delete this question"}`

404 Not Found
- 问题不存在：`{"error": "Question not found"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**GET 获取可见问题列表**

`GET /api/question/pblist`

依据中间件返回的用户Role类型选取可见的问题列表

> 返回结果

200 OK  
- 获取成功：
	```
	[
  		{
    		"id": int,
    		"author": string,
    		"permission": int,
    		"author_id": int,
    		"title": string,
    		"content": string,
    		"post_date": "2006-01-02 15:04:05",
    		"modify_date": "2006-01-02 15:04:05",
    		"likes": int
  		}, ...
	]
	```

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**GET 获取登录用户问题列表**

`GET /api/question/pvlist`

需登录，依据中间件返回的用户id选取自己发布的问题列表

> 返回结果

200 OK  
- 获取成功：
	```
	[
  		{
    		"id": int,
    		"author": string,
    		"permission": int,
    		"author_id": int,
    		"title": string,
    		"content": string,
    		"post_date": "2006-01-02 15:04:05",
    		"modify_date": "2006-01-02 15:04:05",
    		"likes": int
  		}, ...
	]
	```

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

### Answers

**回答类型结构**
```
type Answer struct {
	Id         int    `json:"id"`
	Author     string `json:"author"`
	Permission int    `json:"permission"`
	AuthorId   int    `json:"author_id"`
	QuestionId int    `json:"question_id"`
	Content    string `json:"content"`
	PostDate   string `json:"post_date"`
	ModifyDate string `json:"modify_date"`
	Likes      int    `json:"like"`
	Dislikes   int    `json:"dislike"`
	IsBest     bool   `json:"is_best"`
}
```

**POST 发布回答**

`POST /api/question/:qid/answer`

需登录

> 请求参数

```
{
	'content': string
}
```
|名称|类型|必选|说明|
|----|----|----|----|
|content|string|是|回答内容|

> 返回结果

200 OK  
- 发布成功：`{"status": "Create answer successfully"}`

400 Bad Request  
- qid不为数字：`{"error": "Id must be integer"}`
- 内容为空：`{"error": "Content is empty"}`

401 Unauthorized  
- 用户未登录：`{"error": "You haven't logged in"}`

404 Not Found
- 问题不存在：`{"error": "Question not found"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**GET 获取回答**

`GET /api/question/:qid/answer/:aid`

> 返回结果

200 OK  
- 获取成功：
	```
	{
      "id": int,
      "author": string,
      "permission": int,
      "author_id": int,
      "question_id": int,
      "content": string,
      "post_date": "2006-01-02 15:04:05",
      "modify_date": "2006-01-02 15:04:05",
      "like": int,
      "dislike": int,
      "is_best": bool
    }
	```

400 Bad Request  
- aid不为数字：`{"error": "Id must be integer"}`

403 Forbidden
- 权限不足：`{"error": "You do not have permission to view this answer"}`

404 Not Found
- 问题不存在：`{"error": "Answer not found"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**PUT 修改回答**

`PUT /api/question/:qid/answer/:aid`

需登录

> 请求参数

```
{
	"id": int,
	'content': string
}
```
|名称|类型|必选|说明|
|----|----|----|----|
|id|int|是|问题id(aid)|
|content|string|是|问题内容|

> 返回结果

200 OK  
- 修改成功：`{"status": "Update answer successfully"}`

400 Bad Request  
- 内容为空：`{"error": "Content is empty"}`

401 Unauthorized  
- 用户未登录：`{"error": "You haven't logged in"}`

403 Forbidden
- 回答并非该用户创建且用户并非管理员：`{"error": "You do not have permission to modify this answer"}`

404 Not Found
- 问题不存在：`{"error": "Answer not found"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**DELECT 删除回答**

`DELETE /api/question/:qid/answer/:aid`

需登录

> 返回结果

200 OK  
- 删除成功：`{"status": "Delete answer successfully"}`

400 Bad Request  
- aid不为数字：`{"error": "Id must be integer"}`

401 Unauthorized  
- 用户未登录：`{"error": "You haven't logged in"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`

**GET 获取回答列表**

`GET /api/question/:qid/answer`

> 返回结果

200 OK  
- 获取成功，用户未登录：
	```
	{
		"answer":
		{
      		"id": int,
      		"author": string,
      		"permission": int,
      		"author_id": int,
      		"question_id": int,
      		"content": string,
      		"post_date": "2006-01-02 15:04:05",
      		"modify_date": "2006-01-02 15:04:05",
      		"like": int,
      		"dislike": int,
      		"is_best": bool
    	},
		"userId": -1,
		"role": 2
	}
	```
- 获取成功，用户已登录：
	```
	{
		"answer":
		{
      		"id": int,
      		"author": string,
      		"permission": int,
      		"author_id": int,
      		"question_id": int,
      		"content": string,
      		"post_date": "2006-01-02 15:04:05",
      		"modify_date": "2006-01-02 15:04:05",
      		"like": int,
      		"dislike": int,
      		"is_best": bool
    	},
		"userId": int,
		"role": int
	}
	```

400 Bad Request  
- qid不为数字：`{"error": "Id must be integer"}`

403 Forbidden
- 用户无权限：`{"error": "You do not have permission to view this question"}`

500 Internal Server Error  
- 服务器错误：`{"error": "Internal server error"}`