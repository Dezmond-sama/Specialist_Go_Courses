# Task 2

## Input:

User should enter a password with a right complexity. It should const of:

-   digits,
-   english letters in upper and lower cases
-   special symbols => "\_!@#$%^&"

4 sets of symbols.

Password should include symbols from each of the sets.

Length should be in range 8-15.
Max tries: 5.

Output the current try number.

-   Output the ecplanation, why the password is bad.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "\_!@#$%^&"

## Output:

Write, if the password is appropriate.

## Example:

good -> o58anuahaunH!
good -> aaaAAA111!!!
bad -> saucacAusacu8

---

# Задача №2

## Вход:

Пользователь должен ввести правильный пароль, состоящий из:

-   цифр,
-   букв латинского алфавита(строчные и прописные) и
-   специальных символов special = "\_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.

-   Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "\_!@#$%^&"

## Выход:

Написать, что ввели правильный пароль.

## Пример:

хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
