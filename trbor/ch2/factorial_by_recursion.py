def factorial(number: int) -> int:
    if number == 1:
        # base case
        return 1
    else:
        # recursive case
        return number * factorial(number - 1)


print(factorial(5))
