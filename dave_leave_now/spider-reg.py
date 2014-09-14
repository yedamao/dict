#!/usr/bin/env python
# -*- coding: utf-8 -*-

import re
import urllib
import codecs


pattern = re.compile('<span class="phonetic">(.*)</span>')


def spider_phonetic(word):
    URL = 'http://dict.youdao.com/search?q=' + word + '&keyfrom=dict.index'

    f = urllib.urlopen(URL)
    content = ''
    try:
        content = f.read().decode('utf-8')
    except UnicodeDecodeError:
        print "%s UnicodeDecodeError!!!" % word

    match = pattern.findall(content)
    line = word + ':'
    for _ in match:
        try:
            line = line + _ + ':'
        except:
            print "Code Error"
            pass
    if len(line.split(':')) > 2:
        print line
        write_line(line)


def write_line(line):
    with codecs.open('phonetic', 'a+', encoding='utf-8') as f:
        f.write(line + '\n')
        print "add successfully"


with open('american-english') as f:
    for line in f:
        word = line.strip().lower()
        spider_phonetic(word)
