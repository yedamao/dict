#!/usr/bin/env python
# -*- coding: utf-8 -*-

import re

pattern = re.compile(r'app')

with open('dict/wordlist') as f:
    content = f.read().decode('utf-8')

match = pattern.findall(content)
print match

for _ in match:
    print _
