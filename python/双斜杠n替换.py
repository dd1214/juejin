import json

s = input()

cnt = s.count("\\n")
s.replace("\\n", "\n", cnt)

print(s)
