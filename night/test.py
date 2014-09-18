#!/usr/bin/env python

with open('phone') as f:
    content = f.read()

for line in content.split('\n'):
    for phone in line.split(':')[1:]:
        print phone
