# Solution submitted to: https://adventofcode.com/2020/day/10

# Want to implement pure python
# So don't use np.diff for difference
# import numpy as np 
from functools import lru_cache

def difference(data, n=1):
  def dff(data):
    return [x - y for (x, y) in zip(data[1:], data[:-1])]
  for i in range(n): data = dff(data)
  return data

def part1(data):
  diff = difference(data)
  return sum([1 for i in diff if i == 1]) * sum([1 for i in diff if i == 3]) 

@lru_cache(maxsize=256)
def part2(i):
  if i == len(data) - 1:
    return 1
  return sum(
    [part2(j) for j in range(i + 1, min(i + 4, len(data))) if data[j] - data[i] <= 3]
  )

if __name__ == "__main__":
  with open("../data/day10.txt", "r") as f:
  # with open("test.txt", "r") as f:
    data = list(map(int, f.readlines()))
  data.sort()
  data.insert(0, 0) # Seat output joltage
  data.append(data[-1] + 3) # Device joltage

  res = part1(data)
  print("Part I:", res)

  res = part2(0)
  print("Part II:", res)
