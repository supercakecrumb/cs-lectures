# C: простая среда для Ubuntu и Windows (WSL)

Ниже — минимальные шаги, чтобы компилировать и запускать учебные программы на C с GCC, без make и отладки. Редактор — VS Code. Для Windows предполагается, что установлен WSL с Ubuntu.

## Окружение

0. Открой терминал (для винды WSL профиль Windows Terminal)
1. Обновить индексы: `sudo apt update`
2. Установить GCC и инструменты: `sudo apt install -y build-essential`

## Первый проект (Hello World)
1. Создай файл `hello.c` со следующим содержимым:
   ```c
   #include <stdio.h>

   int main() {
       printf("Hello, world!\n");
       return 0;
   }
   ```
2. Собери: `gcc hello.c -o hello`
3. Запусти: `./hello`

Готово — можно писать и запускать простые учебные программы на C.
