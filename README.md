# douyin
字节跳动第三届青训营，极简抖音后端项目

# 数据库设计E-R图

![E-R图](README.assets/E-R图.png)

- 用户（<u>用户id</u>，用户名，用户密码）
- 视频（<u>视频id</u>，视频文件路径，视频图片路径，修改时间，视频状态<!--审核中/未过审/已提交-->，视频分类）
- 评论（<u>评论id</u>，评论内容，评论时间，回复评论id）
- 观看（<u>用户id，视频id</u>，是否点赞）
- 评论点赞（<u>用户id，评论id</u>，是否点赞）
- 播放量（<u>视频id</u>，播放量）
- 关注/粉丝列表（<u>所关注用户id，粉丝id</u>）

![image-20220506151928800](README.assets/image-20220506151928800.png)