# Solution submitted to: https://adventofcode.com/2020/day/3
def solve(data, st):
  x = 0
  count = 0
  for i in range(st[0], len(data), st[0]):
    x += st[1]
    x = x % len(data[i])
    if data[i][x] == '#':
      count += 1
  return count
      

if __name__ == "__main__":
  with open("../data/day3.txt", "r") as f:
    data = list(map(lambda x: x.strip(), f.readlines()))

  res = solve(data, [1, 3])
  print("Part I:", res)

  requirements = [[1, 1], [1, 5], [1, 7], [2, 1]]
  for req in requirements:
    res *= solve(data, req)
  print("Part II:", res)
