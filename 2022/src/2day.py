# Solution submitted to: https://adventofcode.com/2022/day/2

if __name__ == "__main__":
  with open('../data/day2.txt', 'r') as f:
    data = list(map(lambda x: x.strip().replace(' ', ''), f.readlines()))

  scores = {
    'X': 1,
    'Y': 2,
    'Z': 3
  }

  outcomes = {
    'AX': 3,
    'AY': 6,
    'AZ': 0,
    'BX': 0,
    'BY': 3,
    'BZ': 6,
    'CX': 6,
    'CY': 0,
    'CZ': 3
  }

  outcomes_2 = {
    'AX': ['Z', 0],
    'AY': ['X', 3],
    'AZ': ['Y', 6],
    'BX': ['X', 0],
    'BY': ['Y', 3],
    'BZ': ['Z', 6],
    'CX': ['Y', 0],
    'CY': ['Z', 3],
    'CZ': ['X', 6],
  }

  score_1 = 0
  score_2 = 0
  for st in data:
    score_1 += scores[st[-1]] + outcomes[st]
    score_2 += scores[outcomes_2[st][0]] + outcomes_2[st][1]

  print("Part I:", score_1)
  print("Part II:", score_2)
