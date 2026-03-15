# Метки (labels)

В языке Go есть метки (labels), которые позволяют перемещаться к разным частям кода. 

Метку можно указать для операторов:

- break;
- continue;
- goto (безусловный оператор перехода, позволяет перейти в любое место кода).

Приведём пример для break:

```go
outerLoopLabel:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Printf("[%d, %d]\n", i, j)
            break outerLoopLabel
        }
    }
    fmt.Println("End")
```

Здесь `break outerLoopLabel` прерывает выполнение внешнего цикла.
А вот пример для `continue`:

```go
outerLoopLabel:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Printf("[%d, %d]\n", i, j)
            continue outerLoopLabel
        }
    }
    fmt.Println("End")
```