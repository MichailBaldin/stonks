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
