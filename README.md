# douyin-Goteam
# 项目文档

## 架构设计

![[Pasted image 20230903111618.png|700]](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/695cb697e49c410e9953f0b00a0ccdfc~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp#?w=799&h=636&s=48957&e=png&b=141414)

客户端发送请求给API端，API端对客户进行鉴权，并且把数据进行拆分发送给Service层，Service层的RPC服务分别在ETCD进行注册，API端即可通过ETCD找到相应的RPC服务，再由RPC服务调用数据库和缓存对数据进行增删改查，再返回给API端，API端即可返回数据给客户端

## 技术选型
| 技术栈   | 原因                                                       | 版本    |
| -------- | ---------------------------------------------------------- | ------- |
| Go-Zero  | 集成了web和rpc框架，编写api文件和proto文件快速生成结构代码 | v1.5.4  |
| Nginx    | 简单的配置即可搭建静态资源服务器                           | 1.24.0  |
| Redis    | 优秀的读写效率和缓存能力的非关系型数据库                   | 7.0.12  |
| MySQL    | 具有可靠性的稳定性的关系型数据库                           | 8.0.33  |
| GORM     | 能够简易使用go操作MySQL数据库                              | v1.25.3 |
| Go-Redis | 能够简易使用go操作Redis数据库                              | v6.15.9 |
| ffmpeg   | 用于视频压缩和抽帧                                         |         |

使用Go-Zero搭建API端和Service层，使用GORM操作数据库，Go-Redis操作Redis缓存，Nginx实现图片和视频的静态资源服务器，进行存放和管理，ffmpeg实现视频压缩和截取第一帧作为视频封面

## 数据库设计
![[Pasted image 20230903001224.png|800]](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/5d57b3cf3d7c4201a8dcbeb2cc027794~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp#?w=1142&h=753&s=86508&e=png&b=414149)

### 数据库说明
- user表为用户信息，其中用户密码password做MD5加密，保证即使数据库泄露也无法轻易登录他人的账户（考虑到绝大部分软件有找回密码的功能，这其实缺乏扩展性）
- video表为视频信息，与user表为一对多的关系
- comment表为评论信息，与video表和user表都为一对多的关系
- users_favor_videos表为user是否喜欢相应的video
- fans_follows表为user是否互相关注，其中user不能关注自己

### 设计原因
整个抖音最重要的两大部分为视频和用户，其他的各种功能都围绕这两者展开，所以user和video分别建表，其他功能均可以通过额外建表的方式扩展

## 缓存数据库设计
![[Pasted image 20230902234316.png|800]](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/6419b919819d469eb56df4f955ca4e8c~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp#?w=958&h=696&s=84263&e=png&b=151515)

### 缓存说明
- user_id哈希表通过id查询缓存
- videos按照10个为单位进行分片，每10个创建相应的时间戳存入time
- userid_userid查找相应的消息发送者和接收者（通过条件语句判断，两个好友之间只有一个userid_userid的list表）
**聊天记录和视频是按照时间顺序发送给服务器处理的，利用这点可以使用list表轻松的实现顺序和倒序的返回查询**

### 场景分析
抖音用户一般使用app会有两种操作，一是不断往下刷视频，点赞评论关注，查看用户主页其他视频；二是处理好友分享的视频或消息

以上，用户使用最频繁的功能即是视频和用户模块，如果直接对数据库读写，磁盘的读写速度较慢，并且视频需要按照一定的顺序进行返回，需要设计相应的数据结构。所以对user和video分别建立缓存，并且video使用list结构。
聊天记录只进行缓存，并不写入数据库（redis提供了持久化存储的方式，并且可以设置过期时间，防止堆积在数据库中）同时，一般的聊天功能都是通过时间戳分片返回的，所以后续扩展时间戳分片返回，可以直接仿照视频时间戳分片的方式建立time-list表

### 数据库和缓存的一致性
写入：先操作缓存，修改缓存的数据，再额外创建协程异步处理数据库
读取：先读取缓存，如果未命中则在数据库中查找，写入缓存并且返回数据

## 项目结构
```bash
douyin
|--	apps
	|--	api            对外提供HTTP服务
		|--	user       数据拆分、包装、接收、转发以及身份鉴权
		
	|-- rpc            对内提供RPC服务
		|-- user       用户登录、注册和用户信息功能
		|-- video      视频功能(视频流、评论、点赞)
		|-- relation   社交功能(关注、粉丝、好友)
		|-- chat       聊天功能(发送消息、聊天记录)
		
|--	pkg                公共部分
	|-- authcrypto.go  auth获取token和鉴权
	|-- userdb.go      存放数据库的结构体和公共函数
	
|--	README.md
|--	go.mod
```

## 具体实现
### 具体框架
![[Pasted image 20230903165617.png|800]](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/dddc9ed9eb264d098411ea03ad3d000d~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp#?w=1003&h=550&s=45938&e=png&b=141414)

### Token鉴权
使用HS256算法加密，token的生成和鉴权均由自己编写
```go
// HS256加密部分
mac := hmac.New(sha256.New, []byte(AccessSecret))
secret := base64.URLEncoding.EncodeToString(mac.Sum(nil))

// 验证是否过期
if time.Now().Unix() > payload.Exp {
	return 0, errors.New("token out of data")
}

// 验证是否有效
secret := base64.URLEncoding.EncodeToString(mac.Sum(nil))
if secret != iter[2] {
	return 0, errors.New("token is wrong")
}
```

### 视频压缩和抽帧
使用ffmpeg进行压缩和抽帧
```go
// 压缩30帧 视频码率为2.0MB/s 音频码率为1.5MB/s
cmdplay := exec.Command("ffmpeg", "-i", fornowplay, "-r", "30", "-b:v", "2M", "-b:a", "1.5M", play)
err = cmdplay.Run()

// 抽取视频的第一帧作为图片
cmdpic := exec.Command("ffmpeg", "-i", play, "-y", "-f", "image2", "-vframes", "1", cover)
err = cmdpic.Run()
```

### 视频分片 
```go
// 按照10个一组进行时间戳分片
llen, err := redisDb.LLen("videos").Result()
if llen%10 == 0 {
	redisDb.LPush("time", videoinfo.CreatedAt)
}

// 根据时间戳判断返回的视频位置
var left, right int64
var next_time int64
VLen, err := redisDb.LLen("videos").Result()
TLen, err := redisDb.LLen("time").Result()
for k, v := range timeset {
	t, _ := strconv.ParseInt(v, 10, 64)
	if in.LatestTime >= t {
		right = VLen - 10*(TLen-int64(k)-1) - 1
		if k != 0 {
			left = right - 9
		} else {
			left = 0
		}
		if k != int(TLen-1) {
			n_t, _ := strconv.ParseInt(timeset[k+1], 10, 64)
			next_time = n_t
		} else {
			next_time = time.Now().Unix()
		}
		break
	}
}
```

## 项目总结与反思
### 目前仍存在的问题
-  上传视频压缩需要优化，如果视频有100MB，则可以压缩到5MB以内，如果只有600KB，则会反向压缩到1.5MB左右
- 评论和点赞功能会出现高并发场景，没有设置消息队列进行削峰
- 关注需要设置上限
- Go-Zero生成的架构文件并不支持在logic模块里面处理文件上传，所以需要将所有的代码放在handler模块处理
- MySQL数据库没有设置支持emoji
- 暂无进行性能测试
- 暂无部署到服务器运行
### 已识别的优化项
- [x] 服务器处理超时问题（解决：API服务端，API客户端，RPC服务端都需延长超时时间，数据库操作创建协程异步进行操作）
- [x] 数据库连接数量超过范围，并且每次请求都需要重新建立连接（解决：将原先每个文件都有的数据库和缓存连接整合到svc模块中，加载时即可创立连接）
- [x] 未登录状态下，评论，关注和粉丝列表的查询如果发生错误会打印网络错误（解决：修改鉴权方式，即使没有携带token也可以进行查询）
### 架构演进的可能
- Redis引入分布式集群
- MySQL引入分布式集群
- 引入Docker进行部署
- 将Token鉴权放入中间件
- 使用对象存储云服务放置视频和图片等静态资源
### 项目过程的反思和总结
整个项目大概写了二十来天，很多概念和知识都是第一次接触，所以经常要边学边写的。再加上队友完全不上心，我也处于一个刚入门的状态，也没办法带他们，只能自己一个人写。前期的架构设计和一些框架的操作问题，就在网上找别人的项目实战文章和视频去学习别人是怎么做的，后面测试debug遇到问题，就截图去大群里面问，或者翻聊天记录看看有没有聊到这个问题，或者去网上搜。
