# ezname
created by fristzzz (wxy in HITSZ)

## 安装方法
请确保电脑上已经安装go语言环境（powershell或cmd输入go version查看)
### 手动安装
```
git clone https://github.com/fristzzz/ezname.git
cd ezname
go env -w GO111MODULE=auto
go build .
```
即可通过
`$ ./ezname.exe `
来使用


## 如何使用

添加专业、班级中的同学信息
需要姓名、学号

### 添加姓名学号
可以使用
```
./ezname.exe add
```
来添加同学姓名、学号

### 自动文件重命名
```
$ ./ezname.exe <作业所在目录> <命名格式>
```
然后根据提示输入专业、班级即可

#### 命名格式
ezname将关键词“学号”替换为学号，“姓名”替换为姓名，“班级”替换为专业+班级

### 示例
```
$ ./ezname.exe C:\frist\作业\ 电路1-班级-姓名-学号
```

### 注意
需要原文件名包含学号，否则ezname不起作用
