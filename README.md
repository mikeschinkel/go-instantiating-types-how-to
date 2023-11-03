# GoLang — Insantiating Types — Options Structs vs Option Funcs


Oo the [r/golang](https://www.reddit.com/r/golang/) forum on Reddit user [u/Nottymak88](https://www.reddit.com/user/Nottymak88/) posted [his code](https://go.dev/play/p/2T4St61sN5Y) and asked for help [_"Assigning values into Nested structs"_](https://www.reddit.com/r/golang/comments/17md60x/assigning_values_into_nested_structs/).

That is such a common question for those learning Go I decided to answer it with the repository giving two (2) different approaches.

But first, let's start with some general tips based on their code.


## General Recommendations
Here are the things that came to mind while reading the asker's code. They are not arranged in perfect order, but at least I tried. A bit.

1. **Don't Try to Eat the Entire Apple** — Instead of trying to assemble a large struct with many embedded structs, _treat each struct as its own entity_, complete with methods to enable interaction. This approach allows for divide and conquer and reduces the complexity one needs to manage in ones head all at once.
2. **Use Instantiators** — With Go we can instantiate an object using itsname and braces — e.g. `Policy{}` — but experience tells me you are almost always better off defining at least one instantiator function, typically prefixed with `New`, e.g. `NewPolicy()`.   
3. **Use Pointers** — Usually it is better to use pointers — such as when returning from `New*()` intantiator funcs and when creating slice types like `[]*Coverage` vs `[]Coverge`. Not using pointer ends up creating many small frictions, and it is easier to just use pointers. One example is recursively defined structs must use pointers, and another example is an [interface cannot call a func with a pointer receiver if the object is not a pointer](https://goplay.tools/snippet/AGB7Gv8iZC6). **_However_**, if you do have a valid reason for _not_ using a pointer — _such as minimizing heap allocations and garbage collection for edge-case apps_ or you have the discipline to code Go in a purely functional, non-mutable approach, which by the way is not idiomatic Go — then _by all means return a stack-based value_.
4. **It's Good to String Along Your Types** — If you give each of your types a `String()` property then passing an instance of the type to `fmt.Println()` or any of the similar print/sprint methods will automatically call the type's  `String()` method. This simplifies composing log messages, error messages and other human-readable output. Actually, this is just one use-case of the `Stringer` interface being satisfied by having a `String()` method, but [interfaces](https://go.dev/tour/methods/9) are way [out of scope](https://golangbot.com/interfaces-part-1/) for this discussion.
5. **Give `main()` a package of its own** — Create a package of its own for `main()` and put in the cmd directory, [as the Go team recommends](https://go.dev/doc/modules/layout). This means your `main()` code will need to use package references to call your other code. This can help leads to developing applications as a thin executable veneer around everything else that are reusable packages, which IMO is a best practice.
6. **Be pithy when naming your reusable package** — Packages names need to complete globally _(within a Go application)_ so name your package something short, pithy, and unlikely to be used by another package and/or a variable name. As an aside, I curse the Go team for naming the URL package `net/url` thereby squatting on the perfect variable name for a URL, e.g. `url` and a name that every project I have every worked on it filled with variables of that name thereby making an IDE constantly complain when I read their code. But I digress.
7. **Give Each Type a File** — Although not literally _**every**_ time, but it makes good sense to separate out your code into multiple files so it is easier to wrap your head around what code is where, and also it gives you room to elaborate on each type as you will almost certainly do so when you are writing a real-world application.
8. **Give Each Property a Line** — Don't list multiple properties on the same line in a struct. Sure it may eliminate duplication of the type but doing so can make a struct exceedingly hard to read, especially for properties and types that end up with 20, 30, 40, 50 or more characters of whitespace between the name and the type. Also, you can't add property tags to properties — such as for JSON — when multiple properties are combined on a line.
9. **Create Plural Types** — When you want to use a slice of a struct, such as `[]Location` — or as I suggested `[]*Location` – then go ahead and create a plural type, such as `Locations`. Doing so will allow you to create methods for that type when you realize that's the best way to minimize unnecessary code duplication — [vs. acceptable code duplication](https://research.swtch.com/deps#avoid_the_dependency) — and, you will thank me later.
10. **Use Adder Functions** — When you have a struct with a property whose type is a slice of objects, write a method to add those objects to that object. It makes the code much more readable, encapsulates the code that does the adding which ensures it gets called correctly and added correctly, and reduces more unnecessary code duplication.   
11. **Default to Private** — With Go, struct properties are package-private if they start with lower-case. As a rule of thumb it is better to start with all properties being private and then either expose the ones that needs to be public on an as-needed basis, or better, create methods to access those properties. The latter approach is often preferred in Go because methods can participate in interfaces but properties cannot, and the more complex a project becomes the more likely it will need to use interfaces to resolve cyclical dependencies, among several other benefits of using interfaces.  
12. **Avoid _"Or"_ in type names** — This may just be an opinionated persoanl best practice, but using `PersonOrOrg` when you could use `Insured` seems like it is asking for complex naming in methods that will be needed.  

## Approach A — Using  _"Options"_ structs

[Approach A](tree/main/approach_a) is generally my preferred approach and it uses a bespoke **options struct** for each entity struct's optional properties.   

Here is what that looks like for `Policy`, which I derived from [the asker's original code](https://go.dev/play/p/2T4St61sN5Y): 

```go
type Policy struct {
  number         string
  effectiveDate  time.Time
  expirationDate time.Time
  lines          Lines
  transactions   Transactions
}

type PolicyOpts struct {
  EffectiveDate  time.Time
  ExpirationDate time.Time
}

func NewPolicy(number string, opts *PolicyOpts) *Policy {
  return &Policy{
    number:         number,
    effectiveDate:  opts.EffectiveDate,
    expirationDate: opts.ExpirationDate,
    lines:          make([]*Line, 0),
    transactions:   make([]*Transaction, 0),
  }
}

func (p *Policy) AddLine(line *Line) *Policy {
  p.lines = append(p.lines, line)
  return p
}

func (p *Policy) AddTransaction(tx *Transaction) *Policy {
  p.transactions = append(p.transactions, tx)
  return p
}
```

The above can then be used like so:
```go
func main() {
  policy := NewPolicy("Policy1", &PolicyOpts{
    EffectiveDate:  now,
    ExpirationDate: addYear(now),
})
fmt.Printf("%s\n", policy)
```

However, here is the example the asker wanted to encode into embedded structures, which I embellished a bit by adding a few more values to his instantiation request:

```go 
package main

import (
  "fmt"
  "time"

  "github.com/shopspring/decimal"
  "insure"
)

func main() {
  now := time.Now()

  /*
    var P Policy
    P.Number = "Policy1"
    P.line[0].ID = "Line1"
    P.line[1].ID = "Line2"
    P.transaction[0].ID = "Transaction1"
    P.line[0].coverages[0].Indicator = true
    P.line[0].Risks[0].ID = "Risk1"
    P.line[0].Risks[1].ID = "Risk1"
    P.line[0].Risks[0].coverages[0].Indicator = true
    P.line[0].Risks[1].coverages[0].Indicator = true
    P.line[0].loc[0].Address1 = "Addr1"
  */
  policy := insure.NewPolicy("Policy1", &insure.PolicyOpts{
    EffectiveDate:  now,
    ExpirationDate: addYear(now),
  }).
    AddTransaction(insure.NewTransaction("Transaction1")).
    AddLine(insure.NewLine("Line1", &insure.LineOpts{
      TypeLOB:          insure.AutoLOBType,
      TermPremium:      decimal.NewFromInt(120),
      PriorTermPremium: decimal.NewFromInt(110),
    }).
      AddCoverage(insure.NewCoverage(true)).
      AddRisk(insure.NewRisk("Risk1", &insure.RiskOpts{
        Included: true,
      }).AddCoverage(insure.NewCoverage(true))).
      AddRisk(insure.NewRisk("Risk2", &insure.RiskOpts{
        Included: true,
      }).AddCoverage(insure.NewCoverage(true))).
      AddLocation(insure.NewLocation("Addr1")),
    ).
    AddLine(insure.NewLine("Line2", &insure.LineOpts{}))

  fmt.Printf("%s\n", policy)

}

func addYear(t time.Time) time.Time {
  return t.AddDate(1, 0, 0)
}

// shortFormTime is a format.
// To understand why 2006-01-02, see https://stackoverflow.com/a/52966197/102699
const shortFormTime = "2006-01-02"
```

You can see the [**complete code for Approach A here**](tree/main/approach_a).

Note, the complete code **is not a fully fleshed out app**; it just illustrates the points mentioned here but does not try to go farther.

## Approach B — Using _"Options"_ funcs

[Approach B](tree/main/approach_a) is a variant of  an approach that AFAIK was [first proposed by Dave Cheney](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) which he called _"Functional options for Friendly APIs."_ Dave Calhoun [has covered it](https://www.calhoun.io/using-functional-options-instead-of-method-chaining-in-go/) as well as many others since Dave, too. The proposed v2 of the standard `encoding/json` package [also chose this approach](https://github.com/golang/go/discussions/63397). 

Here is our equivalent to Approach A for `Policy` when using option funcs aka Approach B. Note how the `Set*()` and `Add*()` methods return a closure that gets executed in `NewPolicy()`: 

```go
type Policy struct {
  number         string //GUID
  effectiveDate  time.Time
  expirationDate time.Time
  lines          Lines
  transactions   Transactions
}

type PolicyOptions func(*Policy)

func NewPolicy(number string, opts ...PolicyOptions) *Policy {
  p := &Policy{
    number:       number,
    lines:        make([]*Line, 0),
    transactions: make([]*Transaction, 0),
  }
  for _, opt := range opts {
    opt(p)
  }
  return p
}

func (PolicyOptions) SetEffectiveDate(d time.Time) PolicyOptions {
  return func(p *Policy) {
    p.effectiveDate = d
  }
}

func (PolicyOptions) SetExpirationDate(d time.Time) PolicyOptions {
  return func(p *Policy) {
    p.expirationDate = d
  }
}

func (PolicyOptions) AddLine(line *Line) PolicyOptions {
  return func(p *Policy) {
    p.lines = append(p.lines, line)
  }
}

func (PolicyOptions) AddTransaction(tx *Transaction) PolicyOptions {
  return func(p *Policy) {
    p.transactions = append(p.transactions, tx)
  }
}
```

You can use the above in an example equivalent to the one shown for Approach A, like so:

```go
func main{
  var po insure.PolicyOptions
  policy := insure.NewPolicy("Policy1",
    po.SetEffectiveDate(now),
    po.SetExpirationDate(addYear(now)),
  )
  fmt.Printf("%s\n", policy)
}
```

Then we have the full the example the asker wanted to encode into embedded structures, which I also embellished the same amount here:

```go 
package main

import (
  "fmt"
  "time"

  "github.com/shopspring/decimal"
  "insure"
)

// shortFormTime is a format.
// To understand why 2006-01-02, see https://stackoverflow.com/a/52966197/102699
const shortFormTime = "2006-01-02"

func main() {
  now := time.Now()
  var po insure.PolicyOptions
  var lo insure.LineOptions
  var ro insure.RiskOptions

  /*
    var P Policy
    P.Number = "Policy1"
    P.line[0].ID = "Line1"
    P.line[1].ID = "Line2"
    P.transaction[0].ID = "Transaction1"
    P.line[0].coverages[0].Indicator = true
    P.line[0].Risks[0].ID = "Risk1"
    P.line[0].Risks[1].ID = "Risk1"
    P.line[0].Risks[0].coverages[0].Indicator = true
    P.line[0].Risks[1].coverages[0].Indicator = true
    P.line[0].loc[0].Address1 = "Addr1"
  */
  policy := insure.NewPolicy("Policy1",
    po.SetEffectiveDate(now),
    po.SetExpirationDate(addYear(now)),
    po.AddTransaction(insure.NewTransaction("Transaction1")),
    po.AddLine(
      insure.NewLine("Line1",
        lo.SetTypeLOB(insure.AutoLOBType),
        lo.SetTermPremium(decimal.NewFromInt(120)),
        lo.SetPriorTermPremium(decimal.NewFromInt(110)),
        lo.AddCoverage(insure.NewCoverage(true)),
        lo.AddRisk(insure.NewRisk("Risk1",
          ro.AddCoverage(insure.NewCoverage(true)),
        )),
        lo.AddRisk(insure.NewRisk("Risk2",
          ro.AddCoverage(insure.NewCoverage(true)),
        )),
        lo.AddLocation(insure.NewLocation("Addr1")),
      ),
    ),
    po.AddLine(insure.NewLine("Line2")),
  )

  fmt.Printf("%s\n", policy)

}

func addYear(t time.Time) time.Time {
  return t.AddDate(1, 0, 0)
}
```

You can see the [**complete code for Approach B here**](tree/main/approach_b).


## The Output of Both Approaches
So you can visualize the output of the two different approaches, here is what you'll see from both them:

```
Policy:
  Number: Policy1
  Effective Date: 2023-11-03
  Expiration Date: 2024-11-03
  Lines:
    Line: Line1
      Coverages:
        Coverage: true
      Risks:
        Risk ID: Risk1
          Coverages:
            Coverage: true
        Risk ID: Risk2
          Coverages:
            Coverage: true
      Locations:
        Location: Addr1
    Line: Line2
      Coverages:
      Risks:
      Locations:
  Transactions:
    Transaction ID: Transaction1
```

**Note:** the output does not show all values our code sets because I did not go back any update my `String()` after I embellished the initialization code with a little extra data. 

## Summary

There are many differnt ways to skin a cat — _no offence to felines, or to those who adore them_ — and these are just two approaches to object creation in Go. 

Ask any other Go programmer, and they are certain to have either small tweaks to what I have shown all the way up to major disagreements. 🤷‍ 

But the reality is these approaches work and while no approach is perfect these two approaches can provide developers new to Go with a starting point for learning how to create embedded structures in Go and learning how to overall improve their craft.  

Besides, aren't programmer disagreements what makes the world go round?  🙂     