import sys


def confirm_password(rule):
    count, letter, password = rule.split(" ")
    minimum, maximum = count.split("-")
    letter = letter.strip(":")

    first = password[int(minimum) - 1]
    last = password[int(maximum) - 1]

    if (first == letter and last != letter) or (first != letter and last == letter):
        return 1
    else:
        return 0


values = []
with open(sys.argv[1]) as inputfile:
    for line in inputfile:
        values.append(line.strip())

# values = ["1-3 a: abcde","1-3 b: cdefg","2-9 c: ccccccccc"]
totals = 0
for p in values:
    totals += confirm_password(p)
print(totals)
