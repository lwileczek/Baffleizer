#!/usr/bin/env python3

import time
import numpy as np
from datetime import datetime


def print_hello_world(logical_argument: str = "Hello", second_argument: str = "World", should_print: bool = True):
    """Print hello, world but have a bunch of typing and default args"""
    if should_print:
        print(logical_argument, ", ", second_argument)


def long_docstring(foo, bar):
    """Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam posuere
    enim non ipsum tristique porta. Vivamus nec urna orci. Quisque laoreet
    aliquet leo, sed porta metus rhoncus in. Pellentesque malesuada, erat a
    facilisis vehicula, ipsum nibh sodales ante, id efficitur erat diam vitae massa.
    Nam ut luctus diam, sit amet cursus libero. Fusce urna risus, fringilla quis
    euismod id, gravida eu lorem. Donec dictum metus in diam iaculis, vitae
    aliquam est dapibus. Integer pellentesque est sodales lacus porta, ac volutpat
    dolor convallis. Suspendisse a nisl non sem imperdiet semper. Suspendisse
    malesuada erat sed est venenatis, vel vestibulum tortor commodo.
    Aliquam eget eros vitae sem dictum iaculis non quis nunc. Integer ut ex purus.
    """

    print(foo, bar)


def fib(n):
    """Compute a Fibonacci Number
    making the docstring two lines for complexity"""
    if n <= 1:
        return n
    elif n == 2:
        return 1
    else:
        return fib(n-1) + fib(n-2)


def get_today():
    now = datetime.now()  # current date and time
    year = now.strftime("%Y")
    month = now.strftime("%m")
    day = now.strftime("%d")
    time = now.strftime("%H:%M:%S")
    date_time = now.strftime("%m/%d/%Y, %H:%M:%S")
    return date_time


class ExampleClass:
    random_class_value = 1

    def __init__(self, x, y):
        self.x = x
        self.y = y

    def first_method(self):
        return self._internal_method(self.x) + self.y

    def _internal_method(self, z):
        if z is None:
            z = self.x
        return z**2


if __name__ == "__main__":
    print_hello_world()
    long_docstring("an example of a boolean is", False)
    time.sleep(1)
    x = fib(8)
    assert x == 21, "fibonacci is messed up"
    dt = get_today()
    print(dt)

    # The best prime numbers
    class_mate = ExampleClass(2, 3)
    answer = class_mate.first_method()

    # Numpy stuff I guess
    my_array = np.array([[1, 2, 3], [2, 5, 11], [4, 23, 101]])
    q, r = np.linalg.qr(my_array)
    print(q)
    print(r)
