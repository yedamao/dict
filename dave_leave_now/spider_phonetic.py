# -*- coding: utf-8 -*-
#!/usr/bin/env python

import urlib
import MySQLdb
import re

db = MySQLdb.connect("localhost", user="root", passwd="telecom",
                     db="dictionary", charset="utf8")
cursor = db.cursor()

def spider_phonetic(word):
    """
    this function accept a word, and spider the data.
    and print it.
    """
    URL = 'http://dict.youdao.com/search?q=' + word + '&keyfrom=dict.index'
    pattern = re.compile('<span class="phonetic">(.*)</span>')
    phonetic_us = ''
    phonetic_eng = ''
    content = ''

    f = urllib.urlopen(URL)
    try:
        content = f.read().decode('utf-8')
    except UnicodeDecodeError:
        print "%s UnicodeDecodeError!!!!" % word

    match = pattern.findall(content)

    if match:
        phonetic_eng = match[0]
        phonetic_us = match[1]
        insert_record(word, match.group())
    else:
        print "%s not match!!!!\n" % word


def insert_record(word, content):
    """
    this function insert record to database
    """
    sql = "insert into word_content( word, content) values ("
    sql = sql + "'" + word.decode('utf-8') + "'" + ',' + "'" + content + "'" + ')'

    try:
        cursor.execute(sql)
        db.commit()
        print "%s record insert sucessful" % word
    except MySQLdb.Error, e:
        print "Mysql Error %d: %s" % (e.args[0], e.args[1])
        db.rollback()


with open('american-english') as f:
    for line in f:
        word = line.strip()
        spider_word(word)


db.close()
