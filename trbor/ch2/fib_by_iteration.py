# SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
#
# SPDX-License-Identifier: MPL-2.0

def fibonacci(nth_number) -> int:
    a, b = 1, 1
    print(f"a = {a}, b = {b}")
    for i in range(1, nth_number):
        a, b = b, a + b  # get the next number
        print(f"a = {a}, b = {b}")
    return a


print(fibonacci(10))
