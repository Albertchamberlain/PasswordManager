package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/atotto/clipboard"
	"github.com/syndtr/goleveldb/leveldb"
	"gopkg.in/toast.v1"
)

var mainwin *ui.Window

func makeBasicControlsPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	vbox.Append(ui.NewHorizontalSeparator(), false)
	dateNow := time.Now()
	vbox.Append(ui.NewLabel("当前日期："+dateNow.Format("2006-01-02 Mon")), false)
	vbox.Append(ui.NewHorizontalSeparator(), false)

	group := ui.NewGroup("信息填写")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	website := ui.NewEntry()

	entryForm.Append("网站地址:", website, false)

	// website.OnChanged(func(*ui.Entry) {
	// 	fmt.Println("网站地址:", website.Text())
	// })

	usernameOrEmail := ui.NewEntry()
	entryForm.Append("用户名或邮箱:", usernameOrEmail, false)

	List1 := ui.NewCombobox()
	List1.Append("是")
	List1.Append("否")
	entryForm.Append("是否含小写字母:", List1, false)
	list1Value := -1
	List1.OnSelected(func(*ui.Combobox) {
		list1Value = List1.Selected() //是0，否1
	})

	List2 := ui.NewCombobox()
	List2.Append("是")
	List2.Append("否")
	entryForm.Append("是否含大写字母:", List2, false)
	list2Value := -1
	List2.OnSelected(func(*ui.Combobox) {
		list2Value = List2.Selected() //是0，否1
	})

	List3 := ui.NewCombobox()
	List3.Append("是")
	List3.Append("否")
	entryForm.Append("是否含数字:", List3, false)
	list3Value := -1
	List3.OnSelected(func(*ui.Combobox) {
		list3Value = List3.Selected() //是0，否1
	})

	List4 := ui.NewCombobox()
	List4.Append("是")
	List4.Append("否")
	entryForm.Append("是否含特殊字符:", List4, false)
	list4Value := -1
	List4.OnSelected(func(*ui.Combobox) {
		list4Value = List4.Selected() //是0，否1
	})

	slider := ui.NewSlider(6, 32)
	slider.OnChanged(func(*ui.Slider) {
		slider.SetValue(slider.Value())
	})

	entryForm.Append("密码总长度:", slider, false)

	submit := ui.NewButton("生成密码")
	submit.OnClicked(func(*ui.Button) {
		//点击了按钮干什么
		slidervalue := slider.Value()
		websiteValue := website.Text()
		usernameOrEmailValue := usernameOrEmail.Text()
		res := InsertKV2DB(websiteValue, usernameOrEmailValue, list1Value, list2Value, list3Value, list4Value, slidervalue)

		passResEntry := ui.NewEntry()
		passResEntry.SetText(res)
		vbox.Append(ui.NewLabel("生成的密码为："), false)
		vbox.Append(passResEntry, false)
	})
	vbox.Append(submit, false)
	return vbox
}

// www.taobao.com:::amos
func makeNumbersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("密码查找")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	vbox.Append(ui.NewLabel("输入要查找的目标网站和用户"), false)
	urlsearchEntry := ui.NewEntry()
	entryForm.Append("请输入目标网站:", urlsearchEntry, false)
	usernamesearchEntry := ui.NewEntry()
	entryForm.Append("请输入账户名或邮箱:", usernamesearchEntry, false)
	vbox.Append(entryForm, false)

	find := ui.NewButton("密码查找")
	find.OnClicked(func(*ui.Button) {
		user := usernamesearchEntry.Text()
		website := urlsearchEntry.Text()
		password := GetValueFromDB(website, user)
		if password == "" {
			password = "未找到该密码!"
			fmt.Println("未找到该密码!")
		}
		passLable := ui.NewLabel("↓↓↓↓↓" + user + "在" + website + "上的密码为↓↓↓↓↓")
		passResEntry := ui.NewEntry()
		passResEntry.SetText(password)
		entryForm.Append("", passLable, false)
		entryForm.Append("", passResEntry, false)
		notification := toast.Notification{
			AppID:   "Password Manager",
			Title:   "Your Password",
			Message: "Your Password is: " + "**********",
			Icon:    "D:\\toast\\gopher.png", // This file must exist (remove this line if it doesn't)
			Actions: []toast.Action{
				{Type: "protocol", Label: "Copy the Password", Arguments: ""},
				{Type: "protocol", Label: "Later", Arguments: ""},
			},
			Duration: toast.Long,
		}
		err := notification.Push()
		if err != nil {
			log.Fatalln(err)
		}
		// 复制内容到剪切板
		clipboard.WriteAll(password)
		fmt.Println("got it!") //拿到了slider的值
	})
	vbox.Append(find, false)
	// 一直滚动的滚动条
	// ip := ui.NewProgressBar()
	// ip.SetValue(-1)
	// vbox.Append(ip, false)

	group = ui.NewGroup("密码弃用")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)
	vbox.Append(ui.NewLabel("输入要弃用的目标"), false)

	entryFormDiscard := ui.NewForm()
	entryFormDiscard.SetPadded(true)
	urlsearchEntryDiscard := ui.NewEntry()
	entryFormDiscard.Append("请输入目标网站:", urlsearchEntryDiscard, false)
	usernamesearchEntryDisacrd := ui.NewEntry()
	entryFormDiscard.Append("请输入账户名或邮箱:", usernamesearchEntryDisacrd, false)
	vbox.Append(entryFormDiscard, false)

	sureButton := ui.NewRadioButtons()
	sureButton.Append("确认")
	var sure int
	sureButton.OnSelected(func(*ui.RadioButtons) {
		//点击了按钮干什么
		sure = sureButton.Selected() //选中输出的是0
		fmt.Println(sure)
	})
	vbox.Append(sureButton, false)

	discardButton := ui.NewButton("弃用密码")
	discardButton.OnClicked(func(*ui.Button) {
		urldiscard := urlsearchEntryDiscard.Text()
		usernamediscard := usernamesearchEntryDisacrd.Text()
		if sure == 0 {
			key := mergeTwoString(urldiscard, usernamediscard)
			isdicard := DelKeyAndValue(key)
			if isdicard {
				vbox.Append(ui.NewLabel("密码弃用成功!"), false)
			} else {
				vbox.Append(ui.NewLabel("你要弃用的密码不存在,可能其已经被弃用!"), false)
			}
		} else {
			vbox.Append(ui.NewLabel("请确认你的操作!"), false)
		}
	})
	vbox.Append(discardButton, false)
	return hbox
}

func setupUI() {
	mainwin = ui.NewWindow("PasswordManger", 500*2, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("密码生成", makeBasicControlsPage())
	tab.SetMargined(0, true)

	tab.Append("密码管理", makeNumbersPage())
	tab.SetMargined(1, true)

	mainwin.Show()
}

var DB *leveldb.DB

// 截取时间戳后8位作为文件
var TimeStamp string
var (
	PRIVATEFILE string
	PUBLICFILE  string
)

func SetDirAndFileName() {
	TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)[:8]
	PRIVATEFILE = "myRSA/privateKey.pem"
	PUBLICFILE = "myRSA/publicKey.pem"
}

func main() {
	SetDirAndFileName() //初始化公钥私钥存放的文件夹
	fmt.Println("TimeStamp:", TimeStamp)
	//var err error
	// bits, err := strconv.Atoi(TimeStamp)
	// if err != nil {
	// 	panic(err)
	// }
	GenerateKeyFile(4096) //生成公钥私钥
	var err2 error
	DB, err2 = leveldb.OpenFile("db", nil)
	if err2 != nil {
		fmt.Println("levelDB打开失败, err:", err2)
		panic(err2)
	}
	defer DB.Close()
	ui.Main(setupUI)
}

// Service.go
// ********************************************************************************
func mergeTwoString(a, b string) string {
	return a + ":::" + b
}

func generatePassword(isInLower, isInUpper, isInNumber, isInSpecial, sliderValue int) string {
	//0001
	if isInLower == 0 && isInUpper == 0 && isInNumber == 0 && isInSpecial == 1 {
		excludeSpecialgenerater, err := New(&ExcludeSymbolsConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeSpecial, err2 := excludeSpecialgenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeSpecial
	}
	if isInLower == 0 && isInUpper == 0 && isInNumber == 1 && isInSpecial == 0 {
		excludeNumbergenerater, err := New(&ExcludeNumbersConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeNumber, err2 := excludeNumbergenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeNumber
	}
	//0011
	if isInLower == 0 && isInUpper == 0 && isInNumber == 1 && isInSpecial == 1 {
		excludeNumberAndSpecialgenerater, err := New(&ExcludeSymbolsNumbersConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeNumberAndSpecial, err2 := excludeNumberAndSpecialgenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeNumberAndSpecial
	}
	//0100
	if isInLower == 0 && isInUpper == 1 && isInNumber == 0 && isInSpecial == 0 {
		excludeUppergenerater, err := New(&ExcludeUppercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeUpper, err2 := excludeUppergenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeUpper
	}
	//0101
	if isInLower == 0 && isInUpper == 1 && isInNumber == 0 && isInSpecial == 1 {
		excludeUpperAndSpecialgenerater, err := New(&ExcludeSymbolsUppercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeUpperAndSpecial, err2 := excludeUpperAndSpecialgenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeUpperAndSpecial
	}
	//0110
	if isInLower == 0 && isInUpper == 1 && isInNumber == 1 && isInSpecial == 0 {
		excludeUpperAndNumbergenerater, err := New(&ExcludeNumbersUppercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeUpperAndNumber, err2 := excludeUpperAndNumbergenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeUpperAndNumber
	}
	//0111
	if isInLower == 0 && isInUpper == 1 && isInNumber == 1 && isInSpecial == 1 {
		excludeUpperAndNumberAndSpecialgenerater, err := New(&ExcludeSymbolsNumbersUppercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeUpperAndNumberAndSpecial, err2 := excludeUpperAndNumberAndSpecialgenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeUpperAndNumberAndSpecial
	}
	//1000
	if isInLower == 1 && isInUpper == 0 && isInNumber == 0 && isInSpecial == 0 {
		excludeLowergenerater, err := New(&ExcludeLowercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeLower, err2 := excludeLowergenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeLower
	}
	//1001
	if isInLower == 1 && isInUpper == 0 && isInNumber == 0 && isInSpecial == 1 {
		excludeLowerAndSpecialgenerater, err := New(&ExcludeSymbolsLowercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeLowerAndSpecial, err2 := excludeLowerAndSpecialgenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeLowerAndSpecial
	}
	//1010
	if isInLower == 1 && isInUpper == 0 && isInNumber == 1 && isInSpecial == 0 {
		excludeLowerAndNumbergenerater, err := New(&ExcludeNumbersLowercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeLowerAndNumber, err2 := excludeLowerAndNumbergenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeLowerAndNumber
	}
	//1011
	if isInLower == 1 && isInUpper == 0 && isInNumber == 1 && isInSpecial == 1 {
		excludeLowerAndNumberAndSpecialgenerater, err := New(&ExcludeSymbolsNumbersLowercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeLowerAndNumberAndSpecial, err2 := excludeLowerAndNumberAndSpecialgenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeLowerAndNumberAndSpecial
	}
	//1100
	if isInLower == 1 && isInUpper == 1 && isInNumber == 0 && isInSpecial == 0 {
		excludeLowerAndUppergenerater, err := New(&ExcludeLowercaseUppercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeLowerAndUpper, err2 := excludeLowerAndUppergenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeLowerAndUpper
	}
	//1101
	if isInLower == 1 && isInUpper == 1 && isInNumber == 0 && isInSpecial == 1 {
		excludeLowerAndUpperAndSpecialgenerater, err := New(&ExcludeSymbolsLowercaseUppercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeLowerAndUpperAndSpecial, err2 := excludeLowerAndUpperAndSpecialgenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeLowerAndUpperAndSpecial
	}
	//1110
	if isInLower == 1 && isInUpper == 1 && isInNumber == 1 && isInSpecial == 0 {
		excludeLowerAndUpperAndNumbergenerater, err := New(&ExcludeNumbersLowercaseUppercaseConfig)
		if err != nil {
			panic(err)
		}
		passwordByExcludeLowerAndUpperAndNumber, err2 := excludeLowerAndUpperAndNumbergenerater.GenerateWithWeek(sliderValue)
		if err2 != nil {
			panic(err2)
		}
		return passwordByExcludeLowerAndUpperAndNumber
	}
	//1111
	if isInLower == 1 && isInUpper == 1 && isInNumber == 1 && isInSpecial == 1 {
		return "密码规则过于简单"
	}
	defaultGenerater, err := New(&DefaultConfig)
	if err != nil {
		panic(err)
	}
	passwordByDefaultGenerater, generr := defaultGenerater.GenerateWithWeek(sliderValue)
	if generr != nil {
		panic(generr)
	}
	return passwordByDefaultGenerater
}

func encipherByRSA(str string) []byte {
	//公钥加密
	encryptBytes, err := LockWithPublicKey([]byte(str), PUBLICFILE)
	if err != nil {
		fmt.Println(err)
	}
	return encryptBytes
}

func decipherByRSA(str string) string {
	//私钥解密
	decryptBytes, err := UnlockWithPrivateKey([]byte(str), PRIVATEFILE)
	if err != nil {
		fmt.Println(err)
	}
	return string(decryptBytes)
}

func InsertKV2DB(website, username string, isInLower, isInUpper, isInNumber, isInSpecial, sliderValue int) string {
	orinPass := generatePassword(isInLower, isInUpper, isInNumber, isInSpecial, sliderValue)
	encryptPass := encipherByRSA(orinPass)
	Key := mergeTwoString(website, username)
	flag := SetKeyAndValue(Key, encryptPass)
	if flag {
		return orinPass
	} else {
		return "密码插入失败!"
	}
}

func GetValueFromDB(website, username string) string {
	Key := mergeTwoString(website, username)
	encryptPass, isGetSuccess := GetValueByKey(Key)
	if !isGetSuccess {
		return ""
	}
	decryptPass := decipherByRSA(encryptPass)
	return decryptPass
}

//********************************************************************************

//*********************************密码生成***************************************

// 默认配置常量
const (
	LengthWeak                int = 6
	LengthOK                  int = 12
	LengthStrong              int = 24
	LengthVeryStrong          int = 36
	DefaultLetterSet              = "abcdefghijklmnopqrstuvwxyz"
	DefaultNumberSet              = "0123456789"
	DefaultSymbolSet              = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"
	DefaultLetterAmbiguousSet     = "ijlo"
	DefaultSymbolAmbiguousSet     = "<>[](){}:;'/|\\,"
	DefaultNumberAmbiguousSet     = "01"
)

var (
	// 默认配置 全true配置
	//0000
	DefaultConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true, //包含特殊符号
		IncludeNumbers:             true, //包含小写字母
		IncludeLowercaseLetters:    true, //包含小写字母
		IncludeUppercaseLetters:    true, //包含大写字母
		ExcludeSimilarCharacters:   true, //排除相似字符
		ExcludeAmbiguousCharacters: true, //排除连续的字符
	}
	//0001
	ExcludeSymbolsConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             false, //包含特殊符号
		IncludeNumbers:             true,  //包含小写字母
		IncludeLowercaseLetters:    true,  //包含小写字母
		IncludeUppercaseLetters:    true,  //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//0001
	ExcludeUppercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true,  //包含特殊符号
		IncludeNumbers:             true,  //包含小写字母
		IncludeLowercaseLetters:    true,  //包含小写字母
		IncludeUppercaseLetters:    false, //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//0010
	ExcludeLowercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true,  //包含特殊符号
		IncludeNumbers:             true,  //包含小写字母
		IncludeLowercaseLetters:    false, //包含小写字母
		IncludeUppercaseLetters:    true,  //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//0011
	ExcludeLowercaseUppercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true,  //包含特殊符号
		IncludeNumbers:             true,  //包含小写字母
		IncludeLowercaseLetters:    false, //包含小写字母
		IncludeUppercaseLetters:    false, //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//0100
	ExcludeNumbersConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true,  //包含特殊符号
		IncludeNumbers:             false, //包含小写字母
		IncludeLowercaseLetters:    true,  //包含小写字母
		IncludeUppercaseLetters:    true,  //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//0101
	ExcludeNumbersUppercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true,  //包含特殊符号
		IncludeNumbers:             false, //包含小写字母
		IncludeLowercaseLetters:    true,  //包含小写字母
		IncludeUppercaseLetters:    false, //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//0110
	ExcludeNumbersLowercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true,  //包含特殊符号
		IncludeNumbers:             false, //包含小写字母
		IncludeLowercaseLetters:    false, //包含小写字母
		IncludeUppercaseLetters:    true,  //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//0111
	ExcludeNumbersLowercaseUppercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true,  //包含特殊符号
		IncludeNumbers:             false, //包含小写字母
		IncludeLowercaseLetters:    false, //包含小写字母
		IncludeUppercaseLetters:    false, //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//1001
	ExcludeSymbolsUppercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             false, //包含特殊符号
		IncludeNumbers:             true,  //包含小写字母
		IncludeLowercaseLetters:    true,  //包含小写字母
		IncludeUppercaseLetters:    false, //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//1010
	ExcludeSymbolsLowercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             false, //包含特殊符号
		IncludeNumbers:             true,  //包含小写字母
		IncludeLowercaseLetters:    false, //包含小写字母
		IncludeUppercaseLetters:    true,  //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//1011
	ExcludeSymbolsLowercaseUppercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             false, //包含特殊符号
		IncludeNumbers:             true,  //包含小写字母
		IncludeLowercaseLetters:    false, //包含小写字母
		IncludeUppercaseLetters:    false, //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//1100
	ExcludeSymbolsNumbersConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             false, //包含特殊符号
		IncludeNumbers:             false, //包含小写字母
		IncludeLowercaseLetters:    true,  //包含小写字母
		IncludeUppercaseLetters:    true,  //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//1101
	ExcludeSymbolsNumbersUppercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             false, //包含特殊符号
		IncludeNumbers:             false, //包含小写字母
		IncludeLowercaseLetters:    true,  //包含小写字母
		IncludeUppercaseLetters:    false, //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//1110
	ExcludeSymbolsNumbersLowercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             false, //包含特殊符号
		IncludeNumbers:             false, //包含小写字母
		IncludeLowercaseLetters:    false, //包含小写字母
		IncludeUppercaseLetters:    true,  //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	//1111
	ExcludeSymbolsNumbersLowercaseUppercaseConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             false, //包含特殊符号
		IncludeNumbers:             false, //包含小写字母
		IncludeLowercaseLetters:    false, //包含小写字母
		IncludeUppercaseLetters:    false, //包含大写字母
		ExcludeSimilarCharacters:   true,  //排除相似字符
		ExcludeAmbiguousCharacters: true,  //排除连续的字符
	}
	ErrConfigIsEmpty = errors.New("config is empty")
)

type Generator struct {
	*Config
}

type Config struct {
	Length                     int
	CharacterSet               string
	IncludeSymbols             bool
	IncludeNumbers             bool
	IncludeLowercaseLetters    bool
	IncludeUppercaseLetters    bool
	ExcludeSimilarCharacters   bool
	ExcludeAmbiguousCharacters bool
}

// New一个密码生成器
func New(config *Config) (*Generator, error) {
	if config == nil {
		config = &DefaultConfig
	}
	if !config.IncludeSymbols &&
		!config.IncludeUppercaseLetters &&
		!config.IncludeLowercaseLetters &&
		!config.IncludeNumbers &&
		config.CharacterSet == "" {
		return nil, ErrConfigIsEmpty
	}

	if config.Length == 0 {
		config.Length = LengthStrong
	}

	if config.CharacterSet == "" {
		config.CharacterSet = buildCharacterSet(config)
	}

	return &Generator{Config: config}, nil
}

func buildCharacterSet(config *Config) string {
	var characterSet string
	if config.IncludeLowercaseLetters {
		characterSet += DefaultLetterSet
	}
	if config.IncludeUppercaseLetters {
		characterSet += strings.ToUpper(DefaultLetterSet)
	}
	if config.IncludeNumbers {
		characterSet += DefaultNumberSet
	}
	if config.IncludeSymbols {
		characterSet += DefaultSymbolSet
		// if config.ExcludeAmbiguousCharacters {
		// 	characterSet = removeCharacters(characterSet, DefaultSymbolAmbiguousSet)
		// }
	}
	return characterSet
}

// func removeCharacters(str, characters string) string {
// 	return strings.Map(func(r rune) rune {
// 		if !strings.ContainsRune(characters, r) {
// 			return r
// 		}
// 		return -1
// 	}, str)
// }

func NewWithDefault() (*Generator, error) {
	return New(&DefaultConfig)
}

// 生成一个密码
func (g Generator) Generate() (*string, error) {
	var generated string
	characterSet := strings.Split(g.Config.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))

	for i := 0; i < g.Config.Length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}

// 一次性生成多个密码
func (g Generator) GenerateMany(amount int) ([]string, error) {
	var generated []string
	for i := 0; i < amount; i++ {
		str, err := g.Generate()
		if err != nil {
			return nil, err
		}
		generated = append(generated, *str)
	}
	return generated, nil
}

// 生成一个指定长度的密码
func (g Generator) GenerateWithLength(length int) (*string, error) {
	var generated string
	characterSet := strings.Split(g.Config.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))
	for i := 0; i < length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}

// 一次性生成多个指定长度的密码
func (g Generator) GenerateManyWithLength(amount, length int) ([]string, error) {
	var generated []string
	for i := 0; i < amount; i++ {
		str, err := g.GenerateWithLength(length)
		if err != nil {
			return nil, err
		}
		generated = append(generated, *str)
	}
	return generated, nil
}

// 一次性生成7个密码,然后根据星期几来从生成的密码中挑选一个
func (g Generator) GenerateWithWeek(length int) (string, error) {
	amount := 7
	week := time.Now().Weekday()
	generatPassByWeek, err := g.GenerateManyWithLength(amount, length)
	if err != nil {
		return "", err
	}
	return generatPassByWeek[week], nil
}

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ******************************************RSAFile********************************************
func GenerateKeyFile(bits int) error {
	//判断公私钥文件是否存在
	err := os.Mkdir("./myRSA", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}
	existPublic, _ := PathExists("public.pem")
	existPrivate, _ := PathExists("private.pem")
	if existPublic && existPrivate {
		return errors.New("公私钥文件已存在")
	}
	//生成私钥
	//使用rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	//通过x509标准将得到的ras私钥序列化为ASN.1的DER编码字符串
	PKCS1PrivateBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	//将私钥字符串设置到pem格式块中
	privateBlock := pem.Block{
		Type:  "RSA Private Key",
		Bytes: PKCS1PrivateBytes,
	}
	//通过pem将设置好的数据进行编码，并写入磁盘文件
	privateFile, err := os.Create(PRIVATEFILE)
	fmt.Println("重复创建私钥文件")
	if err != nil {
		return err
	}
	defer privateFile.Close()
	err = pem.Encode(privateFile, &privateBlock)
	if err != nil {
		return err
	}

	//生成公钥
	//从得到的私钥对象中将公钥信息取出
	publicKey := privateKey.PublicKey

	//通过x509标准将得到的ras公钥序列化为ASN.1的DER编码字符串
	PKCS1PublicBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	//将公钥字符串设置到pem格式块中
	publicBlock := pem.Block{
		Type:  "RSA Public Key",
		Bytes: PKCS1PublicBytes,
	}
	//通过pem将设置好的数据进行编码，并写入磁盘文件
	publicFile, err := os.Create(PUBLICFILE)
	fmt.Println("重复创建公钥文件")
	if err != nil {
		return err
	}
	defer publicFile.Close()
	err = pem.Encode(publicFile, &publicBlock)
	if err != nil {
		return err
	}

	return nil
}

// 公钥加密
func LockWithPublicKey(src []byte, pubKeyFile string) ([]byte, error) {
	var err error
	//将公钥文件中的公钥读出，得到使用pem编码的字符串
	file, err := os.Open(pubKeyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	//将得到的字符串解码
	block, _ := pem.Decode(buffer)

	//使用x509将编码之后的公钥解析出来
	pubInner, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey := pubInner.(*rsa.PublicKey)

	//使用得到的公钥通过rsa进行数据加密
	encryptBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, src)
	if err != nil {
		log.Fatal("公钥加密失败")
		return nil, err
	}

	return encryptBytes, nil
}

// 私钥解密
func UnlockWithPrivateKey(src []byte, privateKeyFile string) ([]byte, error) {
	//将私钥文件中的私钥读出，得到使用pem编码的字符串
	file, err := os.Open(privateKeyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := fileInfo.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	//将得到的字符串解码
	block, _ := pem.Decode(buffer)

	//使用x509将编码之后的私钥解析出来
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	//使用得到的私钥通过rsa进行数据解密
	decryptBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	if err != nil {
		log.Fatal("私钥解密失败", err)
		panic(err)
	}
	return decryptBytes, nil
}

//********************************************************************************
//leveldb的key和value都是byte数组类型

// 根据key获取value
func GetValueByKey(key string) (string, bool) {
	data, err := DB.Get([]byte(key), nil)
	if err == leveldb.ErrNotFound {
		fmt.Println("没有找到该密码")
		return "", false
	}
	return string(data), true
}

// 设置key和value
func SetKeyAndValue(key string, value []byte) bool {
	err := DB.Put([]byte(key), value, nil)
	return err == nil
}

// 根据key删除key value
func DelKeyAndValue(key string) bool {
	_, err := DB.Get([]byte(key), nil)
	if err == leveldb.ErrNotFound {
		return false
	}
	err2 := DB.Delete([]byte(key), nil)
	return err2 == nil
}
