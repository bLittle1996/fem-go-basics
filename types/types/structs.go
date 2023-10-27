package types

type MyFirstStruct struct {
	property string
	wow      bool
}

type MyFirstInterface interface {
}

func randomFunctionToWriteCodeInMyFriends() {
	var someStruct MyFirstStruct
	someStruct.wow = true

	struct2 := MyFirstStruct{property: "yep", wow: true}
	struct2.property = "wowowo"

	emptyBoi := MyFirstStruct{}

	if emptyBoi.property == "" {
		print("sadge")
	}
}

func (obj MyFirstStruct) Print() {

	print(obj.property)
}
