#!/usr/bin/python3


def staircase(n):
    level_string = ""
    for i in range(n):
        level_string += " " * (n - 1 - i)
        level_string += "#" * (i + 1)
        level_string += "\n"
    print(level_string)
    return level_string


staircase(6)
