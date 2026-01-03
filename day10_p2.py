import numpy as np
from scipy.optimize import linprog
import re

# Go was not the ideal tool for linear programming :'(

def solve(a, b):
    a = np.array(a)
    b = np.array(b)
    c = np.ones(len(a[0]))
    res = linprog(c, A_eq=a, b_eq=b, bounds=(0, None), integrality=c, method='highs')
    if res.success:
        return int(res.fun)
    return -1

def parse(line):
    buttons = re.findall(r'\(([\d,]+)\)', line)
    buttons = [[int(x) for x in button.split(",")] for button in buttons]
    num_buttons = len(buttons)

    consts = re.search(r'\{([\d,]+)\}', line)
    counters = [int(x) for x in consts.group(1).split(',')] if consts else []
    num_counters=len(counters)

    a = []
    for i, counter in enumerate(counters):
        a_counter = [0 for i in range(num_buttons)]
        for j, button in enumerate(buttons):
            if i in button:
                a_counter[j] = 1
        a.append(a_counter)

    return a, counters

s = 0
for i, line in enumerate(open("input/day10.txt").readlines()):
    a, b = parse(line)
    r = solve(a, b)
    s += r

print(s)
