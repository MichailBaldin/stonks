# Manjaro: Пошаговая установка Python + uv + ruff

## 🔄 Шаг 1: Обновление системы

```bash
# Обновляем систему
sudo pacman -Syu

# Перезагружаемся, если были обновлены ядро или критические компоненты
sudo reboot  # (при необходимости)
```

## 🐍 Шаг 2: Установка Python

```bash
# Устанавливаем Python (обычно уже есть в системе)
sudo pacman -S python python-pip

# Проверяем версию
python --version
python3 --version

# Устанавливаем дополнительные пакеты для разработки
sudo pacman -S python-pipenv python-virtualenv
```

### Проверка установки Python
```bash
# Должно показать версию Python 3.11+ 
python3 --version

# Проверяем pip
pip --version
```

## ⚡ Шаг 3: Установка uv (современный менеджер пакетов Python)

### Способ 1: Через curl (рекомендуемый)
```bash
# Скачиваем и устанавливаем uv
curl -LsSf https://astral.sh/uv/install.sh | sh

# Перезагружаем shell или добавляем в PATH
source $HOME/.cargo/env

# Или добавляем в ~/.bashrc / ~/.zshrc
echo 'export PATH="$HOME/.cargo/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### Способ 2: Через pip
```bash
# Альтернативный способ
pip install uv
```

### Способ 3: Через AUR (если есть yay)
```bash
# Если у вас установлен yay
yay -S uv

# Или через pamac
pamac install uv
```

### Проверка установки uv
```bash
uv --version
# Должно показать версию uv
```

## 🦀 Шаг 4: Установка ruff (быстрый линтер/форматтер)

### Способ 1: Через uv (рекомендуемый)
```bash
uv tool install ruff
```

### Способ 2: Через pacman
```bash
sudo pacman -S ruff
```

### Способ 3: Через pip
```bash
pip install ruff
```

### Способ 4: Через cargo (если есть Rust)
```bash
cargo install ruff
```

### Проверка установки ruff
```bash
ruff --version
# Должно показать версию ruff
```

## 🛠️ Шаг 5: Настройка среды разработки

### Создание проекта с uv
```bash
# Создаем новый проект
mkdir my-python-project
cd my-python-project

# Инициализируем проект с uv
uv init

# Создаем виртуальное окружение
uv venv

# Активируем виртуальное окружение
source .venv/bin/activate

# Устанавливаем пакеты через uv
uv add requests pandas numpy
```

### Настройка ruff
```bash
# Создаем конфигурационный файл ruff
cat > pyproject.toml << 'EOF'
[tool.ruff]
# Длина строки
line-length = 88

# Включаем дополнительные правила
select = [
    "E",  # pycodestyle errors
    "W",  # pycodestyle warnings
    "F",  # Pyflakes
    "I",  # isort
    "B",  # flake8-bugbear
    "C4", # flake8-comprehensions
    "UP", # pyupgrade
]

# Исключаем файлы
exclude = [
    ".git",
    "__pycache__",
    ".venv",
    "build",
    "dist",
]

[tool.ruff.format]
# Используем двойные кавычки
quote-style = "double"
EOF
```

## 📝 Шаг 6: Создание тестового файла

```bash
# Создаем тестовый Python файл
cat > main.py << 'EOF'
def greet(name: str) -> str:
    """Приветствует пользователя по имени."""
    return f"Привет, {name}!"


def main():
    """Главная функция."""
    user_name = input("Введите ваше имя: ")
    message = greet(user_name)
    print(message)


if __name__ == "__main__":
    main()
EOF
```

## 🧪 Шаг 7: Тестирование установки

### Проверяем Python
```bash
python main.py
# Должно запросить имя и поприветствовать
```

### Проверяем ruff
```bash
# Проверка кода на ошибки
ruff check main.py

# Форматирование кода
ruff format main.py

# Исправление автоматически исправимых ошибок
ruff check --fix main.py
```

### Проверяем uv
```bash
# Показать информацию о проекте
uv pip list

# Добавить зависимость
uv add requests

# Запустить скрипт
uv run main.py
```


## ⚙️ Шаг 8: Настройка VS Code (опционально)

### Установка VS Code
```bash
# Через pamac
pamac install visual-studio-code-bin

# Или через AUR
yay -S visual-studio-code-bin
```

### Установка расширений для Python
```bash
# Открываем VS Code
code .

# Устанавливаем расширения (через интерфейс VS Code):
# - Python
# - Pylance  
# - Ruff
```

### Настройка VS Code для проекта
```bash
# Создаем настройки VS Code для проекта
mkdir .vscode
cat > .vscode/settings.json << 'EOF'
{
    "python.defaultInterpreterPath": "./.venv/bin/python",
    "python.formatting.provider": "none",
    "python.linting.enabled": false,
    "[python]": {
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.fixAll.ruff": true,
            "source.organizeImports.ruff": true
        },
        "editor.defaultFormatter": "charliermarsh.ruff"
    },
    "ruff.args": ["--config=pyproject.toml"]
}
EOF
```

## 🚀 Шаг 9: Создание алиасов (опционально)

```bash
# Добавляем удобные алиасы в ~/.bashrc или ~/.zshrc
cat >> ~/.bashrc << 'EOF'

# Python development aliases
alias py="python"
alias venv="python -m venv"
alias activate="source .venv/bin/activate"
alias lint="ruff check"
alias format="ruff format"
alias fix="ruff check --fix"

# uv aliases
alias uv-init="uv init && uv venv && source .venv/bin/activate"
alias uv-add="uv add"
alias uv-run="uv run"
EOF

# Перезагружаем bash
source ~/.bashrc
```

## ✅ Шаг 10: Финальная проверка

```bash
# Проверяем все установленные инструменты
echo "=== Python ==="
python --version

echo "=== uv ==="
uv --version

echo "=== ruff ==="
ruff --version

echo "=== pip ==="
pip --version

echo "=== Тест работы ==="
# Создаем быстрый тест
cat > test.py << 'EOF'
import sys
print(f"Python версия: {sys.version}")
print("Все работает! 🎉")
EOF

# Запускаем через разные способы
python test.py
uv run test.py

# Проверяем ruff
ruff check test.py
ruff format test.py

echo "Установка завершена успешно! ✅"
```

## 📋 Краткая шпаргалка команд

### uv команды
```bash
uv init                    # Создать новый проект
uv venv                   # Создать виртуальное окружение  
uv add package            # Добавить пакет
uv remove package         # Удалить пакет
uv pip install package    # Установить через pip
uv run script.py          # Запустить скрипт
uv sync                   # Синхронизировать зависимости
```

### ruff команды
```bash
ruff check .              # Проверить код
ruff check --fix .        # Исправить автоматически
ruff format .             # Отформатировать код
ruff check --watch .      # Проверять в реальном времени
```

### Активация окружения
```bash
source .venv/bin/activate  # Активировать
deactivate                 # Деактивировать
```

---

**Готово!** Теперь у вас есть современная среда разработки Python с быстрыми инструментами uv и ruff на Manjaro Linux. 🎉