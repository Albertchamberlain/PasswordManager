package main

// import (
// 	"crypto/rand"
// 	"errors"
// 	"math/big"
// 	"strings"
// 	"time"
// )

// // 默认配置常量
// const (
// 	LengthWeak                int = 6
// 	LengthOK                  int = 12
// 	LengthStrong              int = 24
// 	LengthVeryStrong          int = 36
// 	DefaultLetterSet              = "abcdefghijklmnopqrstuvwxyz"
// 	DefaultNumberSet              = "0123456789"
// 	DefaultSymbolSet              = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"
// 	DefaultLetterAmbiguousSet     = "ijlo"
// 	DefaultSymbolAmbiguousSet     = "<>[](){}:;'/|\\,"
// 	DefaultNumberAmbiguousSet     = "01"
// )

// var (
// 	// 默认配置 全true配置
// 	//0000
// 	DefaultConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             true, //包含特殊符号
// 		IncludeNumbers:             true, //包含小写字母
// 		IncludeLowercaseLetters:    true, //包含小写字母
// 		IncludeUppercaseLetters:    true, //包含大写字母
// 		ExcludeSimilarCharacters:   true, //排除相似字符
// 		ExcludeAmbiguousCharacters: true, //排除连续的字符
// 	}
// 	//0001
// 	ExcludeSymbolsConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             false, //包含特殊符号
// 		IncludeNumbers:             true,  //包含小写字母
// 		IncludeLowercaseLetters:    true,  //包含小写字母
// 		IncludeUppercaseLetters:    true,  //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//0001
// 	ExcludeUppercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             true,  //包含特殊符号
// 		IncludeNumbers:             true,  //包含小写字母
// 		IncludeLowercaseLetters:    true,  //包含小写字母
// 		IncludeUppercaseLetters:    false, //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//0010
// 	ExcludeLowercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             true,  //包含特殊符号
// 		IncludeNumbers:             true,  //包含小写字母
// 		IncludeLowercaseLetters:    false, //包含小写字母
// 		IncludeUppercaseLetters:    true,  //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//0011
// 	ExcludeLowercaseUppercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             true,  //包含特殊符号
// 		IncludeNumbers:             true,  //包含小写字母
// 		IncludeLowercaseLetters:    false, //包含小写字母
// 		IncludeUppercaseLetters:    false, //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//0100
// 	ExcludeNumbersConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             true,  //包含特殊符号
// 		IncludeNumbers:             false, //包含小写字母
// 		IncludeLowercaseLetters:    true,  //包含小写字母
// 		IncludeUppercaseLetters:    true,  //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//0101
// 	ExcludeNumbersUppercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             true,  //包含特殊符号
// 		IncludeNumbers:             false, //包含小写字母
// 		IncludeLowercaseLetters:    true,  //包含小写字母
// 		IncludeUppercaseLetters:    false, //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//0110
// 	ExcludeNumbersLowercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             true,  //包含特殊符号
// 		IncludeNumbers:             false, //包含小写字母
// 		IncludeLowercaseLetters:    false, //包含小写字母
// 		IncludeUppercaseLetters:    true,  //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//0111
// 	ExcludeNumbersLowercaseUppercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             true,  //包含特殊符号
// 		IncludeNumbers:             false, //包含小写字母
// 		IncludeLowercaseLetters:    false, //包含小写字母
// 		IncludeUppercaseLetters:    false, //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//1001
// 	ExcludeSymbolsUppercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             false, //包含特殊符号
// 		IncludeNumbers:             true,  //包含小写字母
// 		IncludeLowercaseLetters:    true,  //包含小写字母
// 		IncludeUppercaseLetters:    false, //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//1010
// 	ExcludeSymbolsLowercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             false, //包含特殊符号
// 		IncludeNumbers:             true,  //包含小写字母
// 		IncludeLowercaseLetters:    false, //包含小写字母
// 		IncludeUppercaseLetters:    true,  //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//1011
// 	ExcludeSymbolsLowercaseUppercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             false, //包含特殊符号
// 		IncludeNumbers:             true,  //包含小写字母
// 		IncludeLowercaseLetters:    false, //包含小写字母
// 		IncludeUppercaseLetters:    false, //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//1100
// 	ExcludeSymbolsNumbersConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             false, //包含特殊符号
// 		IncludeNumbers:             false, //包含小写字母
// 		IncludeLowercaseLetters:    true,  //包含小写字母
// 		IncludeUppercaseLetters:    true,  //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//1101
// 	ExcludeSymbolsNumbersUppercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             false, //包含特殊符号
// 		IncludeNumbers:             false, //包含小写字母
// 		IncludeLowercaseLetters:    true,  //包含小写字母
// 		IncludeUppercaseLetters:    false, //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//1110
// 	ExcludeSymbolsNumbersLowercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             false, //包含特殊符号
// 		IncludeNumbers:             false, //包含小写字母
// 		IncludeLowercaseLetters:    false, //包含小写字母
// 		IncludeUppercaseLetters:    true,  //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	//1111
// 	ExcludeSymbolsNumbersLowercaseUppercaseConfig = Config{
// 		Length:                     LengthStrong,
// 		IncludeSymbols:             false, //包含特殊符号
// 		IncludeNumbers:             false, //包含小写字母
// 		IncludeLowercaseLetters:    false, //包含小写字母
// 		IncludeUppercaseLetters:    false, //包含大写字母
// 		ExcludeSimilarCharacters:   true,  //排除相似字符
// 		ExcludeAmbiguousCharacters: true,  //排除连续的字符
// 	}
// 	ErrConfigIsEmpty = errors.New("config is empty")
// )

// type Generator struct {
// 	*Config
// }

// type Config struct {
// 	Length                     int
// 	CharacterSet               string
// 	IncludeSymbols             bool
// 	IncludeNumbers             bool
// 	IncludeLowercaseLetters    bool
// 	IncludeUppercaseLetters    bool
// 	ExcludeSimilarCharacters   bool
// 	ExcludeAmbiguousCharacters bool
// }

// // New一个密码生成器
// func New(config *Config) (*Generator, error) {
// 	if config == nil {
// 		config = &DefaultConfig
// 	}
// 	if !config.IncludeSymbols &&
// 		!config.IncludeUppercaseLetters &&
// 		!config.IncludeLowercaseLetters &&
// 		!config.IncludeNumbers &&
// 		config.CharacterSet == "" {
// 		return nil, ErrConfigIsEmpty
// 	}

// 	if config.Length == 0 {
// 		config.Length = LengthStrong
// 	}

// 	if config.CharacterSet == "" {
// 		config.CharacterSet = buildCharacterSet(config)
// 	}

// 	return &Generator{Config: config}, nil
// }

// func buildCharacterSet(config *Config) string {
// 	var characterSet string
// 	if config.IncludeLowercaseLetters {
// 		characterSet += DefaultLetterSet
// 		// if config.ExcludeSimilarCharacters {
// 		// 	characterSet = removeCharacters(characterSet, DefaultLetterAmbiguousSet)
// 		// }
// 	}
// 	if config.IncludeUppercaseLetters {
// 		characterSet += strings.ToUpper(DefaultLetterSet)
// 		// if config.ExcludeSimilarCharacters {
// 		// 	characterSet = removeCharacters(characterSet, strings.ToUpper(DefaultLetterAmbiguousSet))
// 		// }
// 	}
// 	if config.IncludeNumbers {
// 		characterSet += DefaultNumberSet
// 		// if config.ExcludeSimilarCharacters {
// 		// 	characterSet = removeCharacters(characterSet, DefaultNumberAmbiguousSet)
// 		// }
// 	}
// 	if config.IncludeSymbols {
// 		characterSet += DefaultSymbolSet
// 		// if config.ExcludeAmbiguousCharacters {
// 		// 	characterSet = removeCharacters(characterSet, DefaultSymbolAmbiguousSet)
// 		// }
// 	}
// 	return characterSet
// }

// func removeCharacters(str, characters string) string {
// 	return strings.Map(func(r rune) rune {
// 		if !strings.ContainsRune(characters, r) {
// 			return r
// 		}
// 		return -1
// 	}, str)
// }

// func NewWithDefault() (*Generator, error) {
// 	return New(&DefaultConfig)
// }

// // 生成一个密码
// func (g Generator) Generate() (*string, error) {
// 	var generated string
// 	characterSet := strings.Split(g.Config.CharacterSet, "")
// 	max := big.NewInt(int64(len(characterSet)))

// 	for i := 0; i < g.Config.Length; i++ {
// 		val, err := rand.Int(rand.Reader, max)
// 		if err != nil {
// 			return nil, err
// 		}
// 		generated += characterSet[val.Int64()]
// 	}
// 	return &generated, nil
// }

// // 一次性生成多个密码
// func (g Generator) GenerateMany(amount int) ([]string, error) {
// 	var generated []string
// 	for i := 0; i < amount; i++ {
// 		str, err := g.Generate()
// 		if err != nil {
// 			return nil, err
// 		}
// 		generated = append(generated, *str)
// 	}
// 	return generated, nil
// }

// // 生成一个指定长度的密码
// func (g Generator) GenerateWithLength(length int) (*string, error) {
// 	var generated string
// 	characterSet := strings.Split(g.Config.CharacterSet, "")
// 	max := big.NewInt(int64(len(characterSet)))
// 	for i := 0; i < length; i++ {
// 		val, err := rand.Int(rand.Reader, max)
// 		if err != nil {
// 			return nil, err
// 		}
// 		generated += characterSet[val.Int64()]
// 	}
// 	return &generated, nil
// }

// // 一次性生成多个指定长度的密码
// func (g Generator) GenerateManyWithLength(amount, length int) ([]string, error) {
// 	var generated []string
// 	for i := 0; i < amount; i++ {
// 		str, err := g.GenerateWithLength(length)
// 		if err != nil {
// 			return nil, err
// 		}
// 		generated = append(generated, *str)
// 	}
// 	return generated, nil
// }

// // 一次性生成7个密码,然后根据星期几来从生成的密码中挑选一个
// func (g Generator) GenerateWithWeek(length int) (string, error) {
// 	amount := 7
// 	week := time.Now().Weekday()
// 	generatPassByWeek, err := g.GenerateManyWithLength(length, amount)
// 	if err != nil {
// 		return "", err
// 	}
// 	return generatPassByWeek[week], nil
// }
