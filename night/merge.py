# -*- coding: utf-8 -*-
#!/usr/bin/env python


with open('dictionary') as f:
    dictionary = f.read()

with open('phone') as f:
    phonetic = f.read()

for line in dictionary.split('\n'):
    word = line.split(':')[0]
    means = line.split(':')[1:]
    for _ in phonetic.split('\n'):
        phonetic_word = _.split(':')[0]
        phone = _.split(':')[1:]
        if word.lower() == phonetic_word:
            new_line = word
            for ph in phone:
                new_line = new_line + ':' + ph
            new_line = new_line[:-1]
            for mean in means:
                new_line = new_line + ':' + mean
            line = new_line[:-1]
    print line
    with open('new', 'a+') as f:
        f.write(line + '\n')
