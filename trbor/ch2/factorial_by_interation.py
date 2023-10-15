def factorial(number: int) -> int:
    product: int = 1
    for i in range(1, number+1):
        product = product * i
    return product


print(factorial(5))
