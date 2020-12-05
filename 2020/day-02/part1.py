import sys


def confirm_password(rule):
    count, letter, password = rule.split(" ")
    minimum, maximum = count.split("-")
    letter = letter.strip(":")

    letter_count = password.count(letter)

    if letter_count >= int(minimum) and letter_count <= int(maximum):
        return 1
    else:
        return 0


values = []
with open(sys.argv[1]) as inputfile:
    for line in inputfile:
        values.append(line.strip())

totals = 0
for p in values:
    totals += confirm_password(p)
print(totals)
