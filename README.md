
# Insights/ideas

```typescript

type Language = string
type URI = string
type Code = string // Code is basically Data here
type Grammar = Code
type Translation = Code
type System = {}

// TODO: Steal base types from LSP, use JSON-RPC

type Monotype = string // kind == 1
type TypeConstructor = string // kind > 1
type Type =
    Monotype | TypeConstructor

type LessThan = [string, string]
type Top = string
type Bottom = string
type Relation =
    LessThan |
    Top |
    Bottom

// Consider incremental approach (pipes)

interface Ontology {
    languages: Language[]
    links: { from: Language, to: Language, semantic: string }[]
    types: Type[]
    subtyping: Relation[]
    // grammars: (l: Language) => Grammar
    // translations: (l: Language) => Translation
}

namespace Frontend {
    namespace CodeExtractor {
        interface Query {
            lang: Language
            root: URI
        }

        interface Response {
            code: (Code | URI)[]
        }

        interface Extractor {
            extract(q: Query): Response;
        }
    }

    namespace SystemInfo {
        interface GenerateRequest {
            where: URI
        }

        interface InfoRequest {
            from: URI
        }

        type Query = GenerateRequest | InfoRequest

        interface Response {
            info: Code
        }

    /*class*/ interface SystemInfoProvider {
        /*private*/ generate(system: System): void
            info(q: Query): Response
        }
    }
}

// TODO: develop further
namespace Analyzer {
    type Constraint = {}

    interface Constraints {
        constraints: Constraint[]
    }

    interface Resolution {
        psi: any
        phi: any
    }

    namespace SyntaxTranslator {
        interface GenerateParser {
            lang: Language
            grammar: URI | Grammar
        }

        // translation should be in some sort of language, DSL?
        // L-attributed/S-attributed https://www.csd.uwo.ca/~mmorenom/CS447/Lectures/Translation.html/node4.html
        interface GenerateTranslator {
            lang: Language
            translation: URI | Translation
        }

        interface RunTranslation {
            lang: Language
        }

        type Query =
            GenerateParser |
            GenerateTranslator |
            RunTranslation

        type Response =
            Constraints

        interface Translator {
            constraints(q: Query): Response
        }
    }

    // NOTE: Maybe borrow some constraint solver?
    namespace Solver {
        interface SolveConstraints {
            constraints: Constraint[]
        }

        type Query = SolveConstraints

        type Response = Resolution

        interface Solver {
            solve(q: Query): Response
        }
    }
}

```

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
    - [x] Разобраться со строчкой 47 (doTwoPlusTwo), подумать над env: [3, 5, 8] (откуда там эта инфа?)
- [ ] Сделать обзорную статью методов разрешения проблем с областями (по сути обозреть 3 статьи по scope graphs)
    - [ ] Рассмотреть статьи отсюда https://research.tudelft.nl/en/persons/h-van-antwerpen

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
