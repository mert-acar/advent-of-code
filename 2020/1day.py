# Solution submitted to: https://adventofcode.com/2020/day/1
def part1(data, target=2020):
  for i in range(len(data)):
    if target - data[i] in data[i:]:
      return data[i] * (target - data[i])
  return -1


def part2(data, target=2020):
  for i in range(len(data)):
    residual = target - data[i]
    for j in range(i, len(data)):
      if residual - data[j] in data[j:]:
        return data[j] * (residual - data[j]) * data[i]
  return -1

if __name__ == "__main__":
  with open("../data/day1.txt", "r") as f:
    data = list(map(int, f.readlines()))
  
  res = part1(data)
  print('Part I:', res)

  res = part2(data)
  print('Part II:', res)
