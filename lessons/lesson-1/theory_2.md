[//]: # (Заметки с лекции)
### Указатели:
- ZeroValue для pointer = nil
- Для избежания ошибок, рекоменуедтся использовать функцию new, которая выделяет память для int и возвращает указатель на нее

```
ptrNew := new(int)
*ptrNew = 100
```

### Приведение типов
Boolean можно обрабатывать через управляющие конструкци: 
```
var flag bool = true
var intFlag int = 0
if flag{
    intFlag = 1
}
```
Приведение интерфейсов:
```
var anyValue any

anyValue = 42
intValue, ok := anyValue.(int) // Приведение типа
if ok {
...} else {
...}

```