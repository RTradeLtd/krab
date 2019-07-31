# krab

`krab` 是一个“半安全”级别的密钥库，它适配了IPFS秘钥库接口，并可以与许多现有的IPFS扩展实现配合使用。它将密钥存储于Badger上，并且数据存入存储区前都会加密密钥。每次获取密钥时，首先对其进行解密。一份密码加密所有密钥。

## 多语言

[![](https://img.shields.io/badge/Lang-English-blue.svg)](README.md)  [![jaywcjlove/sb](https://jaywcjlove.github.io/sb/lang/chinese.svg)](README-zh.md)

## 用例

有两种使用`krab`的方式，一个是使用`import "github.com/RTradeLtd/krab"`导入，另一个是使用[RTradeLtd/kaas](https://github.com/RTradeLtd/kaas) gRPC API方式。

## 限制

由于badger被用作底层数据存储层，单个badger数据存储区域无法同时拥有多个独立的读写操作权限，只支持可读。写入数据（大部分情况是存储密钥的需求），则必须使用gRPC版本，因为它允许通过多个不同的服务使用单个badger数据存储区。

## 未来的发展

* 定义密钥组，每个组具有单独的加密密码