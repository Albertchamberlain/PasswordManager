package main

// func mergeTwoString(a, b string) string {
// 	return a + ":::" + b
// }

// func generatePassword(isInLower, isInUpper, isInNumber, isInSpecial, sliderValue int) string {
// 	//0001
// 	if isInLower == 0 && isInUpper == 0 && isInNumber == 0 && isInSpecial == 1 {
// 		excludeSpecialgenerater, err := New(&ExcludeSymbolsConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeSpecial, err2 := excludeSpecialgenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeSpecial
// 	}
// 	if isInLower == 0 && isInUpper == 0 && isInNumber == 1 && isInSpecial == 0 {
// 		excludeNumbergenerater, err := New(&ExcludeNumbersConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeNumber, err2 := excludeNumbergenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeNumber
// 	}
// 	//0011
// 	if isInLower == 0 && isInUpper == 0 && isInNumber == 1 && isInSpecial == 1 {
// 		excludeNumberAndSpecialgenerater, err := New(&ExcludeSymbolsNumbersConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeNumberAndSpecial, err2 := excludeNumberAndSpecialgenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeNumberAndSpecial
// 	}
// 	//0100
// 	if isInLower == 0 && isInUpper == 1 && isInNumber == 0 && isInSpecial == 0 {
// 		excludeUppergenerater, err := New(&ExcludeUppercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeUpper, err2 := excludeUppergenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeUpper
// 	}
// 	//0101
// 	if isInLower == 0 && isInUpper == 1 && isInNumber == 0 && isInSpecial == 1 {
// 		excludeUpperAndSpecialgenerater, err := New(&ExcludeSymbolsUppercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeUpperAndSpecial, err2 := excludeUpperAndSpecialgenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeUpperAndSpecial
// 	}
// 	//0110
// 	if isInLower == 0 && isInUpper == 1 && isInNumber == 1 && isInSpecial == 0 {
// 		excludeUpperAndNumbergenerater, err := New(&ExcludeNumbersUppercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeUpperAndNumber, err2 := excludeUpperAndNumbergenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeUpperAndNumber
// 	}
// 	//0111
// 	if isInLower == 0 && isInUpper == 1 && isInNumber == 1 && isInSpecial == 1 {
// 		excludeUpperAndNumberAndSpecialgenerater, err := New(&ExcludeSymbolsNumbersUppercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeUpperAndNumberAndSpecial, err2 := excludeUpperAndNumberAndSpecialgenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeUpperAndNumberAndSpecial
// 	}
// 	//1000
// 	if isInLower == 1 && isInUpper == 0 && isInNumber == 0 && isInSpecial == 0 {
// 		excludeLowergenerater, err := New(&ExcludeLowercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeLower, err2 := excludeLowergenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeLower
// 	}
// 	//1001
// 	if isInLower == 1 && isInUpper == 0 && isInNumber == 0 && isInSpecial == 1 {
// 		excludeLowerAndSpecialgenerater, err := New(&ExcludeSymbolsLowercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeLowerAndSpecial, err2 := excludeLowerAndSpecialgenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeLowerAndSpecial
// 	}
// 	//1010
// 	if isInLower == 1 && isInUpper == 0 && isInNumber == 1 && isInSpecial == 0 {
// 		excludeLowerAndNumbergenerater, err := New(&ExcludeNumbersLowercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeLowerAndNumber, err2 := excludeLowerAndNumbergenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeLowerAndNumber
// 	}
// 	//1011
// 	if isInLower == 1 && isInUpper == 0 && isInNumber == 1 && isInSpecial == 1 {
// 		excludeLowerAndNumberAndSpecialgenerater, err := New(&ExcludeSymbolsNumbersLowercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeLowerAndNumberAndSpecial, err2 := excludeLowerAndNumberAndSpecialgenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeLowerAndNumberAndSpecial
// 	}
// 	//1100
// 	if isInLower == 1 && isInUpper == 1 && isInNumber == 0 && isInSpecial == 0 {
// 		excludeLowerAndUppergenerater, err := New(&ExcludeLowercaseUppercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeLowerAndUpper, err2 := excludeLowerAndUppergenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeLowerAndUpper
// 	}
// 	//1101
// 	if isInLower == 1 && isInUpper == 1 && isInNumber == 0 && isInSpecial == 1 {
// 		excludeLowerAndUpperAndSpecialgenerater, err := New(&ExcludeSymbolsLowercaseUppercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeLowerAndUpperAndSpecial, err2 := excludeLowerAndUpperAndSpecialgenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeLowerAndUpperAndSpecial
// 	}
// 	//1110
// 	if isInLower == 1 && isInUpper == 1 && isInNumber == 1 && isInSpecial == 0 {
// 		excludeLowerAndUpperAndNumbergenerater, err := New(&ExcludeNumbersLowercaseUppercaseConfig)
// 		if err != nil {
// 			panic(err)
// 		}
// 		passwordByExcludeLowerAndUpperAndNumber, err2 := excludeLowerAndUpperAndNumbergenerater.GenerateWithWeek(sliderValue)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		return passwordByExcludeLowerAndUpperAndNumber
// 	}
// 	//1111
// 	if isInLower == 1 && isInUpper == 1 && isInNumber == 1 && isInSpecial == 1 {
// 		return "密码规则过于简单"
// 	}
// 	defaultGenerater, err := New(&DefaultConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	passwordByDefaultGenerater, generr := defaultGenerater.GenerateWithWeek(sliderValue)
// 	if generr != nil {
// 		panic(generr)
// 	}
// 	return passwordByDefaultGenerater
// }

// func encipherByRSA(str string) []byte {
// 	//公钥加密
// 	encryptBytes, err := LockWithPublicKey([]byte(str), PUBLICFILE)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return encryptBytes
// }

// func decipherByRSA(str string) string {
// 	//私钥解密
// 	decryptBytes, err := UnlockWithPrivateKey([]byte(str), PRIVATEFILE)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return string(decryptBytes)
// }

// type Service struct {
// }

// func (s Service) InsertKV2DB(website, username string, isInLower, isInUpper, isInNumber, isInSpecial, sliderValue int) string {
// 	orinPass := generatePassword(isInLower, isInUpper, isInNumber, isInSpecial, sliderValue)
// 	encryptPass := encipherByRSA(orinPass)
// 	Key := mergeTwoString(website, username)
// 	flag := SetKeyAndValue(Key, encryptPass)
// 	if flag {
// 		return "插入成功"
// 	} else {
// 		return "插入失败"
// 	}
// }

// func (s Service) GetValueFromDB(website, username string) string {
// 	Key := mergeTwoString(website, username)
// 	encryptPass := GetValueByKey(Key)
// 	decryptPass := decipherByRSA(encryptPass)
// 	return decryptPass
// }
