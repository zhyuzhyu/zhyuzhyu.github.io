---
title: "我的文章标题"
date: 2024-05-20
categories: [技术, 教程]
---

- awk调用shell系统函数，命令必须用双引号包起来
  - 使用awk的system函数：awk '{system("echo hello")}'<<<world
  - 使用awk的print函数：awk '{print "echo hello"|"/bin/bash"}'<<<world
  - 使用awk的getline函数：awk '{"echo hello"|getline _ret;print _ret}'<<<world

- awk调用函数，传递参数，需要转义，示例：
```
awk '{
  cmd="echo "$1""
  system(cmd)
  close(cmd)
}'<<<world
```
- awk调用函数，传递参数，获取返回值，只能使用getline，调用代码：
```
awk '{
  cmd="echo "$1""
  cmd|getline _ret
  close(cmd)
  print _ret
  }' <<< world
```
- awk调用shell自定义函数，只能在另外个文件中定义函数，并export -f出来，然后调用时先source此文件，否则会报函数找不到；
  - 定义test.sh：
  ```
  #!/bin/bash
  myfunc() {
    echo "Hello, $1!"
  }
  export -f myfunc
  ```
  - 调用代码
  ```
  awk '{
  cmd="bash -c \"source test.sh && myfunc " $1 "\""
  cmd|getline _ret
  close(cmd)
  print _ret
  }' <<< world
  ```
