# wxlogin

wxlogin 用于通过微信扫码获取 cookie 信息来登录第三方应用. 该项目辅助于 [sign](https://github.com/jdxj/sign) (还未使用)

## 使用方法

1. 克隆

```shell script
$ git clone https://github.com/jdxj/wxlogin.git
```

2. 编译

```shell script
$ cd wxlogin
$ make
```

3. 运行

```shell script
$ ./wxlogin.out
```

4. 使用微信扫码并同意登录. 二维码在当前目录名为 `qrcode.jpeg`
5. 在终端复制出 cookie

## TODO

- 重构, 使其通用
