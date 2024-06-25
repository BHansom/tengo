package tengo

import (
	"fmt"

	"github.com/d5/tengo/v2/allure"
)

var builtinFuncs = []*BuiltinFunction{
	{
		Name:  "len",
		Value: builtinLen,
	},
	{
		Name:  "copy",
		Value: builtinCopy,
	},
	{
		Name:  "append",
		Value: builtinAppend,
	},
	{
		Name:  "delete",
		Value: builtinDelete,
	},
	{
		Name:  "splice",
		Value: builtinSplice,
	},
	{
		Name:  "string",
		Value: builtinString,
	},
	{
		Name:  "int",
		Value: builtinInt,
	},
	{
		Name:  "bool",
		Value: builtinBool,
	},
	{
		Name:  "float",
		Value: builtinFloat,
	},
	{
		Name:  "char",
		Value: builtinChar,
	},
	{
		Name:  "bytes",
		Value: builtinBytes,
	},
	{
		Name:  "time",
		Value: builtinTime,
	},
	{
		Name:  "is_int",
		Value: builtinIsInt,
	},
	{
		Name:  "is_float",
		Value: builtinIsFloat,
	},
	{
		Name:  "is_string",
		Value: builtinIsString,
	},
	{
		Name:  "is_bool",
		Value: builtinIsBool,
	},
	{
		Name:  "is_char",
		Value: builtinIsChar,
	},
	{
		Name:  "is_bytes",
		Value: builtinIsBytes,
	},
	{
		Name:  "is_array",
		Value: builtinIsArray,
	},
	{
		Name:  "is_immutable_array",
		Value: builtinIsImmutableArray,
	},
	{
		Name:  "is_map",
		Value: builtinIsMap,
	},
	{
		Name:  "is_immutable_map",
		Value: builtinIsImmutableMap,
	},
	{
		Name:  "is_iterable",
		Value: builtinIsIterable,
	},
	{
		Name:  "is_time",
		Value: builtinIsTime,
	},
	{
		Name:  "is_error",
		Value: builtinIsError,
	},
	{
		Name:  "is_undefined",
		Value: builtinIsUndefined,
	},
	{
		Name:  "is_function",
		Value: builtinIsFunction,
	},
	{
		Name:  "is_callable",
		Value: builtinIsCallable,
	},
	{
		Name:  "type_name",
		Value: builtinTypeName,
	},
	{
		Name:  "format",
		Value: builtinFormat,
	},
	{
		Name:  "range",
		Value: builtinRange,
	},
    {
        Name: "glob",
        Value: builtinGlob,
    },

    {
        Name: "_caseNew",
        Value: _caseNew,
    },
    {
        Name: "_caseClose",
        Value: _caseClose,
    },
    {
        Name: "_caseCopy",
        Value: _caseCopy,
    },
    {
        Name: "_setLocal",
        Value: _setLocal,
    },
    {
        Name: "_caseAttachment",
        Value: _caseAttachment,
    },
    {
        Name: "_caseParameter",
        Value: _caseParameter,
    },
    {
        Name: "_caseStep",
        Value: _caseStep,
    },
    {
        Name: "_caseDone",
        Value: _caseDone,
    },
}


func builtinGlob(args ...Object)(Object, error){
    fmt.Println(CurrentVM())
    return nil, nil
}

// GetAllBuiltinFunctions returns all builtin function objects.
func GetAllBuiltinFunctions() []*BuiltinFunction {
	return append([]*BuiltinFunction{}, builtinFuncs...)
}

func builtinTypeName(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	return &String{Value: args[0].TypeName()}, nil
}

func builtinIsString(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*String); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsInt(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Int); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsFloat(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Float); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBool(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bool); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsChar(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Char); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBytes(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bytes); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsArray(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Array); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsImmutableArray(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*ImmutableArray); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsMap(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Map); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsImmutableMap(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*ImmutableMap); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsTime(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Time); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsError(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Error); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsUndefined(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0] == UndefinedValue {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsFunction(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	switch args[0].(type) {
	case *CompiledFunction:
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsCallable(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0].CanCall() {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsIterable(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0].CanIterate() {
		return TrueValue, nil
	}
	return FalseValue, nil
}

// len(obj object) => int
func builtinLen(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Array:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *ImmutableArray:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *String:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *Bytes:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *Map:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *ImmutableMap:
		return &Int{Value: int64(len(arg.Value))}, nil
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array/string/bytes/map",
			Found:    arg.TypeName(),
		}
	}
}

//range(start, stop[, step])
func builtinRange(args ...Object) (Object, error) {
	numArgs := len(args)
	if numArgs < 2 || numArgs > 3 {
		return nil, ErrWrongNumArguments
	}
	var start, stop, step *Int

	for i, arg := range args {
		v, ok := args[i].(*Int)
		if !ok {
			var name string
			switch i {
			case 0:
				name = "start"
			case 1:
				name = "stop"
			case 2:
				name = "step"
			}

			return nil, ErrInvalidArgumentType{
				Name:     name,
				Expected: "int",
				Found:    arg.TypeName(),
			}
		}
		if i == 2 && v.Value <= 0 {
			return nil, ErrInvalidRangeStep
		}
		switch i {
		case 0:
			start = v
		case 1:
			stop = v
		case 2:
			step = v
		}
	}

	if step == nil {
		step = &Int{Value: int64(1)}
	}

	return buildRange(start.Value, stop.Value, step.Value), nil
}

func buildRange(start, stop, step int64) *Array {
	array := &Array{}
	if start <= stop {
		for i := start; i < stop; i += step {
			array.Value = append(array.Value, &Int{
				Value: i,
			})
		}
	} else {
		for i := start; i > stop; i -= step {
			array.Value = append(array.Value, &Int{
				Value: i,
			})
		}
	}
	return array
}

func builtinFormat(args ...Object) (Object, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, ErrWrongNumArguments
	}
	format, ok := args[0].(*String)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &String{Value: s}, nil
}

func builtinCopy(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	return args[0].Copy(), nil
}

func builtinString(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*String); ok {
		return args[0], nil
	}
	v, ok := ToString(args[0])
	if ok {
		if len(v) > MaxStringLen {
			return nil, ErrStringLimit
		}
		return &String{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinInt(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Int); ok {
		return args[0], nil
	}
	v, ok := ToInt64(args[0])
	if ok {
		return &Int{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinFloat(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Float); ok {
		return args[0], nil
	}
	v, ok := ToFloat64(args[0])
	if ok {
		return &Float{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinBool(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bool); ok {
		return args[0], nil
	}
	v, ok := ToBool(args[0])
	if ok {
		if v {
			return TrueValue, nil
		}
		return FalseValue, nil
	}
	return UndefinedValue, nil
}

func builtinChar(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Char); ok {
		return args[0], nil
	}
	v, ok := ToRune(args[0])
	if ok {
		return &Char{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinBytes(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}

	// bytes(N) => create a new bytes with given size N
	if n, ok := args[0].(*Int); ok {
		if n.Value > int64(MaxBytesLen) {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: make([]byte, int(n.Value))}, nil
	}
	v, ok := ToByteSlice(args[0])
	if ok {
		if len(v) > MaxBytesLen {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinTime(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Time); ok {
		return args[0], nil
	}
	v, ok := ToTime(args[0])
	if ok {
		return &Time{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

// append(arr, items...)
func builtinAppend(args ...Object) (Object, error) {
	if len(args) < 2 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Array:
		return &Array{Value: append(arg.Value, args[1:]...)}, nil
	case *ImmutableArray:
		return &Array{Value: append(arg.Value, args[1:]...)}, nil
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    arg.TypeName(),
		}
	}
}

// builtinDelete deletes Map keys
// usage: delete(map, "key")
// key must be a string
func builtinDelete(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen != 2 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Map:
		if key, ok := args[1].(*String); ok {
			delete(arg.Value, key.Value)
			return UndefinedValue, nil
		}
		return nil, ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map",
			Found:    arg.TypeName(),
		}
	}
}

// builtinSplice deletes and changes given Array, returns deleted items.
// usage:
// deleted_items := splice(array[,start[,delete_count[,item1[,item2[,...]]]])
func builtinSplice(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return nil, ErrWrongNumArguments
	}

	array, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}
	arrayLen := len(array.Value)

	var startIdx int
	if argsLen > 1 {
		arg1, ok := args[1].(*Int)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int",
				Found:    args[1].TypeName(),
			}
		}
		startIdx = int(arg1.Value)
		if startIdx < 0 || startIdx > arrayLen {
			return nil, ErrIndexOutOfBounds
		}
	}

	delCount := len(array.Value)
	if argsLen > 2 {
		arg2, ok := args[2].(*Int)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int",
				Found:    args[2].TypeName(),
			}
		}
		delCount = int(arg2.Value)
		if delCount < 0 {
			return nil, ErrIndexOutOfBounds
		}
	}
	// if count of to be deleted items is bigger than expected, truncate it
	if startIdx+delCount > arrayLen {
		delCount = arrayLen - startIdx
	}
	// delete items
	endIdx := startIdx + delCount
	deleted := append([]Object{}, array.Value[startIdx:endIdx]...)

	head := array.Value[:startIdx]
	var items []Object
	if argsLen > 3 {
		items = make([]Object, 0, argsLen-3)
		for i := 3; i < argsLen; i++ {
			items = append(items, args[i])
		}
	}
	items = append(items, array.Value[endIdx:]...)
	array.Value = append(head, items...)

	// return deleted items
	return &Array{Value: deleted}, nil
}

type C struct{
    Case *allure.Result
    LocalCase   bool      // whether the case is assigned in the same scope of the symbol
                          // object C was duplicated when propagating across block/func
    //step stack of the current case
    Steps       []*allure.Step
    Domain      string
    Headers     map[string]string
    Suite       string
    ParentSuite string
    SubSuite    string
    Epic        string
    Story       string
    Feature     string
    
    
    
    
}
func (c *C) String() string{
    ret:="" 
    if c.Case!=nil{
        j,_:=c.Case.ToJSON()
        ret+=fmt.Sprintf("Case %v LocalCase %v", string(j), c.LocalCase)
    }
    ret+=fmt.Sprintf("LocalCase %v", c.LocalCase)
    
    if c.Domain!=""{
        ret += fmt.Sprintf(" Domain(%s)", c.Domain)
    }
    if len(c.Headers)>0{
        ret += fmt.Sprintf(" headers(%s)", c.Headers)
    }
    if c.ParentSuite!=""{
        ret += fmt.Sprintf(" ParentSuite(%s)", c.ParentSuite)
    }
    if c.Suite!=""{
        ret += fmt.Sprintf(" Suite(%s)", c.Suite)
    }
    if c.SubSuite!=""{
        ret += fmt.Sprintf(" SubSuite(%s)", c.SubSuite)
    }
    if c.Epic!=""{
        ret += fmt.Sprintf(" Epic(%s)", c.Epic)
    }
    if c.Feature!=""{
        ret += fmt.Sprintf(" Feature(%s)", c.Feature)
    }
    if c.Story!=""{
        ret += fmt.Sprintf(" Story(%s)", c.Story)
    }
    
    return ret 
}
//copy the case
//block/call
func _caseCopy(args ...Object) (Object, error){
    if args[0]==UndefinedValue{
        return nil, nil
    }
    if err := validateArgs(1, []string{"native-ref(*tengo.C)"}, args...); err!=nil{
        return nil, err
    }
    origin:=args[0].(*NativeReference).Value.(*C)
    newHeaders := map[string]string{}
    for k,v := range origin.Headers{
        newHeaders[k]= v
    }
    steps:= []*allure.Step{}
    //the step stack is duplicated
    for _,s := range origin.Steps{
        steps = append(steps, s)
    }
    return &NativeReference{
        Name: "case",
        Value: &C{
            Case: origin.Case,
            LocalCase: false,
            //should this be pointer
            Steps: steps,

            Headers: newHeaders,
            Domain: origin.Domain,
            ParentSuite: origin.ParentSuite,
            Suite:  origin.Suite,
            SubSuite: origin.SubSuite,

            Epic: origin.Epic,
            Feature: origin.Feature,
            Story: origin.Story,


        },//allure.NewResult(caseName.Value, fullName.Value),
    },nil
    
    
}
//build a new case
func _caseNew(args ...Object) (Object, error){
    assert(len(args)>0)
    if args[0]!=UndefinedValue{
        if err := validateArgs(3, []string{"native-ref(*tengo.C)", "string", "string", "string"}, args...); err!=nil{
            return nil, err
        }
    }
 //    if len(args)< 2 {
	// 	return nil, ErrWrongNumArguments
 //    }
 //    if !ok {
	// 	return nil, ErrInvalidArgumentType{
	// 		Name:     "caseName",
	// 		Expected: "string",
	// 		Found:    args[0].TypeName(),
	// 	}
	// }
 //    if !ok {
	// 	return nil, ErrInvalidArgumentType{
	// 		Name:     "fullName",
	// 		Expected: "string",
	// 		Found:    args[0].TypeName(),
	// 	}
	// }
    caseName,_ := args[1].(*String)
    fullName,_ := args[2].(*String)

    origin:= args[0]
    var ret Object
    if origin!=UndefinedValue{
        if origin.(*NativeReference).Value.(*C).Case.Status==""{
            _caseDone(origin, &String{Value: "Pass"})
        }
        ret,_ = _caseCopy(origin)
    }else{
        ret= newEmptyRef()
    }
    c := ret.(*NativeReference).Value.(*C)
    c.LocalCase=true
    c.Case= allure.NewResult(caseName.Value, fullName.Value)
    if len(args)>3{
        desc:=args[3].(*String)
        c.Case.Description = desc.Value
    }
    c.Case.Begin()
    return ret,nil
    
    
    
    // return &NativeReference{
    //     Name: "case",
    //     Value: &C{
    //         Domain: "domain",
    //         Headers: map[string]string{"conent-type":"content-type"},
    //         Suite: "suite",
    //     },//allure.NewResult(caseName.Value, fullName.Value),
    // },nil
}
//implicitly finish the case
func _caseClose(args ...Object) (Object, error){
    //case not set
    if args[0] == UndefinedValue{
        return nil, nil
    }
    if err := validateArgs(1, []string{"native-ref(*tengo.C)"}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    //nil or closed or not allocated in the current scope 
    //nothing to do
    if c.Case==nil || c.Case.Status!="" || 
        !c.LocalCase {
        return nil,nil
    }

    c.Case.Done()
    
    return nil, nil
}

//if the local is nil, alloc a new one and return
func _setLocal(args ...Object)(Object, error){
    if args[0] == UndefinedValue{
        args[0] = newEmptyRef()
    }
    if err := validateArgs(3, []string{"native-ref(*tengo.C)", "string", "string"}, args...); err!=nil{
        return nil, err
    }
    
    currentCase := args[0]

    origin:=currentCase.(*NativeReference).Value.(*C)
    field:=args[1].(*String).Value
    values:=args[2:]
    setLocal(origin, field, values...)
    
    return currentCase,nil
}

func setLocal(c *C, field string, value ...Object){
    arg0:=value[0].(*String).Value
    switch(field){
    case "Domain":
        c.Domain=arg0
    case "Header":
        if len(value)==1 || value[1].(*String).Value=="" {
            delete(c.Headers, arg0)
        }else{
            arg1:=value[1].(*String).Value
            c.Headers[arg0] = arg1
        }
    case "ParentSuite":
        c.ParentSuite=arg0
    case "Suite":
        c.Suite=arg0
    case "SubSuite":
        c.SubSuite=arg0

    case "Epic":
        c.Epic=arg0
    case "Feature":
        c.Feature=arg0
    case "Story":
        c.Story=arg0
    }
}
func newEmptyRef()(*NativeReference){
    return &NativeReference{
        Name: "case",
        Value: &C{
            Headers: map[string]string{},
        },
    }
}

//attachment/param/step set to current step/case
//circumstances of no case or step
//1. case undefined
//2. case finished
func _caseAttachment(args ...Object) (Object, error){
    if args[0] == UndefinedValue{
        warnNoCaseOrStep()
        return nil,nil
    }
    if err := validateArgs(4, []string{"native-ref(*tengo.C)", "string", "string", "string"}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    if c.Case.Status!=""{
        warnNoCaseOrStep()
        return nil,nil
    }

    name := args[1].(*String).Value
    //TODO check mime type
    mime := args[2].(*String).Value
    content := args[3].(*String).Value

    if len(c.Steps)>=0 {
        c.Steps[len(c.Steps)-1].WithAttachments(allure.NewAttachment(name, allure.MimeType(mime), []byte(content)))
            
    }else{
        c.Case.Attachments = append(c.Case.Attachments, (allure.NewAttachment(name, allure.MimeType(mime), []byte(content))))
    }
    
    return nil, nil
}
//param
func _caseParameter(args ...Object) (Object, error){
    if args[0] == UndefinedValue{
        warnNoCaseOrStep()
        return nil, nil
    }
    if err := validateArgs(3, []string{"native-ref(*tengo.C)", "string", "string"}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    if c.Case.Status!=""{
        warnNoCaseOrStep()
        return nil, nil
    }
    key := args[1].(*String).Value
    value := args[2].(*String).Value
    
    if len(c.Steps)>=0{
        //current step
        c.Steps[len(c.Steps)-1].WithParameters(allure.NewParameter(key, value))
    }else{
        c.Case.Parameters = append(c.Case.Parameters, allure.NewParameter(key, value))
    }
    
    return nil, nil
}
//a stack is used to store the step
func _caseStep(args ...Object) (Object, error){
    if args[0] == UndefinedValue{
        warnNoCaseOrStep()
        return nil,nil
    }
    if err := validateArgs(2, []string{"native-ref(*tengo.C)", "string"}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    if c.Case.Status!=""{
        warnNoCaseOrStep()
        return nil,nil
    }
    stepName := args[1].(*String).Value
    var newStep *allure.Step 
    assert(c.Case.Status=="")
    if len(c.Steps)>0{
        newStep = allure.NewSimpleStep(stepName)
        newStep.Begin()

        currentStep := c.Steps[len(c.Steps)-1]
        assert(currentStep.Status=="")
        currentStep.WithChild(newStep)
    }else{
        newStep = allure.NewSimpleStep(stepName)
        newStep.Begin()
        
        c.Case.Steps = append(c.Case.Steps, newStep)
    }
    c.Steps = append(c.Steps, newStep)
    return nil, nil
}
//manually pass/fail current case/step
//PassStep()/Pass()
//FailCase("message")/FailStep("message")
func _caseDone(args ...Object) (Object, error){
    if args[0] == UndefinedValue{
        warnNoCaseOrStep()
        return nil, nil
    }
    if err := validateArgs(2, []string{"native-ref(*tengo.C)", "string", "string"}, args...); err!=nil{
        return nil, err
    }
    c:=args[0].(*NativeReference).Value.(*C)
    t:=args[1].(*String)
    if t.Value == "Fail" || t.Value=="FailStep"{
        if len(args)< 3{
            //requires fail message as arg
            return nil, ErrWrongNumArguments
        }
    }
    
    if c.Case.Status!=""{
        warnNoCaseOrStep()
        return nil,nil
    }
    switch(t.Value){
    case "Fail","Pass": //case
        //finish the steps
        if len(c.Steps)>0{
            //finish the cases, the sub steps are done recursively
            c.Steps[0].Done()
        }
        //finish the case
        if t.Value=="Fail"{
            //fail the case
            msg:=args[2].(*String).Value
            trace:= CurrentVM().getStackTrace()
            c.Case.StatusDetails = allure.StatusDetail{Message: msg, Trace: trace}
            c.Case.Status        = allure.Failed
            c.Case.Done()
            
        }else{
            c.Case.Done()
        }
    case "FailStep","PassStep": //step
        //check steps empty 
        if len(c.Steps)==0{
            //no steps
            warnNoCaseOrStep()
            return nil,nil
        }else{
            if t.Value=="FailStep"{
                var detail *allure.StatusDetail
                msg:=args[2].(*String).Value
                trace:=CurrentVM().getStackTrace()
                for i:=len(c.Steps)-1;i>=0;i--{
                    if i==len(c.Steps)-1{
                        c.Steps[i].Fail(&allure.StatusDetail{Message: msg, Trace: trace})
                        detail = &allure.StatusDetail{
                            Message: fmt.Sprintf("step %s failure",  c.Steps[i].Name),
                            Trace: trace,
                        }
                    }else{
                        c.Steps[i].Fail(detail)
                    }
                }
                c.Steps=c.Steps[:0] 
                //failing a step makes the full case fail
                c.Case.StatusDetails = *detail
                c.Case.Status        = allure.Failed
                c.Case.Done()
                
            }else{
                c.Steps[len(c.Steps)-1].Done()
                //pop element
                c.Steps = c.Steps[:len(c.Steps)-1]
            }
        }
    }
    return nil, nil
}

func warnNoCaseOrStep() {
    fmt.Println("Warning: no case or step found")
    fmt.Println(CurrentVM().getStackTrace())
}

func builtinCase(args ...Object) (Object, error){
    return nil, nil
}
func builtinDesc(args ...Object) (Object, error){
    return nil, nil
}
func builtinDomain(args ...Object) (Object, error){
    return nil, nil
}
func builtinHeader(args ...Object) (Object, error){
    return nil, nil
}
func builtinAttachment(args ...Object) (Object, error){
    return nil, nil
}
func builtinStep(args ...Object) (Object, error){
    return nil, nil
}
func builtinParameter(args ...Object) (Object, error){
    return nil, nil
}
func builtinFail(args ...Object) (Object, error){
    return nil, nil
}
func builtinPass(args ...Object) (Object, error){
    return nil, nil
}

func validateArgs(requiredArgs int, types []string, args ...Object) error{
    if len(args)<requiredArgs{
        return ErrWrongNumArguments 

    }
    maxLen:=len(types)
    for i,arg := range args{
        if i >= maxLen{
            break
        }
        
        if arg !=nil && arg.TypeName()!=types[i]{
            return ErrInvalidArgumentType{
                Name:     "arg" + string(i),
                Expected: types[i],
                Found:    args[i].TypeName(),
            }
        }
    }
    return nil
}

func assert(v bool){
    if (!v){
        panic("assertion failed")
    }
}
