package main

import (
	"fmt"
	"math/rand"
	"time" //시간 관련 패키지(seed 생성용)
)

func main() {
	//seed 설정(난수 사용 시 seed설정이 필요하다.)
	seed := time.Now().Unix() //unix 기준 현재 시간을 얻는 명령어
	rand.Seed(seed)

	//number:=rand.Intn(6)  //0부터 ()안에 숫자 사이에 랜덤 숫자가 나온다.
	// dice := rand.Intn(6)
	// fmt.Println(dice)
	for i := 1; i < 6; i++ { //6번 난수를 출력
		dice := rand.Intn(6) + 1
		fmt.Println(dice)
	}
}
