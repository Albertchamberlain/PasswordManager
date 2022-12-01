package main

// import (
// 	"crypto/rand"
// 	"crypto/rsa"
// 	"crypto/x509"
// 	"encoding/pem"
// 	"log"
// 	"os"
// )

// // var (
// // 	PRIVATEFILE string
// // 	PUBLICFILE  string
// // )

// // // 截取时间戳后8位作为文件
// // var TimeStamp string

// // func SetDirName() {
// // 	TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)[:8]
// // 	PRIVATEFILE = "./" + TimeStamp + "/privateKey.pem"
// // 	PUBLICFILE = "./" + TimeStamp + "/publicKey.pem"
// // }

// // 生成公钥和私钥文件

// func GenerateKeyFile(bits int) error {
// 	//生成私钥
// 	//使用rsa中的GenerateKey方法生成私钥
// 	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
// 	if err != nil {
// 		return err
// 	}

// 	//通过x509标准将得到的ras私钥序列化为ASN.1的DER编码字符串
// 	PKCS1PrivateBytes := x509.MarshalPKCS1PrivateKey(privateKey)

// 	//将私钥字符串设置到pem格式块中
// 	privateBlock := pem.Block{
// 		Type:  "RSA Private Key",
// 		Bytes: PKCS1PrivateBytes,
// 	}

// 	//通过pem将设置好的数据进行编码，并写入磁盘文件
// 	privateFile, err := os.Create(PRIVATEFILE)
// 	if err != nil {
// 		return err
// 	}
// 	defer privateFile.Close()
// 	err = pem.Encode(privateFile, &privateBlock)
// 	if err != nil {
// 		return err
// 	}

// 	//生成公钥
// 	//从得到的私钥对象中将公钥信息取出
// 	publicKey := privateKey.PublicKey

// 	//通过x509标准将得到的ras公钥序列化为ASN.1的DER编码字符串
// 	PKCS1PublicBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
// 	if err != nil {
// 		return err
// 	}

// 	//将公钥字符串设置到pem格式块中
// 	publicBlock := pem.Block{
// 		Type:  "RSA Public Key",
// 		Bytes: PKCS1PublicBytes,
// 	}

// 	//通过pem将设置好的数据进行编码，并写入磁盘文件
// 	publicFile, err := os.Create(PUBLICFILE)
// 	if err != nil {
// 		return err
// 	}
// 	defer publicFile.Close()
// 	err = pem.Encode(publicFile, &publicBlock)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // 公钥加密
// func LockWithPublicKey(src []byte, pubKeyFile string) ([]byte, error) {
// 	var err error
// 	//将公钥文件中的公钥读出，得到使用pem编码的字符串
// 	file, err := os.Open(pubKeyFile)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		return nil, err
// 	}

// 	buffer := make([]byte, fileInfo.Size())
// 	_, err = file.Read(buffer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//将得到的字符串解码
// 	block, _ := pem.Decode(buffer)

// 	//使用x509将编码之后的公钥解析出来
// 	pubInner, err := x509.ParsePKIXPublicKey(block.Bytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	publicKey := pubInner.(*rsa.PublicKey)

// 	//使用得到的公钥通过rsa进行数据加密
// 	encryptBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, src)
// 	if err != nil {
// 		log.Fatal("公钥加密失败")
// 		return nil, err
// 	}

// 	return encryptBytes, nil
// }

// // 私钥解密
// func UnlockWithPrivateKey(src []byte, privateKeyFile string) ([]byte, error) {
// 	//将私钥文件中的私钥读出，得到使用pem编码的字符串
// 	file, err := os.Open(privateKeyFile)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		return nil, err
// 	}
// 	size := fileInfo.Size()
// 	buffer := make([]byte, size)
// 	_, err = file.Read(buffer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//将得到的字符串解码
// 	block, _ := pem.Decode(buffer)

// 	//使用x509将编码之后的私钥解析出来
// 	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//使用得到的私钥通过rsa进行数据解密
// 	decryptBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
// 	if err != nil {
// 		log.Fatal("私钥解密失败")
// 		return nil, err
// 	}
// 	return decryptBytes, nil
// }

// // 测试RSA非对称加解密信息，公钥加密，私钥解密
// // func TestRSA() {
// // 	//生成钥匙对
// // 	err := myRSA.GenerateKey(4096)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}

// // 	srcInfo := "GO 密码学 —— RSA非对称加解密测试"

// // 	//公钥加密
// // 	encryptBytes, err := myRSA.LockWithPublicKey([]byte(srcInfo), myRSA.PUBLICFILE)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}

// // 	//私钥解密
// // 	decryptBytes, err := myRSA.UnlockWithPrivateKey(encryptBytes, myRSA.PRIVATEFILE)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}

// // 	fmt.Println("测试非对称加密结果：")
// // 	fmt.Println("元数据：", srcInfo)
// // 	encryptHex := hex.EncodeToString(encryptBytes)
// // 	fmt.Println("公钥加密数据：", encryptHex)
// // 	fmt.Println("私钥解密数据：", string(decryptBytes))

// // }

// //首先会在规定的目录生成钥匙对文件：privateKey.pem、publicKey.pem

// // OUTPUT:
// // 测试非对称加密结果：
// // 元数据： GO 密码学 —— RSA非对称加解密测试
// // 公钥加密数据： 582a4c0a2acdfb61cf2ee28bf6e85696b11415c5d84b515548268d15789752da50cc82dade81062bdbaccfeee227120eaf0cf4fcc46d174ccd1187a0a1318c987cc5f722bc8abeb9f425d3b890d6842dcda6731e9527528e9bdbbbe84220d972fb07049c1b7c5615731d5b61b148f42800d7c30ef5812da80d7b54bdea1ef93388d4eab33ccccb518da205672e3e9d9a2196b41cb000fc53c21e6a75f34ba2d5d941aac4f6cec9cdfd2ed443fe3207f32d636b4973ed94cc9d12362c55e399eea8163f253210ac8ff05b456046e643bdc26ca50648c2337b56399b5f4296903ca58fb0149fbbfb4ecf9c88d7d8b0beb2e0ea60fa2d3df2b7ffd0987259e720627f7f51070412a2d1e4da75c3dc7de2b9010b3f45e7b40c4f681621fd5aa686081e91fc00d578c0b7c8c41738c3d9b7655abc9963fd2d72f4006d0ad0af4e6e832868296fbe36098c8f226c6bc4557250039e6d631234bf7dabf7df9352d0b42e04cce8dc0d24961e927b996b99d87cf786cc48e865e4c093fc65f5fc047fea2141e5fc3648ae30ef7fc51837e4494954d7535512109ed7ed59e3fbf4e39b1e38a0bef1a771bcba656e77e1d6e522bf813cbf11715e1e781dd0e99f1c8ed5dce856ef1468e3a519c191c476ab778147a6ea9a6c311a9bf17beca2d80712180564d0986c6d02022e3640b8128f1cc4ea473aec79f5c7bd865b4de5c9e33cd9909d
// // 私钥解密数据： GO 密码学 —— RSA非对称加解密测试
