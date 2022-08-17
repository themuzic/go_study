package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type emp struct {
	name   string
	gender string
	age    int
	salary int
}

func main() {
	employees := []emp{}
	fileName := "table.csv"
	// 파일오픈
	file, err := os.Open(fileName)
	checkErr(err)
	// csv reader 생성
	fileReader := csv.NewReader(bufio.NewReader(file))
	// 내용 모두 읽기
	rows, rErr := fileReader.ReadAll()
	checkErr(rErr)

	// 첫번째 row는 컬럼명이기 때문에 두번째 row부터 emp에 담기
	for i := 1; i < len(rows); i++ {
		var temp emp
		temp.name = strings.TrimSpace(rows[i][0])
		temp.gender = strings.TrimSpace(rows[i][1])
		temp.age, _ = strconv.Atoi(rows[i][2])
		r := strings.NewReplacer("$", "", ",", "", ".", "")
		temp.salary, _ = strconv.Atoi(r.Replace(rows[i][3]))
		employees = append(employees, temp)
	}

	var mostSalaryEmp string
	mostSalary := 0
	totalSalary3040 := 0
	count := 0
	for _, e := range employees {
		// 1번답 구하기
		if mostSalary < e.salary {
			mostSalary = e.salary
			mostSalaryEmp = e.name
		}
		// 2번답 구하기
		if 30 <= e.age && e.age <= 49 {
			count++
			totalSalary3040 += e.salary
		}
	}
	// 2번답 가공
	totalSalary3040 /= count
	var arr [3]string
	arr[2] = calcZero(totalSalary3040, 100)
	arr[1] = calcZero(totalSalary3040/100, 1000)
	arr[0] = strconv.Itoa(totalSalary3040 / 100000)
	// 결과 출력
	fmt.Println("1. 우리 회사에서 연봉이 가장 높은사람의 이름을 출력해주세요. :", mostSalaryEmp)
	fmt.Println("2. 우리 회사 30,40대(age between 30 and 49) 직원의 평균 연봉을 출력해주세요. : " +
		"$" + arr[0] + "," + arr[1] + "." + arr[2])
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// 2번답 가공 중 "0"을 붙여서 자릿수 맞춰주기 위한 function
func calcZero(num1, num2 int) string {
	var length int
	var returnStr string
	if num2 == 100 {
		length = 2
	} else if num2 == 1000 {
		length = 3
	}

	temp1 := num1 % num2
	temp2 := len(strconv.Itoa(temp1))

	if temp2 != length {
		for i := 0; i < length-temp2; i++ {
			returnStr += "0"
		}
		returnStr += strconv.Itoa(temp1)
	} else {
		returnStr = strconv.Itoa(temp1)
	}
	return returnStr
}
