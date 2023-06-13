# Общие сведения

Обучающийся: Орловский Максим Юрьевич

Тема проекта: Выработка методов к анализу мультиязыковых текстов программ

Место выполнения проекта: Университет ИТМО

Куратор от университета: Логинов И.П. к.т.н.

Цель проекта: Выработка оптимальных методов анализа программ для построения эффективных инструментальных средств по рефакторингу и кодогенерации 


# Пожелания на проект
Набор оптимальных методов анализа мультиязыковых текстов программ, нацеленных на интеграцию в различные инструментальные средства программирования.

# Задачи на проект
## 1 семестр (сентябрь 2022 - январь 2023)
### 1 эпоха
1.	Написать техническое задание на проект.
2.	Согласовать техническое задание на проект.
3.	Выбор дисциплин для изучения.
### 2 эпоха
1.	Определить сценарии использования мультиязыковых анализаторов.
2.	Определить множество анализаторов исходных текстов программ для исследования и составить их краткие характеристики по критериям: функциональные возможности, место (этап) использования в жизненном цикле проекта.
3.	Исследовать возможности существующих средств анализа в части расширения наборов поддерживаемых сочетаний языков программирования.
4.	Определить показатели эффективности как для самих средств, так и для процесса расширения их функциональности.
### 3 эпоха
1.	Выбрать сочетание языков программирования и сценарии использования анализатора для пробной программной реализации. При выборе руководствоваться доступностью и функциональными возможностями программных интерфейсов существующих анализаторов для выбранных языков программирования с целью их совместного использования.
2.	Определить критерии для выполнения анализа мультиязыковых программ.
3.	Сформулировать требования к программной реализации прототипа анализатора.
### 4 эпоха
1.	Выполнить программную реализацию прототипа анализатора.
2.	Выполнить тестирование анализатора и анализ показателей эффективности, определить область применимости полученного решения.
3.	Представить отчет по проделанной работе.
## 2 семестр (февраль 2023 - июнь 2023)
### 1 эпоха
1. Исследовать возможность обобщения предлагаемого метода на сочетание различных языков программирования.
2. Исследовать наиболее часто используемые парадигмы программирования (процедурное, ОО, декларативное) в контексте мультиязыкового анализа.
3. Рассмотреть различные подходы к представлению семантической информации.
TODO:
1.  Use concepts of programming languages as a source for 2.
1.  Search sources for 3. (ПРИС не нашел, нужно искать аналогичные ресурсы, возможно на русском)
1.  1 is written easily i guess
### 2 эпоха
1. Провести анализ текущих решений в области мультиязыкового анализа на прикладном уровне (IDE и инструментальные средства).
2. Провести анализ моделей представления семантической информации программ в контексте различных ЯП и технологических стеков.
3. Исследование влияния прикладных областей языков на возможность обобщения метода мультиязыкового анализа.
TODO:
1.  Go to research of IDEs (Jetbrains, pvs studio, stuff...)
1.  Go and get their information models (if available)
1.  Go into categorization of various languages and try not to drown in them
### 3 эпоха
1. Сформулировать формально метод языкового анализа.
2. Рассмотреть природу входной метаинформации и способы её обработки и хранения.
3. Выбрать оптимальные структуры данных для представления метаинформации и хранения результатов анализа.
TODO:
1.  Formal definition of the analysis process (mathlike, idk)
1.  Software engineering for previous analyser
### 4 эпоха
1. Провести разработку прототипа для тестирования метода в различных сценариях.
2. Разработать избранную модель представления семантической информации для использования в прототипе.
3. Провести тестирование и собрать соответствующую статистику.
4. Исследовать ограничения предлагаемого метода.
5. Рассмотреть иные варианты получения информации для работы метода.
TODO:
1.  Do another prototype, consult language and implications of that (maybe go for abstract analyzer in something simple like Go?)
1.  Do prototype for at least three scenarious, all different kind
1.  Test, maybe benchmark stuff
1.  Write discussion basically (maybe not full)
## 3 семестр (сентябрь 2023 - январь 2024)
## 4 семестр (февраль 2024 - июнь 2024)

# Требования к проекту

## Требования к производительности

1. Разработать метод, обеспечивающий быстрый анализ в реальном времени (до 3 сек)
1. Разработку методов вести с расчетом на большие объемы анализируемых данных (до ? МБ)

# Пользовательские сценарии

## Использование методов анализа в IDE

1. Подсветка синтаксических конструкций 
1. Навигация по различным синтаксическим элементам
1. Обеспечение связи идентификаторов из разных языковых фрагментов

## Использование методов анализа при разработке инструментальных средств

1. Автоматический рефакторинг смешанного кода
1. Сбор статистических данных о смешанном коде

# Тестирование и проверка 

(Использование тестовых фрагментов смешанного кода)

# Планируемые курсы к изучению
[C#](https://www.udemy.com/course/c-sharp-oop-ultimate-guide-project-master-class/)

[Compilers](https://www.edx.org/course/compilers)

[NLP in Python](https://www.udemy.com/course/natural-language-processing-in-python)

[Haskell](https://stepik.org/course/75)