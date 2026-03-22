# SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
#
# SPDX-License-Identifier: MPL-2.0

def factorial(number: int) -> int:
    product: int = 1
    for i in range(1, number+1):
        product = product * i
    return product


print(factorial(5))
