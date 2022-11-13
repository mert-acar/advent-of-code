# Solution submitted to: https://adventofcode.com/2020/day/9

def part1(data, N=25):
  for i in range(N, len(data)):
    pre = data[i-N:i]
    target = data[i]
    flag = False
    for j in range(len(pre)):
      diff = target - pre[j]
      if (diff in pre[j:]) and (diff != pre[j]):
        flag = True
        break
    if not flag:
      return target
  return -1

def part2(data, num, N):
  n = len(data)
  l = []
  for i in range(n):
    for j in range(i + 1, n):
      if sum(data[i:j]) == num:
        return min(data[i:j]) + max(data[i:j]) 

if __name__ == "__main__":
  with open("../data/day9.txt", "r") as f:
  # with open("test.txt", "r") as f:
    data = list(map(int, f.readlines()))
  N = 25
  res1 = part1(data, N)
  print("Part I:", res1)
  res2 = part2(data, res1, N)
  print("Part II:", res2)
