# solution submitted to: https://adventofcode.com/2022/day/1

if __name__ == "__main__":
  with open('../data/day1.txt', 'r') as f:
    elves = f.read().split('\n\n')
  calories = []

  for elf in elves:
    food = sum(list(map(int, elf.strip().split('\n'))))
    calories.append(food)
  calories.sort()

  print("Part I:", calories[-1])
  print("Part II:", sum(calories[-3:]))
