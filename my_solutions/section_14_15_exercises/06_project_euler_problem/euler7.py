"""
Problem 7. 10001st prime.
http://static.projecteuler.net/problem=7
"""

import math


def is_prime(num):
    """
    Checks that given number is a prime number.
    """
    if num <= 3:
        return num == 2 or num == 3

    for i in range(5, int(math.sqrt(num)) + 1, 6):
        if num % i == 0 or num % (i + 2) == 0:
            return False
    return True


def main():
    counter = 1
    num = 1
    while counter < 10001:
        num += 2
        if is_prime(num):
            counter += 1
    print(num)


if __name__ == '__main__':
    main()
