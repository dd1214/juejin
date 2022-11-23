import re

txt = ""
while True:
    s = input()
    txt += s
    if s == '0 0':
        break
bs64_str = re.findall('mark_content:"(.*?)",display_count:d', txt)[0]
print(bs64_str)
