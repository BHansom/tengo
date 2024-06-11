package tengo

// SymbolScope represents a symbol scope.
type SymbolScope string
type ApiTestVariable string

// List of symbol scopes
const (
	ScopeGlobal  SymbolScope = "GLOBAL"
	ScopeLocal   SymbolScope = "LOCAL"
	ScopeBuiltin SymbolScope = "BUILTIN"
	ScopeFree    SymbolScope = "FREE"
    
    // local variable for api tests
    // all case related data stored in this var
    // data/case/error/initilized
    VarCase   ApiTestVariable = "var_case"
    //inheritable data  label/request
    VarData   ApiTestVariable = "var_data"
    //VarDomain ApiTestVariable = "var_domain"
    //VarHeader ApiTestVariable = "var_header"
    //Step      ApiTestVariable = "var_step"
    ////all the labels stored in map
    //Label     ApiTestVariable = "var_label"

    //variables 
    // Epic        LabelType = "epic"
	// Layer       LabelType = "layer"
	// Feature     LabelType = "feature"
	// Story       LabelType = "story"
	// ID          LabelType = "as_id"
	// Severity    LabelType = "severity"
	// ParentSuite LabelType = "parentSuite"
	// Suite       LabelType = "suite"
	// SubSuite    LabelType = "subSuite"
	// Package     LabelType = "package"
	// Thread      LabelType = "thread"
	// Host        LabelType = "host"
	// Tag         LabelType = "tag"
	// Framework   LabelType = "framework"
	// Language    LabelType = "language"
	// Owner       LabelType = "owner"
	// Lead        LabelType = "lead"
	// AllureID    LabelType = "ALLURE_ID"

    
    
    
)

// Symbol represents a symbol in the symbol table.
type Symbol struct {
	Name          string
	Scope         SymbolScope
	Index         int
	LocalAssigned bool // if the local symbol is assigned at least once
}

// SymbolTable represents a symbol table.
type SymbolTable struct {
	parent         *SymbolTable
	block          bool
	store          map[string]*Symbol
	numDefinition  int
	maxDefinition  int
	freeSymbols    []*Symbol
	builtinSymbols []*Symbol
}

// NewSymbolTable creates a SymbolTable.
//TODO lifecycle of symboltable 
// new fork resolve addressing
func NewSymbolTable() *SymbolTable {
    table := &SymbolTable{
		store: make(map[string]*Symbol),
	}
    
    // table.Define(string(VarCase))
    //how to assign a var internally 
    // s := table.Define(string(VarCase))
    return table
}

// Define adds a new symbol in the current scope.
func (t *SymbolTable) Define(name string) *Symbol {
	symbol := &Symbol{Name: name, Index: t.nextIndex()}
	t.numDefinition++

	if t.Parent(true) == nil {
		symbol.Scope = ScopeGlobal

		// if symbol is defined in a block of global scope, symbol index must
		// be tracked at the root-level table instead.
		if p := t.parent; p != nil {
			for p.parent != nil {
				p = p.parent
			}
			t.numDefinition--
			p.numDefinition++
		}

	} else {
		symbol.Scope = ScopeLocal
	}
	t.store[name] = symbol
	t.updateMaxDefs(symbol.Index + 1)
	return symbol
}

// DefineBuiltin adds a symbol for builtin function.
func (t *SymbolTable) DefineBuiltin(index int, name string) *Symbol {
	if t.parent != nil {
		return t.parent.DefineBuiltin(index, name)
	}

	symbol := &Symbol{
		Name:  name,
		Index: index,
		Scope: ScopeBuiltin,
	}
	t.store[name] = symbol
	t.builtinSymbols = append(t.builtinSymbols, symbol)
	return symbol
}

// Resolve resolves a symbol with a given name.
func (t *SymbolTable) Resolve(
	name string,
	recur bool,
) (*Symbol, int, bool) {
	symbol, ok := t.store[name]
	if ok {
		// symbol can be used if
		if symbol.Scope != ScopeLocal || // it's not of local scope, OR,
			symbol.LocalAssigned || // it's assigned at least once, OR,
			recur { // it's defined in higher level
			return symbol, 0, true
		}
	}

	if t.parent == nil {
		return nil, 0, false
	}

	symbol, depth, ok := t.parent.Resolve(name, true)
	if !ok {
		return nil, 0, false
	}
	depth++

	// if symbol is defined in parent table and if it's not global/builtin
	// then it's free variable.
	if !t.block && depth > 0 &&
		symbol.Scope != ScopeGlobal &&
		symbol.Scope != ScopeBuiltin {
		return t.defineFree(symbol), depth, true
	}
	return symbol, depth, true
}

// Fork creates a new symbol table for a new scope.
func (t *SymbolTable) Fork(block bool) *SymbolTable {
    table:= NewSymbolTable()
    table.parent=t
    table.block=block
    return table

	// return &SymbolTable{
	// 	store:  make(map[string]*Symbol),
	// 	parent: t,
	// 	block:  block,
	// }
}

// Parent returns the outer scope of the current symbol table.
func (t *SymbolTable) Parent(skipBlock bool) *SymbolTable {
	if skipBlock && t.block {
		return t.parent.Parent(skipBlock)
	}
	return t.parent
}

// MaxSymbols returns the total number of symbols defined in the scope.
func (t *SymbolTable) MaxSymbols() int {
	return t.maxDefinition
}

// FreeSymbols returns free symbols for the scope.
func (t *SymbolTable) FreeSymbols() []*Symbol {
	return t.freeSymbols
}

// BuiltinSymbols returns builtin symbols for the scope.
func (t *SymbolTable) BuiltinSymbols() []*Symbol {
	if t.parent != nil {
		return t.parent.BuiltinSymbols()
	}
	return t.builtinSymbols
}

// Names returns the name of all the symbols.
func (t *SymbolTable) Names() []string {
	var names []string
	for name := range t.store {
		names = append(names, name)
	}
	return names
}

func (t *SymbolTable) nextIndex() int {
	if t.block {
		return t.parent.nextIndex() + t.numDefinition
	}
	return t.numDefinition
}

func (t *SymbolTable) updateMaxDefs(numDefs int) {
	if numDefs > t.maxDefinition {
		t.maxDefinition = numDefs
	}
	if t.block {
		t.parent.updateMaxDefs(numDefs)
	}
}

func (t *SymbolTable) defineFree(original *Symbol) *Symbol {
	// TODO: should we check duplicates?
	t.freeSymbols = append(t.freeSymbols, original)
	symbol := &Symbol{
		Name:  original.Name,
		Index: len(t.freeSymbols) - 1,
		Scope: ScopeFree,
	}
	t.store[original.Name] = symbol
	return symbol
}
