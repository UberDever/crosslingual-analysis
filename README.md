
# Insights/ideas



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
- [ ] Провести выборку проектов на гитхабе с несколькими языками (Usecases!!!)
- [ ] Написать/адаптировать анализаторы с соответствующими языками
- [x] Сделать прототип основного анализатора
- [x] Сделать dot рисовалку графов
- [x] Доработать истоки-стоки - что является фактами, а что можно линковать?
    - [x] Разобраться со строчкой 47 (doTwoPlusTwo), подумать над env: [3, 5, 8] (откуда там эта инфа?)
- [x] Сделать обзорную статью методов разрешения проблем с областями (по сути обозреть 3 статьи по scope graphs)
    - [x] Рассмотреть статьи отсюда https://research.tudelft.nl/en/persons/h-van-antwerpen
- [ ] Использовать/изучить spoofax workbench для построения анализаторов
- [x] Подумать над использованием SSA вместо AST для повышения полноты анализа
    - потом, возможно никогда...
- [x] Рассмотреть расширение скопграфов (или это уже к ним не относится?) в отношении:
    - [x] Семантики user/used для идентификаторов (it references that)
        * Учитывается в семантике лейблов у ребер графа. Взято из третей по счету статьи
    - [x] Семантика языковых барьеров (окраски ребер графа). Возможно, достаточно последующей интерпретации и отметания связей
        * То же, что и выше. Правда, возможно необходимо будет немного повозится
- [x] Расписать evaluation/usecases
    - [x] Источники данных
    - [x] Конфигурация окружения (и нужные данные из него)


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
- [ ] составить ТЗ
- [ ] Сделать начальные трансляторы (отображения text -> AST + траверс) для каждого из включенных языков
    * [X] Golang
    * [x] JSON
    * [x] HTML
    * [x] XML
    * [x] Python
    * [ ] Makefile/shell?
    * [ ] Shell
    * [ ] C#
    * [ ] VB
    * [ ] JS
    * [ ] Fortran
    * [ ] C++
- [ ] Сделать арбитр/медиатор/фасад/TranslationBroker
    * Штука, которая берет на себя все инфраструктурные издержки по приему/передаче запросов между трансляторами.
    Нужен, так как любому парсеру i может понадобится послать часть разбираемого текста парсеру j.
    Также, такой арбитр может быть ответственен за распознавание текста для получения языка. Пока этот момент
    можно не прорабатывать, а тупо использовать на незнакомой строке все виды трансляторов какие есть. Хотя, концептуально
    стоит такой модуль сделать внешним, а у арбитра лишь оставить возможность обращаться к такому модулю
- [ ] Реструктурировать протокол
- [ ] Переделать всё под JSON
    - [ ] toDot
    - [ ] трансляторы