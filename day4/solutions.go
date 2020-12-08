package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var requiredfields = [7]ppField{{"byr"}, {"iyr"}, {"eyr"}, {"hgt"}, {"hcl"}, {"ecl"}, {"pid"}}
var optionalfields = [1]ppField{{"cid"}}

//byr (Birth Year) - four digits; at least 1920 and at most 2002.
//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
//hgt (Height) - a number followed by either cm or in:
//If cm, the number must be at least 150 and at most 193.
//If in, the number must be at least 59 and at most 76.
//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
//pid (Passport ID) - a nine-digit number, including leading zeroes.
//cid (Country ID) - ignored, missing or not.

// PassportData the data for the program
type PassportData struct {
	datamap map[string]string
}

// Passport a passport
type Passport struct {
	fields map[ppField]ppValue
}

func makePassport() *Passport {
	passport := new(Passport)
	passport.fields = map[ppField]ppValue{}
	return passport
}

type ppValue struct {
	field ppField
	value string
}

type ppField struct {
	key string
}

// PassportMachine validates data
type PassportMachine interface {
	isValid() bool
	readPassport(data []string)
}

var hgtregex = regexp.MustCompile(`^(\d+)(cm|in)$`)

//a # followed by exactly six characters 0-9 or a-f
var hclregex = regexp.MustCompile(`^#[a-f0-9]{6}$`)

var pidregex = regexp.MustCompile(`^([0-9]{9})$`)

var valideyes = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func (f ppValue) isValid() bool {
	intval, _ := strconv.Atoi(f.value)
	strval := f.value
	switch f.field.key {
	case ("byr"):
		if intval < 1920 || intval > 2002 {
			return false
		}
	case ("iyr"):
		if intval < 2010 || intval > 2020 {
			return false
		}
	case ("eyr"):
		if intval < 2020 || intval > 2030 {
			return false
		}
	case ("hgt"):
		d := hgtregex.FindStringSubmatch(strval)
		if d == nil {
			return false
		}
		num, _ := strconv.Atoi(d[1])
		t := d[2]
		if t == "cm" {
			if num < 150 || num > 193 {
				return false
			}
		} else if t == "in" {
			if num < 59 || num > 76 {
				return false
			}
		} else {
			return false
		}
	case ("hcl"):
		return hclregex.MatchString(strval)
	case ("ecl"):
		for _, x := range valideyes {
			if strval == x {
				return true
			}
		}
		return false
	case ("pid"):
		x := pidregex.FindStringSubmatch(strval)
		if x == nil {
			return false
		}
		return true
	case ("cid"):
		return true
	default:
		fmt.Println("Unknown field: ", f.field.key)
	}
	return true
}

func makeIntoPassportData(data []string) PassportData {
	passport := PassportData{}
	passport.datamap = make(map[string]string)
	for _, d := range data {
		ts := strings.Split(d, " ")
		for _, t := range ts {
			kv := strings.Split(t, ":")
			passport.datamap[kv[0]] = kv[1]
		}
	}
	return passport
}

func hasData(ppt PassportData, key string) bool {
	_, prs := ppt.datamap[key]
	return prs
}

func readPassports(passports []PassportData, currentPassport []string, data []string) []PassportData {
	if len(data) == 0 {
		return passports
	}
	nextLine := data[0]
	if nextLine == "" {
		passports = append(passports, makeIntoPassportData(currentPassport))
		currentPassport = make([]string, 0)
	} else {
		currentPassport = append(currentPassport, nextLine)
	}
	return readPassports(passports, currentPassport, data[1:])
}

func (d Passport) readPassport(line string) Passport {
	ts := strings.Split(line, " ")
	for _, t := range ts {
		kv := strings.Split(t, ":")
		ppf := ppField{kv[0]}
		value := ppValue{ppf, kv[1]}
		d.fields[ppf] = value
	}
	return d
}

func scanPassports(passports []Passport, currentPassport Passport, data []string) []Passport {
	if len(data) == 0 {
		return passports
	}
	nextLine := data[0]
	if nextLine == "" {
		passports = append(passports, currentPassport)
		currentPassport = *makePassport()
	} else {
		currentPassport = currentPassport.readPassport(nextLine)
	}
	return scanPassports(passports, currentPassport, data[1:])
}

// Solution2 main
func Solution2(data []string) int {
	passports := make([]Passport, 0)
	currentPassport := makePassport()
	passports = scanPassports(passports, *currentPassport, data)
	fmt.Println(len(passports))
	validcount := 0
	for _, passport := range(passports) {
		valid := true
		for _, f := range requiredfields {
			v, pres := passport.fields[f]
			if pres != true || v.isValid() != true {
				fmt.Printf("failed %v %v %v\n", f, v, pres)
				valid = false
				break
			}
		}
		if valid {
			validcount++
		}
	}

	return validcount
}

// Solution 4.1
func Solution(data []string) int {
	fmt.Printf("%v rows of data\n", len(data))
	passports := make([]PassportData, 0)
	currentPassport := make([]string, 0)
	passports = readPassports(passports, currentPassport, data)
	fmt.Println(passports)

	invalidppts := 0

	//for _, ppt := range passports {
//		for _, k := range requiredfields {
//			if hasData(ppt, k) == false {
//				fmt.Printf("missing field %v -> %v\n", k, ppt)
//				invalidppts++
//				break
//			}
//		}
//	}
	fmt.Printf("%d invalid %d total = %d valid\n", invalidppts, len(passports), len(passports)-invalidppts)
	return len(passports) - invalidppts

}
