# Tengo QA

This is a fork of the script language [tengo](https://github.com/d5/tengo)
specialized for api test automation.

```golang
/* The Tengo Language */
fmt := import("fmt")


Domain("https://localhost:8080")
//declare a test case
Case("name of case 1", "fullname of case 1")

Parameter("author", "bob")

resp := Request("/get", "GET")
fmt.println(resp)
//jsonpath
names:= Extract("$[*].name", resp.body)
AssertThat(len(names)>0)
```

## Features

- Features from tengo
- Builtin functions for test(Domain, Feauture, Case, Step, ...)
- Variable inheritance & error propagation
- Test report generation

## Installation
```bash
git clone https://github.com/BHansom/tengo
cd tengo
make cli
```

## Guide

To make it powerful for api-test, some functions are added to tengo.

### builtin-functions

#### `Case()`, `Step()`

`Case()` initialize a **allure** test case, where subsequent `Step()` or Tag function calls will be contained in the final test report.

`Step()` initialize a **allure** test case step. Steps can be recusive, if there steps that are not closed when calling the function.

#### Tag functions: 

1.  Header()
2.  Domain()
3.  ParentSuite()
4.  Suite()
5.  SubSuite()
6.  Epic()
7.  Feature()
8.  Story()
9.  Package()
10. Attachment()
11. Parameter()

`Header()` and `Domain()` are functions that set `request()` related data to be used subsequently.

The rest functions listed above are tag functions which set the test cases' tags, which will be used in the allure report.

Note: `Attachment()` and `Parameter()` can acts on step or case, depending on the context.


#### `Pass()/Fail()/PassStep()/FailStep()`

These are functions manully closing the case/step.

Implicitly, unclosed case will be close either on the next call to `Case()` or at the end of the function block(condition/loop/function).

Steps has to be closed explicitly, or closed as the effect of case close (`Step()` will act on unclosed step, as mentioned above).

#### `Request()`

`Request()` do a http request, and returns a response object. This function is a `Step`, thus it will not work when no case/step available.

#### `Extract()`

`Extract()` makes a jsonpath operation and returns a array as result. This function is a `Step`.

#### `AssertThat()/AssertEqual()/AssertNotEqual()`

Assertions makes check of the parameters. Assertions are steps.

### variable inheritance

Builtin variables are used when interact with cases/steps. The case/step/tags will be propagated as implicit args in func calls.

E.g., you can declare tags before and use them in all the cascaded scopes without redeclaration(tags declared globally in ohter modules are prior to params
and will thus override the declared tags). The builtin variables cannot be upward propagated.


### error-propagation

Errors will cause the following case to be skipped in the same function.

Note: Errors here are not grammaric/semantic/runtime errors that cause the program halting but the errors from `Fail()/FailStep()` or the builtin steps.
To bubble up errors, a `!` should be specified when calling the function: `a:= loginModule.login()!`

### allure report
If `Case()` are called at least once during the execution, a `allure-results` directory will be created and allure cases results will be outputed. All the
data gathered during the execution will be shown in the workflow of allure.

For more details, see [allure](https://allurereport.org/) and [allure-go](https://github.com/ozontech/allure-go)


## See Also

- The Tengo Lang: [https://github.com/d5/tengo](https://github.com/d5/tengo)
- The Test Report framework: [allure](https://github.com/allure-framework/allure2)

