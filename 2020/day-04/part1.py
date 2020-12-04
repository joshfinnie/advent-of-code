import sys

def get_answers(arr):
    must_include = [
        "byr",
        "iyr",
        "eyr",
        "hgt",
        "hcl",
        "ecl",
        "pid",
    ]
    valid_passports = 0
    for a in arr:
        id = a.split()
        if len(id) == 8:
            valid_passports += 1
        else:
            parts = []
            valid = True
            for p in id:
                parts.append(p.split(":")[0])
            for include in must_include:
                if include not in parts:
                    valid = False
            if valid:
                valid_passports += 1


    return valid_passports



with open(sys.argv[1]) as f:
    arr = f.read().split("\n\n")
    print(get_answers(arr))
