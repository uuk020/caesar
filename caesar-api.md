---
title: 凯撒密码API v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# 凯撒密码API

> v1.0.0

Base URLs:

* <a href="http://localhost:8888/api">凯撒密码API: http://localhost:8888/api</a>

**注意事项**
- /auth 开头 需要登录后获取 token 后操作, 在 Headers Authorization 头上 Bearer <Token>

# 用户

## GET 验证码

GET /captcha

> 返回示例

> 200 Response

```json
{
  "id": "string",
  "image": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|string|true|none|验证码ID|none|
|» image|string|true|none|验证码|base 编码|

## POST 注册用户

POST /register

> Body 请求参数

```json
{
  "user_name": "string",
  "real_name": "string",
  "email": "string",
  "password": "string",
  "re_password": "string",
  "main_password": "string",
  "phone": "string",
  "captcha": "string",
  "captcha_id": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» user_name|body|string| 是 | 用户名|none|
|» real_name|body|string| 是 | 姓名|none|
|» email|body|string| 是 | 邮箱|none|
|» password|body|string| 是 | 密码|none|
|» re_password|body|string| 是 | 重复密码|none|
|» main_password|body|string| 是 | 主密码|none|
|» phone|body|string| 是 | 手机|none|
|» captcha|body|string| 是 | 验证码|none|
|» captcha_id|body|string| 是 | 验证码|none|

> 返回示例

> 200 Response

```json
{
  "id": 0,
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|true|none|用户ID|none|
|» message|string|true|none|消息|none|

## PUT 发送邮件

PUT /email

用于激活账号或重置密码

> Body 请求参数

```json
{
  "email": "string",
  "type": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» email|body|string| 是 | 邮箱|none|
|» type|body|string| 是 | 类型|none|

> 返回示例

> 200 Response

```json
{
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none|消息|none|

## POST 激活账号

POST /activation

> Body 请求参数

```json
{
  "email": "string",
  "code": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» email|body|string| 是 | 邮箱|none|
|» code|body|string| 是 | 激活码|none|

> 返回示例

> 200 Response

```json
{
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none|消息内容|none|

## POST 重置密码

POST /password

> Body 请求参数

```json
{
  "email": "string",
  "code": "string",
  "password": "string",
  "re_password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» email|body|string| 是 | 邮箱|none|
|» code|body|string| 是 | 验证码|none|
|» password|body|string| 是 | 密码|none|
|» re_password|body|string| 是 | 重复密码|none|

> 返回示例

> 200 Response

```json
{
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none|消息内容|none|

## POST 登录

POST /login

> Body 请求参数

```json
{
  "user_name": "string",
  "password": "string",
  "captcha": "string",
  "captcha_id": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» user_name|body|string| 是 | 用户名或邮箱|none|
|» password|body|string| 是 | 密码|none|
|» captcha|body|string| 是 | 验证码|none|
|» captcha_id|body|string| 是 | 验证码ID|none|

> 返回示例

> 200 Response

```json
{
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» token|string|true|none|登录 token|none|

## DELETE 注销账号

DELETE /auth/logout

> 返回示例

> 200 Response

```json
{
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none|消息内容|none|

## GET 个人资料

GET /auth/me

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "email": "string",
  "phone": "string",
  "real_name": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» email|string|true|none|邮箱|none|
|» phone|string|true|none|手机|none|
|» real_name|string|true|none|姓名|none|

## PUT 更新个人资料

PUT /auth/me

> Body 请求参数

```json
{
  "real_name": "string",
  "email": "string",
  "phone": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» real_name|body|string| 是 ||none|
|» email|body|string| 是 ||none|
|» phone|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none|消息内容|none|

# 平台账号

## POST 新建平台账号信息

POST /auth/account

> Body 请求参数

```json
{
  "name": "string",
  "email": "string",
  "password": "string",
  "main_password": "string",
  "platform": "string",
  "url": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 用户名|none|
|» email|body|string| 是 | 邮箱|none|
|» password|body|string| 是 | 密码|none|
|» main_password|body|string| 是 | 主密码|none|
|» platform|body|string| 是 | 平台|none|
|» url|body|string| 是 | 链接|none|

> 返回示例

> 200 Response

```json
{
  "id": 0,
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|true|none|平台ID|none|
|» message|string|true|none|消息内容|none|

## GET 平台列表

GET /auth/account

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|page|query|string| 是 ||页数|
|page_size|query|string| 是 ||页码|
|platform|query|string| 否 ||平台, 支持模糊搜索|
|date_start|query|string| 否 ||开始时间|
|date_end|query|string| 否 ||结束时间|

> 返回示例

> 200 Response

```json
{
  "data": [
    {
      "id": 0,
      "user_id": 0,
      "name": "string",
      "email": "string",
      "password": "string",
      "platform": "string",
      "url": "string",
      "created_at": 0,
      "updated_at": 0
    }
  ],
  "count": 0,
  "page": 0,
  "page_size": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|列表内容|none|
|»» id|integer|true|none|平台ID|none|
|»» user_id|integer|true|none|用户ID|none|
|»» name|string|true|none|账号名|加密|
|»» email|string|true|none|邮箱|加密|
|»» password|string|true|none|密码|加密|
|»» platform|string|true|none|平台|none|
|»» url|string|true|none|链接|none|
|»» created_at|integer|true|none|创建时间|时间戳|
|»» updated_at|integer|true|none|更新时间|时间戳|
|» count|integer|true|none|总数|none|
|» page|integer|true|none|页数|none|
|» page_size|integer|true|none|页码|none|

## POST 查看平台账号

POST /auth/account/{id}

> Body 请求参数

```json
{
  "main_password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|path|string| 是 ||平台ID|
|body|body|object| 否 ||none|
|» main_password|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "email": "string",
  "name": "string",
  "password": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» email|string|true|none|邮箱|none|
|» name|string|true|none|账号|none|
|» password|string|true|none|密码|none|

## PUT 更新平台账号

PUT /auth/account/{id}

> Body 请求参数

```json
{
  "main_password": "string",
  "name": "string",
  "email": "string",
  "password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» main_password|body|string| 是 | 主密码|none|
|» name|body|string| 是 | 账号名|none|
|» email|body|string| 是 | 邮箱|none|
|» password|body|string| 是 | 密码|none|

> 返回示例

> 200 Response

```json
{
  "email": "string",
  "id": 0,
  "name": "string",
  "password": "string",
  "updated_at": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» email|string|true|none|加密邮箱|none|
|» id|integer|true|none|平台ID|none|
|» name|string|true|none|加密用户名|none|
|» password|string|true|none|加密密码|none|
|» updated_at|integer|true|none|更新时间|时间戳|

## DELETE 删除平台账号

DELETE /auth/account/{id}

> Body 请求参数

```json
{
  "main_password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» main_password|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none|消息内容|none|

## GET 查看操作日志

GET /auth/account/{id}/log

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|path|string| 是 ||none|
|page|query|integer| 是 ||页数|
|page_size|query|integer| 是 ||页码|

> 返回示例

> 200 Response

```json
{
  "data": [
    {
      "id": 0,
      "account_id": 0,
      "type": 0,
      "created_at": 0,
      "updated_at": 0
    }
  ],
  "count": 0,
  "page": 0,
  "page_size": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|日志内容|none|
|»» id|integer|true|none|ID|none|
|»» account_id|integer|true|none|平台ID|none|
|»» type|integer|true|none|类型|0-创建 1-查看 2-编辑|
|»» created_at|integer|true|none|创建时间|时间戳|
|»» updated_at|integer|true|none|更新时间|时间戳|
|» count|integer|true|none|总数|none|
|» page|integer|true|none|页数|none|
|» page_size|integer|true|none|页码|none|

# 数据模型

