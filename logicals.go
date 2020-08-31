package main
type And struct{
	condition1 Boolean
	condition2 Boolean
}

type Or struct{
	condition1 Boolean
	condition2 Boolean
}

type Is struct{
	condition1 Boolean
	condition2 Boolean
}

type Not struct{
	condition1 Boolean
	condition2 Boolean
}

type If struct{
	condition1 Boolean
	action *func()
}

type Else struct{
	action *func()
}

func (and And) EvaluateLogic() bool{
	return and.condition1.Bin == true && and.condition2.Bin == true
}
func (or Or) EvaluateLogic() bool{
	return or.condition1.Bin == true || or.condition2.Bin == true
}

func (is Is) EvaluateLogic() bool{
	return is.condition1.Bin == is.condition2.Bin
}

func (not Not) EvaluateLogic() bool{
	return not.condition1.Bin != not.condition2.Bin
}

func RunIf(condition Boolean, true func(), false func()){
	if condition.Bin {
		true()
	} else{
		false()
	}
}
/*func And(string1 string, string2 string) string{
	return "("+string1+")" + "&&" + "("+ string2 + ")"
}

func Or(string1 string, string2 string) string{
	return "("+string1+")" + "||" + "("+ string2 + ")"
}


func Is(string1 string, string2 string) string{
	return "("+string1+")" + "==" + "("+ string2 + ")"
}


func Not(string1 string, string2 string) string{
	return "("+string1+")" + "!=" + "("+ string2 + ")"
}

func If(conditional string, then string) string{
	return "if ("+conditional+"){ "+then+"}"
}
func Else(then string) string{
	return "else {"+then+"}"
}*/