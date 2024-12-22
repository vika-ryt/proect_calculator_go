package calculation

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)


func tokeng(expres string) []string {
    var tokens []string
    var currentToken strings.Builder

    for _, char := range expres {
        switch char {
        case ' ':
            continue
        case '+', '-', '*', '/', '(', ')':
            if currentToken.Len() > 0 {
                tokens = append(tokens, currentToken.String())
                currentToken.Reset()
            }
            tokens = append(tokens, string(char))
        default:
            currentToken.WriteRune(char)
        }
    }

    if currentToken.Len() > 0 {
        tokens = append(tokens, currentToken.String())
    }

    return tokens
}


func infpf(tokens []string) ([]string, error) {
    var output []string
    var operators []string

    for _, token := range tokens {
        if isNumber(token) {
            output = append(output, token)
        } else if token == "(" {
            operators = append(operators, token)
        } else if token == ")" {
            for len(operators) > 0 && operators[len(operators)-1] != "(" {
                output = append(output, operators[len(operators)-1])
                operators = operators[:len(operators)-1]
            }
            if len(operators) == 0 {
                return nil, errors.New("incorrect input")
            }
            operators = operators[:len(operators)-1] 
        } else if isOperator(token) {
            for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
                output = append(output, operators[len(operators)-1])
                operators = operators[:len(operators)-1]
            }
            operators = append(operators, token)
        } else {
            return nil, fmt.Errorf("invalid character")
        }
    }

    for len(operators) > 0 {
        if operators[len(operators)-1] == "(" {
            return nil, errors.New("incorrect input")
        }
        output = append(output, operators[len(operators)-1])
        operators = operators[:len(operators)-1]
    }

    return output, nil
}


func evaluatepf(postfix []string) map[string]string {
    //var jsonBytes []byte
    var stack []float64
    results := make(map[string]string)
   // type Res struct {
	   // Result float64 `json:"result"`
	//}
	//type Mistakes struct {
	   // Error string`json:"error"`
	//}
    for _, token := range postfix {
        if isNumber(token) {
            num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if !isOperator(token) {
            results["error"] = "Expression is not valid"
			//otv4 := Mistakes{Error: "Expression is not valid"}
			//jsonBytes4, _ := json.Marshal(otv4)
					
			
			//return jsonBytes4
		}

		if len(stack) < 2 {
			//otv1 := Mistakes{Error: `Expression is not valid`}
			//jsonBytes7, _ := json.Marshal(otv1)
			results["error"] = "Expression is not valid"
				
			//return jsonBytes7
		}
		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		switch token {
		    case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
                    results["error"] = "Expression is not valid"
					//otv2 := Mistakes{Error: `Expression is not valid`}
					//jsonBytes2, _ := json.Marshal(otv2)
				
				
					//return jsonBytes2
				}
				stack = append(stack, a/b)
			default:
                results["error"] = "Expression is not valid"
				//otv3 := Mistakes{Error: "Expression is not valid"}
				//jsonBytes3, _ := json.Marshal(otv3)
				
				
				//return jsonBytes3
		}
	

		if len(stack) != 1 {
            results["error"] = "Expression is not valid"
			//otv5 := Mistakes{Error: `Expression is not valid`}
			//jsonBytes5, _ := json.Marshal(otv5)
					
					
			//return jsonBytes5
		}
        results["result"] = fmt.Sprintf("%f", stack[0])
        return results
		//otv := Res{Result: stack[0]}
		//jsonBytes, _ := json.Marshal(otv) 
		//return jsonBytes
	}
    return results
	//return jsonBytes
}

func isNumber(token string) bool {
    if _, err := strconv.ParseFloat(token, 64); err == nil {
        return true
    }
    return false
}

func isOperator(token string) bool {
    return token == "+" || token == "-" || token == "*" || token == "/"
}

func precedence(op string) int {
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    default:
        return 0
    }
}
func Calc(expression []byte) map[string]string{
	type Calcul struct { 
		Expression string `json:"expression"`
	}
	
	var expr []Calcul
    f := make(map[string]string)
    //var jsonData []byte
    //jsonStr := `{"ex": expression}`
	err := json.Unmarshal([]byte(expression), &expr)
	if err != nil {
		return f
	}
    tokens := tokeng(expr.Expression)
    pf, err := infpf(tokens)
    if err != nil {
        return f
    }
    return evaluatepf(pf)
}