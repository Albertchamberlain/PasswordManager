# PasswordManager
基于图形用户界面，使用纯Go实现的密码本地管理系统。包含密码生成、密码查找、密码管理基本功能。

感谢以下开源库
1. "github.com/andlabs/ui"
2. "github.com/syndtr/goleveldb/leveldb"
3. "gopkg.in/toast.v1"
4. "github.com/atotto/clipboard"

![image](https://user-images.githubusercontent.com/47343901/205073399-66f68628-4cf2-42a6-8ad7-c05f84145011.png)


![image](https://user-images.githubusercontent.com/47343901/205073447-525120dd-6f55-40be-9c7f-523a9c8c1312.png)


![9de1b2266fa16787115ff36786eeb87](https://user-images.githubusercontent.com/47343901/205430237-35079219-0623-4362-aa5c-0acba2a5adde.png)



## 代码执行流程

1. 调用crypto/rsa 生成公私钥
   
2. 调用crypto/x509 将公私钥序列化    
                                                           
3. 通过encoding/pem 将公私钥编码为PEM格式
   
4. 接收到用户的密码规则
   
5. 后台根据用户的规则选取对应的Config，然后New generator生成七个密码，根据当前日期挑选一个出来
   
6. 使用x509.ParsePKIXPublicKey 解析公钥
   
7. 调用rsa.EncryptPKCS1v15方法使用解析出的公钥加密密码
   
8. 将生成的密码回写到用户界面，将加密后的密码存储至Level DB


`Step  1，2，3     如果检测到已经存在公私钥就不再生成`



- [x] 用户自定义规则生成密码
- [x] 查找密码
- [x] Win弹窗通知密码生成
- [x] 密码写入剪切板，可直接ctrl-v粘贴
- [ ] soon。。。。
