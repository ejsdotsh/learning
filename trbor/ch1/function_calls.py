# SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
#
# SPDX-License-Identifier: MPL-2.0

def a():
    print("a() was called.")
    b()
    print("a() is returning.")


def b():
    print("b() was called.")
    c()
    print("b() is returning.")


def c():
    print("c() was called.")
    print("c() is returning.")


a()
