package tengo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/d5/tengo/v2/allure"
	"github.com/spyzhov/ajson"
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
    {
        Name: "_caseCombine",
        Value: _caseCombine,
    },
    {
        Name: "_emptyRef",
        Value: _emptyRef,
    },
    {
        Name: "_toggleError",
        Value: _toggleError,
    },
    {
        Name: "_caseRequest",
        Value: _caseRequest,
    },
    {
        Name: "_caseAssert",
        Value: _caseAssert,
    },
    {
        Name: "_caseExtract",
        Value: _caseExtract,
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

type StepTrace struct{
    Steps       []*allure.Step
}
func (st *StepTrace)empty()bool{ return len(st.Steps)==0}
func (st *StepTrace)current() *allure.Step{ return st.Steps[len(st.Steps)-1]} 
func (st *StepTrace)root() *allure.Step{ return st.Steps[0]} 
func (st *StepTrace)push(s *allure.Step) { st.Steps = append(st.Steps, s)} 
func (st *StepTrace)pop()   *allure.Step { 
    assert(len(st.Steps)>0)
    ret:= st.current()
    st.Steps = st.Steps[:len(st.Steps)-1]
    return ret
} 
func (st *StepTrace)fail(detail *allure.StatusDetail){
    for i:=len(st.Steps)-1;i>=0;i--{// fail all the steps
        if i==len(st.Steps)-1{
            st.Steps[i].Fail(detail)
            detail = &allure.StatusDetail{
                Message: fmt.Sprintf("step [%s] failure",  st.Steps[i].Name),
                Trace: detail.Trace,
            }
        }else{
            st.Steps[i].Fail(detail)
        }
    }
}

type C struct{
    Case *allure.Result
    LocalCase   bool      // whether the case is assigned in the same scope of the symbol
                          // object C was duplicated when propagating across block/func
    Error       bool      // whether error propagation is accepted 
    Trace       *allure.StatusDetail //error trace

    //step stack of the current case
    St          *StepTrace
    Domain      string
    Headers     map[string]string
    Suite       string
    ParentSuite string
    SubSuite    string
    Epic        string
    Story       string
    Feature     string
    Package     string
    
    
    
    
}

func (c *C) addStep(name string) {
    newStep := allure.NewSimpleStep(name)
    newStep.Begin()
    if current:=c.currentStep(); current!=nil{
        current.WithChild(newStep)
    }else{
        //TODO append on close
        c.Case.Steps = append(c.Case.Steps, newStep)
    }
    c.St.push(newStep)
}
func (c *C) popStep() *allure.Step{
    return c.St.pop()
}
func (c *C) currentStep() *allure.Step{
    if c.St.empty(){
        return nil
    }else{
        return c.St.current()
    }
}
func (c *C) caseAvailable() bool{
    return c.Case!=nil && c.Case.Status==""
}
func (c *C) String() string{
    ret:="" 
    if c.Case!=nil{
        j,_:=c.Case.ToJSON()
        ret+=fmt.Sprintf("Case %v LocalCase %v", string(j), c.LocalCase)
    }
    
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
    if c.Story!=""{
        ret += fmt.Sprintf(" Package(%s)", c.Package)
    }
    if c.Trace!=nil{
        ret += fmt.Sprintf(" Trace(%s)", c.Trace)
    }
    if c.Error{
        ret += fmt.Sprintf(" Error(%v)", c.Error)
    }
    
    
    return ret 
}

func _toggleError(args ...Object)(Object , error){
    if err := validateArgs(1, []string{"native-ref(*tengo.C)"}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    c.Error=true

    return args[0],nil
}
//copy the case
//block/call
func _caseCopy(args ...Object) (Object, error){
   if err := validateArgs(1, []string{"native-ref(*tengo.C)"}, args...); err!=nil{
        return nil, err
    }
    origin:=args[0].(*NativeReference).Value.(*C)
    newHeaders := map[string]string{}
    for k,v := range origin.Headers{
        newHeaders[k]= v
    }
    return &NativeReference{
        Name: "case",
        Value: &C{
            Case: origin.Case,
            LocalCase: false,
            St: origin.St,

            Headers: newHeaders,
            Domain: origin.Domain,
            ParentSuite: origin.ParentSuite,
            Suite:  origin.Suite,
            SubSuite: origin.SubSuite,

            Epic: origin.Epic,
            Feature: origin.Feature,
            Story: origin.Story,
            Package: origin.Package,
            Trace: origin.Trace,


        },//allure.NewResult(caseName.Value, fullName.Value),
    },nil
    
    
}
//build a new case
func _caseNew(args ...Object) (Object, error){
    if err := validateArgs(3, []string{"native-ref(*tengo.C)", "string", "string", "string"}, args...); err!=nil{
        return nil, err
    }
    caseName,_ := args[1].(*String)
    fullName,_ := args[2].(*String)

    origin:= args[0].(*NativeReference).Value.(*C)
    if origin.caseAvailable(){
        _caseDone(args[0], &String{Value: "Pass"})
    }
    ret,_ := _caseCopy(args[0])

    c := ret.(*NativeReference).Value.(*C)
    c.LocalCase=true
    c.Case= allure.NewResult(caseName.Value, fullName.Value)
    c.St  = &StepTrace{}
    if len(args)>3{
        desc:=args[3].(*String)
        c.Case.Description = desc.Value
    }
    c.Case.Begin()
    if origin.Trace!=nil{
        //skip it now
        // _caseDone(ret, &String{Value: "Fail"}, &String{ Value: origin.Trace.Message}, &String{ Value: origin.Trace.Trace})
        caseSkip(c, origin.Trace)
        
        
    }
    return ret,nil
    
    
    
}
//implicitly finish the case
func _caseClose(args ...Object) (Object, error){
    if err := validateArgs(2, []string{"native-ref(*tengo.C)", "native-ref(*tengo.C)", "bool"}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    //nil or closed or not allocated in the current scope 
    //nothing to do
    if c.Case==nil || c.Case.Status!="" || 
        !c.LocalCase {
        //TODO extract function for status check 
    }else{
        applyTags(c)
        c.Case.Done()
    }
    
    param := args[1].(*NativeReference).Value.(*C)
    if c.Trace!=nil{
        //TODO extract method to fail the C
        isBlock := args[2].(*Bool).value
        if (isBlock || param.Error) && 
            param.Trace == nil {
            param.Trace = c.Trace
        }
    }
    param.Error=false//reset

    
    
    return nil, nil
}

//if the local is nil, alloc a new one and return
func _setLocal(args ...Object)(Object, error){
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
    case "Package":
        c.Package=arg0
    }
}
func _emptyRef(args ...Object)(Object, error){
    return &NativeReference{
        Name: "case",
        Value: &C{},
    },nil
}

//attachment/param/step set to current step/case
//circumstances of no case or step
//1. case undefined
//2. case finished
func _caseAttachment(args ...Object) (Object, error){
    if err := validateArgs(4, []string{"native-ref(*tengo.C)", "string", "string", "string"}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    if !c.caseAvailable(){
        warnNoCaseOrStep(c)
        return nil,nil
    }

    name := args[1].(*String).Value
    //TODO check mime type
    mime := args[2].(*String).Value
    content := args[3].(*String).Value

    if !c.St.empty() {
        c.currentStep().WithAttachments(allure.NewAttachment(name, allure.MimeType(mime), []byte(content)))
            
    }else{
        c.Case.Attachments = append(c.Case.Attachments, (allure.NewAttachment(name, allure.MimeType(mime), []byte(content))))
    }
    
    return nil, nil
}
//param
func _caseParameter(args ...Object) (Object, error){
   if err := validateArgs(3, []string{"native-ref(*tengo.C)", "string"}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    key := args[1].(*String).Value
    value := fmt.Sprintf("%v", ToInterface(args[2]))
    return caseParameter(c, key, value)
}
func caseParameter(c *C, k string, v any)(Object , error){
    if !c.caseAvailable(){
        warnNoCaseOrStep(c)
        return nil,nil
    }
    if !c.St.empty(){
        c.currentStep().WithParameters(allure.NewParameter(k, v))
    }else{
        c.Case.Parameters = append(c.Case.Parameters, allure.NewParameter(k, v))
    }
    return nil,nil
}
//a stack is used to store the step
func _caseStep(args ...Object) (Object, error){
   if err := validateArgs(2, []string{"native-ref(*tengo.C)", "string"}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    if !c.caseAvailable(){
        warnNoCaseOrStep(c)
        return nil,nil
    }
    stepName := args[1].(*String).Value
    c.addStep(stepName)
    return nil, nil
}
//manually pass/fail current case/step
//PassStep()/Pass()
//FailCase("message")/FailStep("message")
//optional param msg/trace
func _caseDone(args ...Object) (Object, error){
   if err := validateArgs(2, []string{"native-ref(*tengo.C)", "string", "string", "string"}, args...); err!=nil{
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
    
    if !c.caseAvailable(){
        warnNoCaseOrStep(c)
        return nil,nil
    }
    switch(t.Value){
    case "Fail","Pass": //case
        //finish the steps
        if !c.St.empty(){
            //finish the cases, the sub steps are done recursively
            c.St.root().Done()
        }
        //finish the case
        if t.Value=="Fail"{
            //fail the case
            msg:=args[2].(*String).Value
            var trace string
            if len(args)>=4{
                trace = args[3].(*String).Value
            }else{
                trace= CurrentVM().getStackTrace()
            }
            c.Case.StatusDetails = allure.StatusDetail{Message: msg, Trace: trace}
            c.Case.Status        = allure.Failed
            applyTags(c)
            c.Case.Done()
            c.Trace = &allure.StatusDetail{
                Message: fmt.Sprintf("case [%s] failure", c.Case.Name),
                Trace: trace,
            }
            
        }else{
            applyTags(c)
            c.Case.Done()
        }
    case "FailStep","PassStep": //step
        //check steps empty 
        if c.St.empty(){
            //no steps
            warnNoCaseOrStep(c)
            return nil,nil
        }else{
            if t.Value=="FailStep"{
                msg:=args[2].(*String).Value
                trace:=CurrentVM().getStackTrace()
                detail := &allure.StatusDetail{Message: msg, Trace: trace}

                c.St.fail(detail)
                //failing a step makes the full case fail
                //TODO use pointer
                c.Case.StatusDetails = *detail
                c.Case.Status        = allure.Failed
                applyTags(c)
                c.Case.Done()
                c.Trace = &allure.StatusDetail{
                    Message: fmt.Sprintf("case [%s] failure", c.Case.Name),
                    Trace: trace,
                }
                
            }else{
                c.St.pop().Done()
            }
        }
    }
    return nil, nil
}

//return a copy the param_case and override the data with var_case
//param: param_case var_case
//TODO merge into method _caseCopy
func _caseCombine(args ...Object) (Object, error){
    if err := validateArgs(2, []string{"native-ref(*tengo.C)", "native-ref(*tengo.C)"}, args...); err!=nil{
        return nil, err
    }
    newCase, err:= _caseCopy(args[0])
    if err!=nil{return nil, err}
    closure := args[1].(*NativeReference).Value.(*C)
    c:= newCase.(*NativeReference).Value.(*C)

    //override
    if closure.Domain!=""{ c.Domain=closure.Domain}
    for k,v := range closure.Headers{
        c.Headers[k]= v
    }
    if closure.Suite!=""{ c.Suite=closure.Suite}
    if closure.ParentSuite!=""{ c.ParentSuite=closure.ParentSuite}
    if closure.SubSuite!=""{ c.SubSuite=closure.SubSuite}
    if closure.Epic!=""{ c.Epic=closure.Epic}
    if closure.Story!=""{ c.Story=closure.Story}
    if closure.Feature!=""{ c.Feature=closure.Feature}
    if closure.Package!=""{ c.Package=closure.Package}

    return newCase, nil
}

//http request
//request is a step
//TODO how to deal with nil returned by this func 
func _caseRequest(args ...Object)(Object, error){
    if err := validateArgs(3, []string{"native-ref(*tengo.C)", "string", "string", "", "map"}, args...); err!=nil{
        return nil, err
    }
    //request things
    c:= args[0].(*NativeReference).Value.(*C)
    if !c.caseAvailable(){
        warnNoCaseOrStep(c)
        return nil, nil
    }
    _caseStep(args[0], &String{Value: "http request"})

    
    path:=args[1].(*String).Value
    uri := c.Domain
    if uri==""{
        return nil, errors.New("Domain not set")

    }
    if strings.HasSuffix(uri, "/"){
        uri = uri[:len(uri)-1]
    }
    uri = uri + path
    method:=strings.ToUpper(args[2].(*String).Value)
    caseParameter(c, "url", uri)
    caseParameter(c, "method", method)

    //new req
    var body io.Reader
    if len(args)>=4{
        //TODO use json in tengo stdlib
        bodyMap:= args[3]
        if !bodyMap.IsFalsy(){
            rawBody:= wrapMarshal((bodyMap))
            caseParameter(c, "request-body", rawBody)
            body= strings.NewReader(rawBody)
        }
    }
    req,err:= http.NewRequest(method, uri, body)
    if err!=nil{return nil,err}


    //header
    if len(args)>=5{
        for k,v := range args[4].(*Map).Value{
            //TODO why are string quoted
            s:=v.String()
            if st,ok:= v.(*String); ok{
                //get the raw string
                s=st.Value
            }
            req.Header.Add(k, s)
        }
    }
    for k,v := range c.Headers{
        req.Header.Add(k, v)
    }
    caseParameter(c, "request-headers", wrapMarshal(buildHeader(req.Header)))
    resp, err := http.DefaultClient.Do(req)
    defer resp.Body.Close()
    if err!=nil{
        //do fail step
        _caseDone(args[0], wrapFromInterface("FailStep"), wrapFromInterface(err.Error()))
        return nil, nil
    }
    if resp.StatusCode >= 400{
        //resp code over 400 are considered error
        //do fail step
        _caseDone(args[0], wrapFromInterface("FailStep"), wrapFromInterface(fmt.Sprintf("error response code : %d" , resp.StatusCode)))
    }
    

    //step info 
    //uri/method/body/resp
    ret:=map[string]Object{}

    header:= buildHeader((resp.Header))
    ret["status"]= &Int{Value: int64(resp.StatusCode)}
    ret["url"]= &String{Value: resp.Request.RequestURI}
    ret["method"]= &String{Value: resp.Request.Method}
    //XXX use the FromInterface func to convert all types 
    ret["headers"]= header
    responseBody, err := io.ReadAll(resp.Body)
    if err!=nil{return nil, err}
    ret["body"]=&String{Value: string(responseBody)}


    //response status/headers/body/
    caseParameter(c, "status", fmt.Sprintf("%v", resp.StatusCode))
    caseParameter(c, "response-headers", wrapMarshal(header))
    caseParameter(c, "response-body", string(responseBody))
    _caseDone(args[0], wrapFromInterface("PassStep"))
    return &Map{
        Value: ret,
    },nil
}

//assert step
//assert equal   0 expr arg0 arg1  
//assert not eq  1 expr arg0 arg1  
//assert that    2 expr arg0 nil
func _caseAssert(args ...Object)(Object, error){
    if err := validateArgs(4, []string{"native-ref(*tengo.C)", "int", "string", ""}, args...); err!=nil{
        return nil, err
    }
    c:= args[0].(*NativeReference).Value.(*C)
    if !c.caseAvailable(){
        warnNoCaseOrStep(c)
        return nil, nil
    }

    _caseStep(args[0], wrapFromInterface("assertion"))
    t:= args[1].(*Int).Value// assert result
    expr:= args[2].(*String).Value

    caseParameter(c, "expr" , expr)
    for i:=3;i<len(args);i++{
        caseParameter(c, fmt.Sprintf("arg%d", i-3), ToInterface(args[i]))
    }
    if t==0{
        check:= args[3].(*Bool).value
        if check{
            _caseDone(args[0], wrapFromInterface("PassStep"))
        }else{
            _caseDone(args[0], wrapFromInterface("FailStep"), wrapFromInterface("assertion failed"))
        }
    }else{
        expect:= true
        if t==2 { expect=false }
        check:= args[3].Equals(args[4])
        if check==expect{
            _caseDone(args[0], wrapFromInterface("PassStep"))
        }else{
            _caseDone(args[0], wrapFromInterface("FailStep"), wrapFromInterface("assertion failed"))
        }
    }

    
    

    return nil, nil
}
//extract data from json
func _caseExtract(args ...Object)(Object, error){
    if err := validateArgs(3, []string{"native-ref(*tengo.C)", "string", "string"}, args...); err!=nil{
        return nil, err
    }
    c:=args[0].(*NativeReference).Value.(*C)
    if !c.caseAvailable(){
        warnNoCaseOrStep(c)
        return nil,nil
    }
    _caseStep(args[0], wrapFromInterface("extract data"))

    expr:= args[1].(*String).Value
    obj:= args[2].(*String).Value
    caseParameter(c, "expr", expr)
    caseParameter(c, "data", obj )
    root, err := ajson.Unmarshal([]byte(obj))
    if err!=nil{
        //fail step
        return _caseDone(args[0], wrapFromInterface("FailStep"), wrapFromInterface(err.Error()))
    }
    nodes, err:= root.JSONPath(expr)
    if err!=nil{
        //fail step
        return _caseDone(args[0], wrapFromInterface("FailStep"), wrapFromInterface(err.Error()))
    }
    result:= ajson.ArrayNode("", nodes)
    j,_ := ajson.Marshal(result)
    caseParameter(c, "result", string(j))
    _caseDone(args[0], wrapFromInterface("PassStep"))
    

    //TODO reduce the times of marshalling and unmarshalling
    var o interface{}
    json.Unmarshal(j, &o)
    return wrapFromInterface(o), nil
}
func buildHeader(h http.Header) *Map{
    m := map[string]Object{}
    for k,v := range h{
        m[k]  = &String{ Value: strings.Join(v, "")}
    }

    return &Map{
        Value: m,
    }
}

func applyTags(c *C){
    if c.Case!=nil{
        // if c.
            
        if c.Suite       !=""{c.Case.WithSuite(c.Suite)}
        if c.ParentSuite !=""{c.Case.WithParentSuite(c.ParentSuite)}
        if c.SubSuite    !=""{c.Case.WithSubSuites(c.SubSuite)}
        if c.Epic        !=""{c.Case.WithEpic(c.Epic)}
        if c.Story       !=""{c.Case.WithStory(c.Story)}
        if c.Feature     !=""{c.Case.WithFeature(c.Feature)}
        if c.Package     !=""{c.Case.WithPackage(c.Package)}
    }
}

//case skip only happends at the begining of case
func caseSkip(c *C, details *allure.StatusDetail){
    if c.Case!=nil{
        applyTags(c)
        c.Case.Status=allure.Skipped
        c.Case.StatusDetails = *details
        c.Case.Done()
    }
}

func warnNoCaseOrStep(c *C) {
    //if the status is failed or skipped, may be the preceding operation cause the operation 
    //unavailable, discard it
    if c.Case==nil || c.Case.Status=="passed"{
        fmt.Println("Warning: no case or step found")
        fmt.Println(CurrentVM().getStackTrace())
    }
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

//types  acceptable types splited by /,  empty string for any type 
func validateArgs(requiredArgs int, types []string, args ...Object) error{
    if len(args)<requiredArgs{
        return ErrWrongNumArguments 

    }
    maxLen:=len(types)
    for i,arg := range args{
        if i >= maxLen{
            break
        }
        candidates:= strings.Split(types[i], "/")// multiple types, mainly undefined values
        if ! slices.ContainsFunc(candidates, func(s string)bool{
            return s =="" || s==arg.TypeName()// type is "" or type equals
        }){
            return ErrInvalidArgumentType{
                Name:     fmt.Sprintf("arg%d", i),
                Expected: types[i],
                Found:    args[i].TypeName(),
            }
        }
    }
    return nil
}
func wrapMarshal(i interface{})string{
    var inf interface{}
    switch i:=i.(type){
    case Object:
        inf = ToInterface(i)
        return wrapMarshal(inf)
    case []byte:
        inf = string(i)
    default:
        inf = i
    }
    bytes, err := json.Marshal(inf)
    if err!=nil{
        panic(err)
    }
    return string(bytes)

}
//convert an golang object to tengo object
func wrapFromInterface(i interface{})Object{
    ret, err := FromInterface(i)
    if err!=nil{
        panic("error convert from interface:" + err.Error())
    }
    return ret
}

func assert(v bool){
    if (!v){
        panic("assertion failed")
    }
}
