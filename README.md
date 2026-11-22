# Homework 5 - Functional Programming

This submission presents two solutions for the "What is left?" problem on Eolymp: https://eolymp.com/problems/1964.

One of the solution was written in a functional programming language - Haskell. The other one was written in an imperative programming language - Go.

## Haskell

The Haskell solution uses recursion, pattern matching, and the `zip` function to solve the problem. This solution does not declare any variables. All lists are used only once and recreated whenever a change is needed, as is expected in a functional programming language. It scores 97% on Eolymp and fails on the remaining 3% due to the time limit.

To test the score of this code, select `Haskell (ghc 8.8)` in the programming language selection menu on the problem's page on Eolymp (the link is provided above). After that, copy-paste the solution from `haskell/Main.hs` into the editor and click "Submit".

If you want to run this code in an interactive mode, you can use this OnlineGDB link: https://onlinegdb.com/9RvOvbDH0

If you wish to run this code locally, install GHC ([instructions here](https://www.haskell.org/downloads/)) and compile the Haskell code using the following command in your terminal:

```bash
ghc haskell/Main.hs
```

After this, run the compiled program:

```bash
haskell/Main
```

Input the desired number and you will see the printed result shortly.

## Go

The Go solution, on the other hand, declares an array and some variables, and then manipulates the array using these variables in a loop. All changes are done in-place. This solution scores 100% on Eolymp. Even though this is not an optimal solution, it still gets this score for some reason - probably Go is fast enough to pull this off.

To test the score of this code, select `Go 1.24` in the programming language selection menu on the problem's page on Eolymp (the link is provided above). After that, copy-paste the solution from `go/main.go` into the editor and click "Submit".

If you want to run this code in an interactive mode, you can use this OnlineGDB link: https://onlinegdb.com/DWxqNayxL

If you wish to run this code locally, install GHC ([instructions here](https://go.dev/doc/install)) and run the Go program using the following command in your terminal:

```bash
go run go/main.go
```

Input the desired number and you will see the printed result shortly.
