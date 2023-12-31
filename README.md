
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

## Грамматика языка фрагментов анализатора
```ebnf
# Fragment
# Implies ({ start: file_line_col, end: file_line_col }, lang: <ID from ontology>)
F ::= ID

# Alias for a type
A ::= ID '=' T

# Type of a fragment
T ::= 
    | T '->' T                        # Function
    | T '*' T                         # Product
    | T '+' T                         # Sum
    | T '&' T                         # Intersection
    | T '|' T                         # Union
    | '{' (ID ':' T ',')* '}'       # Record
    | Unit | Any | Opaque           # Builtins

# Typed fragment (term)
t ::= '(' F ')' ':' T

# Collection of fragment terms
# In case of lhs of judgement ',' means conjunction 
# In case of rhs of judgement ',' means multiple conclusions
ts ::= t (',' t)*

# Judgement
# ';' here means disjunction 
J ::= ts (';' ts)* '|-' ts

# Block of judgements
# Judgements on the left of `with` use judgements on the right for inference, but only one level deep
# Judgements within the block are not structured lexically, hence
# conclusions below can be used to infer judgments above
B ::= (J '\n')+ ('with' B)?

Start ::= (A '\n')* 'then' B
```

## Определение функции связываемости фрагментов
Является функцией логического вывода
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
    linked \x.(f g) \y.(f g) == false
    linked 123 id == false
    linked \x.\y.z \x.(y z) == false
*)
let linked lhs: Term rhs: Term = (*lexical comparison for now*)
```

## Общее
Терм - фрагмент кода, уникально идентифицируемый в рамках всей системы, имеет тип

Тип - тип из системы типизации lambda2

Онтология включает:
- Правила взаимодействия языков (видимость одного языка из другого, название связи такой комбинации)
- Предметно-ориентированные типы (допустимые грамматикой выше)
- Ссылки на пару <Грамматика:Трансляция> для каждого вовлеченного языка

1. Универсальный анализатор типовых сигнатур - на вход получает все типовые сигнатуры сущностей в форме вложенных фрагментов и выявляет их связи между собой.
    - Связи используются для обнаружения зависимостей 
    - и поддержания корректности использования сущностей в различных контекстах
1. Такой анализатор использует набор фрагментов и онтологию предметной области, последняя получается из контекста проекта (файла конфигурации)
1. Набор фрагментов получается из результатов работы различных моноязыковых анализаторов, переводящих изначальную типовую информацию 
какого-либо языка в общую форму
1. Система типов универсального тайпчекера предположительно что-то вроде lambda 2
1. Фрагменты подчиняются определенному правилу линковки модулей

## Моноязыковые анализаторы
1. Я думаю стоит использовать [LSP](https://microsoft.github.io/language-server-protocol/)
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
- [x] Хороший первый прототип - анализатор который полностью корректно пробрасывает все запросы от ide в моно lsp, но не делает анализа (identity анализатор) 
- [x] Почитать [BSP](https://build-server-protocol.github.io/)
    - Фигня сырая...
- [x] Порисерчить парсинг для МЯ программ
    - [x] Дочитать диссер
- [x] Использовать ad hoc directed translation для построения фрагментов
- [x] Описать фрагменты и их семантику более формально
    - [~] Дочитать Карделли
- [x] Спросить про Fuzzy logic
    - Не первостепенно
- [x] Решить вопрос с грамматиками и syntax-directed translation [SDT](http://www.cse.iitm.ac.in/~krishna/cs3300/lecture4.pdf)
    - Решился через идею о протоколе
- [ ] Провести выборку проектов на гитхабе с несколькими языками
- [ ] Написать/адаптировать анализаторы с соответствующими языками
- [ ] Сделать прототип основного анализатора
- [ ] Сделать dot рисовалку графов
- [x] Доработать истоки-стоки - что является фактами, а что можно линковать?
    - [ ] Разобраться со строчкой 47 (doTwoPlusTwo), подумать над env: [3, 5, 8] (откуда там эта инфа?)

## Срочно неважно
  
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
