import sys


def get_values(arr):
    sum = 2020
    for i in arr:
        for j in arr:
            if i + j == sum:
                return i * j


values = []
with open(sys.argv[1]) as inputfile:
    for lines in inputfile:
        values.append(int(lines))

print(get_values(values))
