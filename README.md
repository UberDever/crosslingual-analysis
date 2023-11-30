
# Insights/ideas

## Универсальный анализатор (прототип реализован как prototype2)

```haskell
    type LSPMessage = -- <message from lsp basically json>
    type Fragment = Text
    type Tree = -- <tree of text>
    type Parser = Fragment -> EBNF -> Tree -- basically interpreter
    type Judgement = -- <AST of lambda 2>
    type Module = Graph [Judgement]
    type Relation = (Judgement, Judgement)

    translate :: Parser -> Productions -> Fragment -> Module
    analyze :: [Module] -> [Relation]

    parseLSPMessage :: LSPMessage -> Fragment
    dumpCache :: Module -> IO Filepath -- and also readCache
```


```bnf
Fragment ::= '[' ((Language|'_' ':' Isolation|'_') | '_')? ']' Source ':' Signature ('<=' Environment)?

# assert signature.isolation >= fragment.isolation
# assert signature.language == fragment.language
# isolation/languages do make sense for environments, but not so much
Language ::= ID
Isolation ::= NUMBER

# TextSpan: { start: file_line_col, end: file_line_col }
# Нет смысла использовать TextSpan по тому же принципу, что и внизу
Source ::= ID

Signature ::= Type

Environment ::= Term

Term ::= 
    | Fragment
    # Нет смысла описывать такую грамматику
    # Она позволяет либо сделать модули первоклассными - что хорошо для #define
    # но усложняет и без того сложное
    # Либо она позволяет выражать "изоморфизм" связанных фрагментов
    # но фрагменты разных языков __всегда__ связаны именовано, либо структурно/по месту, третьего не дано
    # | ID
    # | \ID.Term
    # | (Term Term) 

Type ::=
    | Type -> Type # Function
    | '(' (Term,)* ')' # Tuple
    | Primitive

Primitive ::=
    | Unit
    | Int
    | Opaque
    # Пока что всё...
```

```ocaml

let linked lhs: Fragment rhs: Fragment =
    (linked lhs.term rhs.term) `and` (linked lhs.type rhs.type)

(*
    linked Bool Bool == true
    linked Int Unit == false
    linked (Bool -> Int) (Bool -> Int) == true
    linked (Int -> Int) (Unit -> Int) == false
    linked _ Bot == false; linked Bot _ == false
    linked _ Any == true; linked Any _ == true
*)
let linked lhs: Type rhs: Type = (*structural comparison for now*)

(*
    linked twoPlusTwo twoPlusTwo == true
    linked \x.(f g) \y.(f g) == true
    linked 123 id == false
    linked \x.\y.z \x.(y z) == false
*)
let linked lhs: Term rhs: Term = (*structural comparison for now*)
```


Терм - фрагмент кода, уникально идентифицируемый в рамках всей системы, имеет тип

Тип - тип из системы типизации lambda2

Архетип - common known структура термов (1) и её тип (2), которые отражают дополнительную информацию о предметной области

Онтология включает:
- Правила связи между языками (видимость, наличие связи) и семантику (название) этой связи
- Наверняка правила связи внутри языка (видимость)
- Архетипы (возможно однопараметрические конструкторы)

1. Универсальный анализатор типовых сигнатур - на вход получает все типовые сигнатуры сущностей в форме вложенных фрагментов и выявляет их связи между собой.
    - Связи используются для обнаружения зависимостей 
    - и поддержания корректности использования сущностей в различных контекстах
1. Такой анализатор использует набор фрагментов и онтологию предметной области, последняя получается из контекста проекта (файла конфигурации)
1. Набор фрагментов получается из результатов работы различных моноязыковых анализаторов, переводящих изначальную типовую информацию 
какого-либо языка в общую форму
1. Система типов универсального тайпчекера предположительно что-то вроде lambda 2
1. Фрагменты подчиняются правилу линковки модулей, предположительно из пейпера Карделли

## Моноязыковые анализаторы
1. Я думаю стоит использовать [LSP](https://microsoft.github.io/language-server-protocol/) и собирать типовые сигнатуры через него
    * LSP Дает языко-независимую информацию. Можно считать, что из него можно вытянуть максимум что-то вроде `fn some_func(a: float, c: bool) -> Arc<unit>`
    * Следовательно, нужен микро-транслятор для каждого языка, переводящий такие сигнатуры в структурированные фрагменты
    * Т.к. трансляция чисто индуктивная, можно использовать Postorder и получится ad hoc syntax directed translation
1. Типовые сигнатуры собираются только для того что можно в каком-то отношении считать `public`
1. Каждый терм имеет окружение и сигнатуру состоящую из окружений и сигнатур своих поддтеров + свою собственную
1. Деление на фрагменты случайное, но хорошо если оно напоминает действительную иерархию сущностей в проекте
1. Сохранение scoping rules может быть произведено через вложенность фрагментов. First-class фрагменты вещь путанная и через чур сложная для поставленных задач
1. Термы имеют области видимости (наверняка характеризуемые числом) которые позволяют контроллировать name-aliasing

# TODO

## Несрочно неважно

- [x] Подумать, нужен ли анализ неисполняемых сущностей, в первую очередь определений типов данных для межфрагментного анализа
    * Не нужен

## Несрочно важно

- [ ] Изменить описание про MLSA в итоговой ВКР
- [ ] Почитать https://arxiv.org/abs/1808.01210
- [x] Расписать про дихотомию исполняется/неисполняется, использовать статью <?>
- [ ] Расписать про ценности анализа
    - [ ] Анализ на основе формальной модели, завязанной на фрагментах (модулях) и их линковке
    - [ ] Механизм дальнейшей эксплуатации информации о связях
- [x] Почитать про [LSP](https://microsoft.github.io/language-server-protocol/)
- [x] Возможно, инференс эффектов и парсер можно сделать тупо через конфигурацию? No
- [x] Проработать fallback анализ (для кейсов с например встроенными языками)
- [ ] Хороший первый прототип - анализатор который полностью корректно пробрасывает все запросы от ide в моно lsp, но не делает анализа (identity анализатор) 
- [x] Почитать [BSP](https://build-server-protocol.github.io/)
    - Фигня сырая...
- [ ] Порисерчить парсинг для МЯ программ
    - [x] Дочитать диссер
- [ ] Использовать ad hoc directed translation для построения фрагментов
- [ ] Описать фрагменты и их семантику более формально
    - [ ] Дочитать Карделли

## Срочно неважно

Type - как (и следует ли) прикрутить другую логику не сломав всё?
- Что насчет fuzzy logic?
- Надо всё переложить в логику... так будет проще оценить достоинства и недостатки подхода

Syntax directed translation
- Грамматики нужно где-то брать, в каком-то формате...
- На каком языке приделывать translation part? [SDT](http://www.cse.iitm.ac.in/~krishna/cs3300/lecture4.pdf)

Infrastructure
    

- [x] Сделать разноплановые тесты на прототип 2
- [x] Обзор текущих конференций (напр КМУ)
- [x] Ису достижения
- [x] Манир по результатам прошлого года, ближайшая конференция - майоровские чтения

## Срочно важно

- [x] Отчет по 2 семестру
- [x] Отчет: Do another prototype, consult language and implications of that (maybe go for abstract analyzer in something simple like Go?)
- [x] Отчет: Do prototype for at least three scenarious, all different kind
- [x] Отчет: Test, maybe benchmark stuff
- [x] Отчет: Write discussion basically (maybe not full)
    - [x] Отписать про внутримодульный и межмодульный анализ
- [x] Прототип 2, гетерогенный лист
- [x] Прототип 2, семантическая сеть и анализатор
- [x] Прототип 2, юзкейс 1
    - [x] анализ
- [x] Прототип 2, юзкейс 2
    - [x] анализ
- [x] Прототип 2, юзкейс 3
    - [x] анализ
- [x] Рассмотреть идею взаимодействия языков в юзкейсах как модулей
- [x] Ввести онтологию для такого универсального межмодульного взаимодействия
- [x] Презентация


# Examples
### 1. JS and C
```cpp
#include "some file"

int foo() {
    a += 2;
}
```
```js
a = require("some module")

export function bar() {
    a.foo()
}
```
```ocaml
(* Note: file.js has interlinked dependencies, 
but they are not listed here because this is not a tool concern for now*)

[C:0] file.c 
:   _ foo: Unit -> Int <= _ a: Int
<= ( _ a: Int, _ some file: File )

[Js:0] file.js
:   _ bar: Unit -> Int
    <= _ foo: Unit -> Any 
        <= _ a: Opaque
<= _ some module: File
```
### 2. TS
```ts
function foo(): number {
    let bar = () => 21 
    let n = bar()
    return n * 2
}
```
```ocaml
[Ts:0] file.ts:
<= TODO
```

### 3. Js
```ts
export default function (a: int) => (b: int) => plus(a, b) // assuming plus is declared in stdlib and opaque/implicit
```
```js
f = require("file.ts")
let a = f(1, 2)
```
```ocaml
(*
This example proves that all entities, that interact in code are always either:
 - named
 - nested/arranged in some pattern 
 *)

[Ts:0] file.ts: Int -> Int -> Any

[Js:0] file.js: () <= [Ts:0] file.ts: File 
```