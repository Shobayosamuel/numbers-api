package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type NumberProperties struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

func classifyNumber(number string) string {
	num, err := strconv.Atoi(number)
	if err != nil {
		return fmt.Sprintf(`{"error": "invalid number"}`)
	}

	properties := NumberProperties{
		Number:     num,
		IsPrime:    isPrime(num),
		IsPerfect:  isPerfect(num),
		Properties: getProperties(num),
		DigitSum:   digitSum(num),
		FunFact:    getFunFact(num),
	}

	response, err := json.Marshal(properties)
	if err != nil {
		return fmt.Sprintf(`{"error": "could not marshal response"}`)
	}

	return string(response)
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isPerfect(num int) bool {
	sum := 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func getProperties(num int) []string {
	var properties []string
	if num%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}
	if isArmstrong(num) {
		properties = append(properties, "armstrong")
	}
	return properties
}

func isArmstrong(num int) bool {
	sum := 0
	temp := num
	n := len(strconv.Itoa(num))
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(n)))
		temp /= 10
	}
	return sum == num
}

func digitSum(num int) int {
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func getFunFact(num int) string {
	resp, err := http.Get(fmt.Sprintf("http://numbersapi.com/%d", num))
	if err != nil {
		return "could not retrieve fun fact"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "could not read fun fact"
	}
	return strings.TrimSpace(string(body))
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	r.GET("/api/classify-number", func(c *gin.Context) {
		fmt.Println("Received number:", c.Query("number"))
		number := c.Query("number")
		if number == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "number is required",
			})
			return
		}
		num, err := strconv.Atoi(number)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid number",
			})
			return
		}

		properties := NumberProperties{
			Number:     num,
			IsPrime:    isPrime(num),
			IsPerfect:  isPerfect(num),
			Properties: getProperties(num),
			DigitSum:   digitSum(num),
			FunFact:    getFunFact(num),
		}

		c.JSON(http.StatusOK, properties)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}