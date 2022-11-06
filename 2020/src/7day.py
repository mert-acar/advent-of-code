# Solution submitted to: https://adventofcode.com/2020/day/7
import re

def countTotal(table, bag):
  return 1 + sum(int(n) * countTotal(table, color) for n, color in table[bag])

def containsSG(table, color):
  if color == "shiny gold": 
    return True
  else:
    return any(containsSG(table, c) for _, c in table[color])

def solve(data):
  table = {}
  c1, c2 = 0, 0
  for rule in data:
    color = re.match(r"(.+?) bags contain", rule)[1]
    table[color] = re.findall(r"(\d+?) (.+?) bags?", rule)

  for bag in table:
    if containsSG(table, bag):
      c1 += 1
  c1 = c1 - 1                              # A shiny gold bag cannot contain itself!
  c2 = countTotal(table, 'shiny gold') - 1 # A shiny gold bag cannot contain itself!
  return c1, c2

if __name__ == "__main__":
  with open("../data/day7.txt", "r") as f:
    data = f.read().splitlines()

  res1, res2 = solve(data)
  print("Part I:", res1)
  print("Part II:", res2)
