# SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
#
# SPDX-License-Identifier: MPL-2.0

def factorial(number: int) -> int:
    if number == 1:
        # base case
        return 1
    else:
        # recursive case
        return number * factorial(number - 1)


print(factorial(5))
