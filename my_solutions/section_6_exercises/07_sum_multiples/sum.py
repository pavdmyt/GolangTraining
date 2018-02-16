# -*- coding: utf-8 -*-


MAX = 10 ** 8


def impl1():
    multiples = (
        i for i in range(MAX)
        if i % 3 == 0 or i % 5 == 0
    )
    print(sum(multiples))


def impl2():
    res = 0
    for i in range(MAX):
        if i%3 == 0 or i%5 == 0:
            res += i
    print(res)


def main():
    impl2()


if __name__ == '__main__':
    main()


# impl1
#
# $ time python sum.py
# 2333333316666668
# python sum.py  18.42s user 0.06s system 99% cpu 18.567 total

# impl2
#
# time python sum.py
# 2333333316666668
# python sum.py  18.52s user 0.06s system 99% cpu 18.647 total
