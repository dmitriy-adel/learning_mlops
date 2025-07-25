# Теоря по Go

- Здесь будет собрана теория в текстовом виде ради увеличения скорости ее запоминания и возможность подглядеть при необходимости

## Самая база

1. Go является языком со строгой типизацией. Есть вариант объявлять и инициализировать переменные в обычном формате, однако при этом есть возможность пользоваться утиной типизацией. Также наименования в Go записываются в camelCase: 

    ```go
    var characterName string = "Darth Vader"  // строгая типизация
    characterName := "Darth Vader" // утиная типизация
    ```

    При этом нельзя менять тип хранимых данных в переменной после ее объявления. Условно, если в переменной хранится интовое значение, то в переменную в последствии можно записывать только инты. А вот другие типы данных не получится


2. Go является компилируемым языком (код, написанный на Go сразу читается как машинный код (бинарка) без предварительного перевода, как в Python). Например, прогон цикла for от 1 до 100млн. с суммированием каждого значения займет на питоне 6 секунд, в то время как гошка справится за 50 мс, что в 120 раз быстрее. Также сама компиляция кода на Go занимает очень мало времени. А еще в Go есть нормальный параллелизм

--- 

## Типы данных

- В принципе, здесь тоже все стандартно за исключением `float` - при использовании плавающей точки требуется указывать кол-во бит на переменную (32 или 64), хотя тот же int можно определить как простой int без указывания его размерности, имхо она выберется автоматом

- Для каста совместимых типов данных (int во float, например), требуется просто обернуть переменную для каста в требуемых тип данных
    ```go
    var intNum32 int32 = 23
    var result float64 = float64(intNUm32)
    ```

- При этом есть возможность создавать константы - переменные, которые нельзя изменить. При их объявлении требуется указать `const` вместо `var`. При попытке изменить константу интерпретатор выдаст ошибку

- Для создания многострочной строки (тройные кавычки в питоне) стоит использовать '`':
    ```go
    var text string = `oh
    my
    god`
    ```

- Также в Go есть интересный тип данны - `rune`(руны). Это символы из ASCII кодировки, которые имеют слишком большие индексы, из-за чего кодируются несколькими байтами. Работа с ними может несколько ломать метод **len()**, поскольку он считает именно что кол-во байт, которые труебются на хранение переменной. Для перевода рун в чары используется метод `utf8.RuneCountInString(rune)` из пакета `unicode/utf8`

- Еще существует такое понятие, как 'значение по умолчанию'. Если просто объявить переменную и не инициализировать ее, переменная получит значение по умолчанию для ее типа данных. Для всех видов `int`, `float`, `rune` это значение равно 0, для `string` - '', для `bool` - False

---

## Функции и структуры данных

---

## Циклы

---

## Терминал

- Если посмотреть на вывод **go help**, можно увидеть не очень малое множество допустимых команд (у каждой еще свои подкоманды), но основные вот: 

    1. `go run`. Компиляция и запуск go-скриптов
    2. `go build`. Компиляция go-скриптов в исполняемый файл без последующего запуска скомпилированного файла
    3. `go install`. Установка пакетов в **$GOBIN**
    4. `go test`. Запуск тестов в указанном пакете. Используеся для модульного тестирования
    5. `go mod`. Управление зависимостями (**init**, **tidy**)
    6. `go fmt`. Автоматическое форматирование кода
    7. `go vet`. Проверка кода на типичные ошибки
    8. `go doc`. Показывает документацию к пакету
    9. `go list`. Выводит список зависимостей в проекте

---

## Корутины

--- 

## Каналы

---

## Дженерики

---

## Стиль

- Стиль в гошке весьма интересный. Например, не стоит повторять в названиях методов наименования пакетов, а также для методов не стоит повторять наименования аргументов, принимаемых методом, в его названии. Честно, не понимаю, почему - при использовании имен аргументов в имени метода становится лучше понятно, с чем мы взаимодействуем, но здесь это харам. Т.е. является харамом что-то вроде: 

    ```go
    package yamlconfig

    func ParseYAMLConfig(input string) (*Config, error)

    func (c *Config) WriteConfigTo(w io.Writer) (int64, error)

    func OverrideFirstWithSecond(dest, source *Config) error
    ```

    Но при этом в случае наличия метода со схожим названием при инициализации нового метода допускается указание дополнительной информации, которую лучше не включать в исходное название в соответствии с описанным выше: 

    ```go
    func (c *Config) WriteTextTo(w io.Writer) (int64, error)
    func (c *Config) WriteBinaryTo(w io.Writer) (int64, error)
    ```

    - Также не стоит добавлять фразы, описывающее получение информации. `Get`, например. Для метода который возвращает, допустим, наименование работы, лучше использовать название `JobName`, а не **Get**JobName
    - А для методов, которые что-то делают (записывают, например), будет правильным использовать глаголы в их названии - `WriteDetail`
    - Если разница метода только в байтах принимаемого значения (int32, int64), стоит указывать это в названии (ParseInt32 - ParseInt64)

# Советы по написанию поддерживаемых программ

## Идентификаторы

- Идентификаторы - синонимы имен. В принципе, если не смотреть на 'правильный стиль' с точки зрения самого кода, а смотреть на 'правильный стиль' с точки зрения человека, который будет его читать и поддерживать, здесь очень много аналогий с тем же питоном. Т.е. в ход идет этика разработчика, а не этика языка
- Имена переменным лучше давать исходя из их ясности, а не краткости, о которой я писал выше. Хорошее имя лаконично, оно является спасательным (имя метода должно описывать применение переменной, а не ее содержимое. Хорошее имя описывает результат функции или поведение метода, а не его реализацию. назначение пакета, а не его содержимое). Хорошее имя предсказуемо. Также важно упомянуть, что Go вляется языков со строгой типизацией, ввиду чего просто не позволит передать, например в параметр метода с типом **map** объект скалярного типа, ввиду чего суффиксы с наименованием обрабатываемого типа данных просто излишни:
    ```go
    var usersMap map[string]*User  // суффикс 'Map' излишен
    ```
- Интересная практика, о которой я прочитал - давать переменным/методам тем большее имя, чем дальше от момента инициализации переменной/метода она/он будет использован(-а). Т.е. если метод используется единожд сразу после инициализации, нет смысла делать его имя длинным, а если метод используется на протяжении длинного участка кода, стоит обозвать его более длинным именем и привнести больше контекста в это имя:

    ```go
    type Person struct {
        Name string
        Age  int
    }

    // AverageAge returns the average age of people.
    func AverageAge(people []Person) int {
        if len(people) == 0 {
            return 0
        }

        var count, sum int
        for _, p := range people {
            sum += p.Age
            count += 1
        }

        return sum / count
    }
    ```

    Грубо говоря, длина имени должна показывать важность переменной. Это логично, но раньше этот принцип я использовал на подсознательном уровне

- В Go принято соглашение, что интересно, о максимальном сокращении имен объектов. Зачастую, если рассматривать аргументы методов, например, аргумент является сокращением от его типа (структура Person -> аргумент p)

- Также рекомендуется использовать единый стиль деклораций (объявление и инициализация). Мне больше всего понравился вариант `var x int = 5`. Т.е. `var x int` преобладает над `var x = 0`. Остановлюсь на нем. Также, если требуется задать нулевое значение определенного типа, переменная просто объявляется, без инициализации (тоже использовать). Если 2 переменные близко связаны между собой, стоит проводить их декларацию в одной строке: 

    ```go
    var min, max := 0, 1000
    ```

- Если есть разногласия по стилю, их стоит выносить на обсуждения только в том случае, если код не проходит через `go fmt`

---

## Комментарии

- Важно понимать, что комментарий должен делать одну и только одну вещь из 3-х приведенных: 
    1. Объяснить, что делает код
    2. Объяснить, как код что-то делает
    3. Объяснить, зачем код что-то делает

- Комментарии в переменных и константах должны описывать их содержимое, а не предназначение: 
    ```go
    // sizeCalculationDisabled указывает, безопасно ли
    // рассчитать ширину и выравнивание типов. См. dowidth.
    var sizeCalculationDisabled bool
    ```

- Иногда лучшее имя для переменной скрывается в комментарии. В примере приведено имя `register` - понятно, что регистр, но регистр чего? если добавить уточнение, что он связан с sql, например `sqlDrivers`, то станет лучше: 
    ```go
    // реестр драйверов SQL
    var registry = make(map[string]*sql.Driver)
    ```

- При разработке своих пакетов лучше всего комментировать все, что не является кратким и явным одновременно. При этом любая функция должна быть прокомментирована вне зависимости от ее краткости и сложности. Однако есть исключение - не стоит документировать методы, реализующие интерфейс: 
    ```go
    // Read реализует интерфейс io.Reader
    func (r *FileReader) Read(buf []byte) (int, error)
    ```

    Приведенный комментарий приосит больше ущерба, чем пользы - в нем нет никакой информации кроме той, что и так ясна из строки кода

- Если стало понятно, что код перед глазами попахивает, лучше всего его переписать, чем насыщать его комментариями. Если столкнусь с подобным комментарием, лучше завести тикет с напоминанием о рефакторинге. Также буде правильным решением в тикете указать ник человека, который либо заметил проблему, либо к которому лучше обратиться за решением проблемы

---

## Структура пакета

---

## Структура проекта

---

## Структура API

---

## Обработка ошибок

---

## Параллелизм

